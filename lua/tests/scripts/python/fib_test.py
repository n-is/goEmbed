from fib import *
import pytest
import time

def test_fib():
        assert fib(-20) == -1
        assert fib(-1) == -1
        assert fib(0) == 0
        assert fib(1) == 1
        assert fib(2) == 1
        assert fib(3) == 2
        assert fib(4) == 3
        assert fib(5) == 5
        assert fib(6) == 8
        assert fib(20) == 6765
        assert fib(40) == 102334155
        assert fib(50) == 12586269025
        assert fib(65) == 17167680177565
        assert fib(75) == 2111485077978050
        assert fib(79) == 14472334024676221


@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_FibM20(benchmark):
        res = benchmark(fib, -20)
        assert res == -1

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_FibM2(benchmark):
        res = benchmark(fib, -2)
        assert res == -1

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_FibM1(benchmark):
        res = benchmark(fib, -1)
        assert res == -1

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib0(benchmark):
        res = benchmark(fib, 0)
        assert res == 0

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib1(benchmark):
        res = benchmark(fib, 1)
        assert res == 1

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib2(benchmark):
        res = benchmark(fib, 2)
        assert res == 1

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib3(benchmark):
        res = benchmark(fib, 3)
        assert res == 2

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib10(benchmark):
        res = benchmark(fib, 10)
        assert res == 55

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib20(benchmark):
        res = benchmark(fib, 20)
        assert res == 6765

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib40(benchmark):
        res = benchmark(fib, 40)
        assert res == 102334155

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib60(benchmark):
        res = benchmark(fib, 60)
        assert res == 1548008755920

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib80(benchmark):
        res = benchmark(fib, 80)
        assert res == 23416728348467685

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib100(benchmark):
        res = benchmark(fib, 100)

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=20,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib1000(benchmark):
        res = benchmark(fib, 1000)

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib10000(benchmark):
        res = benchmark(fib, 10000)

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib100000(benchmark):
        res = benchmark(fib, 100000)

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib1000000(benchmark):
        res = benchmark(fib, 1000000)

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib10000000(benchmark):
        res = benchmark(fib, 10000000)

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib100000000(benchmark):
        res = benchmark(fib, 100000000)

@pytest.mark.benchmark(
    group="Math",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Fib1000000000(benchmark):
        res = benchmark(fib, 1000000000)
