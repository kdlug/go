package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"go.uber.org/automaxprocs/maxprocs"
)

func main() {
	b := runtime.GOMAXPROCS(0)
	log.Printf("Incorrect CPU (of the machine): CPU[%d]", b)

	// Set the correct number of threads for the service
	// based on what is available either by the machine or quotas.
	if _, err := maxprocs.Set(); err != nil {
		fmt.Println("maxprocs: %w", err)
		os.Exit(1)
	}

	g := runtime.GOMAXPROCS(0)
	log.Printf("Correct CPU (of k8s quotas) CPU[%d]", g)
	defer log.Println("service ended")

}
