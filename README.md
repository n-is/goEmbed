# Lua-Helper

[![alt text](https://godoc.org/github.com/n-is/lua-helper/lua?status.svg)](https://godoc.org/github.com/n-is/lua-helper/lua)
[![Build Status](https://travis-ci.org/n-is/lua-helper.svg)](https://travis-ci.org/n-is/lua-helper)
[![Coverage Status](https://coveralls.io/repos/github/n-is/lua-helper/badge.svg)](https://coveralls.io/github/n-is/lua-helper?branch=master)

A helper repository for gopher-lua, to make coding easier. Contains some tests
and benchmarking codes, to compare the performance with python3.

Examples on using this package to load and run a lua script can be found in
[examples_test.go](lua/examples_test.go).

[scripts/lua](lua/tests/scripts/lua) contains lua test scripts that are loaded
using the [gopher-lua](https://github.com/n-is/gopher-lua) package, and run for
benchmarking, using `go test`.
[scripts/python](lua/tests/scripts/python) contains python test scripts that
perform identical work to the respective lua script. `pytest` module was used
for testing and benchmarking the python scripts.

[tests/.benchmark](lua/tests/.benchmark) contains the result of benchmarking.
Any file in [.benchmark/Windows-CPython-3.7-64bit](lua/tests/.benchmark/Windows-CPython-3.7-64bit)
contains the detail info on the machine, on which the benchmarks were run.

Following Programs are used for benchmarking purposes:

* Fibonacci Series
* String Concatenation
* NBody Simulation
* Binary Tree
