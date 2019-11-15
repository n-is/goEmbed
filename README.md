# lua-helper

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
