// helm search repo kinguin-charts/app --versions | awk '/^kinguin-charts\/app/ {print $1 ":" $2}'
// Put chart in reversed order
// helm search repo kinguin-charts/app --versions | awk '/^kinguin-charts\/app/ {print $1 ":" $2}' | tac > charts.txt
package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path"
	"strings"
)

const defaultHelmRepo = "oci://europe-docker.pkg.dev/kinguin-artifacts/helm-charts/kinguin"

func main() {
	var file string
	var chart string
	var destRepo string
	var charts []string

	cmd := &cobra.Command{
		Use:     "chart-sync <chart-name:version>",
		Short:   "Syncs charts between repositories",
		Long:    "Example: chart-sync kinguin-charts/app:8.14.0",
		Args:    cobra.MaximumNArgs(1),
		PreRunE: validateArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(file)

			// Pipe
			if file != "" {
				file, err := os.Open(file)
				if err != nil {
					fmt.Printf("Couldn't open file: %s\n", err)
					return
				}
				defer file.Close()

				scanner := bufio.NewScanner(file)

				// Read file
				var line int
				for scanner.Scan() {
					line++
					chart = scanner.Text()
					err := validateChartFormat(chart)
					if err != nil {
						fmt.Printf("line %d: %s\n", line, err)
						return
					}
					charts = append(charts, chart)
				}

			}

			if len(args) == 1 {
				charts = append(charts, args[0])
			}

			for _, chart := range charts {
				err := fetchAndPushChart(chart, destRepo)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		},
	}

	cmd.Flags().StringVarP(&file, "file", "f", "", "A .txt file containing chart name with version in each line, f.ex: kinguin-charts/app:8.14.0")
	cmd.Flags().StringVarP(&destRepo, "destRepo", "r", defaultHelmRepo, "Destination helm repository")

	cmd.PreRunE = validateArgs
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return nil
	}

	return validateChartFormat(args[0])
}

func validateChartFormat(name string) error {
	if !strings.Contains(name, ":") {
		return fmt.Errorf("not valid Chart format: %s. Required format is: chart:version", name)
	}

	return nil
}

func fetchAndPushChart(chart string, destRepo string) error {
	if err := validateChartFormat(chart); err != nil {
		return err
	}

	parts := strings.SplitN(chart, ":", 2)
	chartName := parts[0]
	version := parts[1]
	chartPkg := fmt.Sprintf("%s-%s.tgz", strings.TrimSuffix(path.Base(chartName), "/"), version)

	// Fetch
	fetchCmd := exec.Command("helm", "fetch", chartName, "--version", version)
	fmt.Printf("Running: %s\n", fetchCmd.String())
	if err := fetchCmd.Run(); err != nil {
		return fmt.Errorf("Couldn't fetch chart: %v\n", err)
	}

	// Push
	pushCmd := exec.Command("helm", "push", chartPkg, destRepo)
	fmt.Printf("Running: %s\n", pushCmd.String())
	if err := pushCmd.Run(); err != nil {
		return fmt.Errorf("Couldn't push chart: %v\n", err)
	}

	// Remove chart file
	fmt.Printf("Running: %s\n", "rm "+chartPkg)
	if err := os.Remove(chartPkg); err != nil {
		return fmt.Errorf("Coulnd't remove chart package file: %v\n", err)
	}

	return nil
}
