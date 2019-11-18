# GoEmbed

[![alt text](https://godoc.org/github.com/n-is/goEmbed/lua?status.svg)](https://godoc.org/github.com/n-is/goEmbed)
[![Go Report Card](https://goreportcard.com/badge/github.com/n-is/goEmbed)](https://goreportcard.com/report/github.com/n-is/goEmbed)
[![Build Status](https://travis-ci.org/n-is/goEmbed.svg)](https://travis-ci.org/n-is/goEmbed)
[![Coverage Status](https://coveralls.io/repos/github/n-is/goEmbed/badge.svg)](https://coveralls.io/github/n-is/goEmbed?branch=master)

A helper repository for testing various languages that can be embedded in golang,
in a sandboxed environment. Contains some tests and benchmarking codes, to
compare the performance with python3.

Examples on using lua package to load and run a lua script can be found in
[examples_test.go](lua/examples_test.go).

[scripts/lua](tests/scripts/lua) contains lua test scripts that are loaded
using the [gopher-lua](https://github.com/n-is/gopher-lua) package, and run for
benchmarking, using `go test`.
[scripts/tengo](tests/scripts/tengo) contains tengo test scripts that are loaded
using the [tengo](https://github.com/n-is/tengo) package, and run for
benchmarking, using `go test`.
[scripts/python](tests/scripts/python) contains python test scripts that
perform identical work to the respective lua script. `pytest` module was used
for testing and benchmarking the python scripts.

[tests/.benchmark](tests/.benchmark) contains the result of benchmarking.
Any file in [.benchmark/Windows-CPython-3.7-64bit](tests/.benchmark/Windows-CPython-3.7-64bit)
contains the detail info on the machine, on which the benchmarks were run.

Following Programs are used for benchmarking purposes:

* Fibonacci Series
* String Concatenation
* NBody Simulation
* Binary Tree
