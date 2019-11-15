# Lua-Helper

[![alt text](https://godoc.org/github.com/n-is/lua-helper/lua?status.svg)](https://godoc.org/github.com/n-is/lua-helper/lua)
[![Build Status](https://travis-ci.org/n-is/lua-helper.svg?branch=master)](https://travis-ci.org/n-is/lua-helper)

A helper repository for gopher-lua, to make coding easier. Contains some tests
and benchmarking codes, to compare the performance with python3.

[scripts/lua](lua/tests/scripts/lua) contains lua test scripts that are loaded
using the [gopher-lua](https://github.com/n-is/gopher-lua) package, and run for
benchmarking, using `go test`.
[scripts/python](lua/tests/scripts/python) contains python test scripts that
perform identical work to the respective lua script. `pytest` module was used
for testing and benchmarking the python scripts.

[tests/.benchmark](lua/tests/.benchmark) contains the result of benchmarking.
Any file in [.benchmark/Windows-CPython-3.7-64bit](lua/tests/.benchmark/Windows-CPython-3.7-64bit)
contains the detail info on the machine, on which the benchmarks were run.

Examples on using this package to load and run a lua script can be found in
[examples_test.go](lua/examples_test.go).
