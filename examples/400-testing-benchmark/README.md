# Benchmark tests

Run all tests in the current project in verbose mode.

```bash
$go test -v

=== RUN   TestIfBasic
--- PASS: TestIfBasic (0.00s)
=== RUN   TestIfTableDriven
=== RUN   TestIfTableDriven/1
=== RUN   TestIfTableDriven/0
=== RUN   TestIfTableDriven/2
=== RUN   TestIfTableDriven/100
--- PASS: TestIfTableDriven (0.00s)
--- PASS: TestIfTableDriven/1 (0.00s)
--- PASS: TestIfTableDriven/0 (0.00s)
--- PASS: TestIfTableDriven/2 (0.00s)
--- PASS: TestIfTableDriven/100 (0.00s)
PASS
ok      github.com/kdlug/benchmark      0.345s
```

Run all benchmarks in the current project. All tests are run prior to benchmarks. The bench flag filters benchmark function names with a regexp.

```bash
go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/kdlug/benchmark
BenchmarkIf-8                   1000000000               0.2952 ns/op
BenchmarkSwitch-8               1000000000               0.2931 ns/op
BenchmarkSwitchConst-8          1000000000               0.2931 ns/op
PASS
ok      github.com/kdlug/benchmark      1.062s
```

So there is no difference in performance between switch and if-else. It is purely for aesthetic and code readability.
I prefer switch over if-else because it is easier to read 