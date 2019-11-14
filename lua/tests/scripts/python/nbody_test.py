from nbody import *
import pytest
import time

def test_nbody():
        nb = NBodySystem()
        energy = nb.run(10)
        target = -0.16907302171469984
        assert abs(energy - target) < 1e-8

def runNBody(n):
        nb = NBodySystem()
        nb.run(n)

@pytest.mark.benchmark(
    group="NBody",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_NBodyM20(benchmark):
        res = benchmark(runNBody, -20)

@pytest.mark.benchmark(
    group="NBody",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_NBody0(benchmark):
        benchmark(runNBody, 0)

@pytest.mark.benchmark(
    group="NBody",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_NBody10(benchmark):
        benchmark(runNBody, 10)

@pytest.mark.benchmark(
    group="NBody",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_NBody100(benchmark):
        benchmark(runNBody, 100)

@pytest.mark.benchmark(
    group="NBody",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_NBody1000(benchmark):
        benchmark(runNBody, 1000)

@pytest.mark.benchmark(
    group="NBody",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_NBody10000(benchmark):
        benchmark(runNBody, 10000)

@pytest.mark.benchmark(
    group="NBody",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_NBody100000(benchmark):
        benchmark(runNBody, 100000)

@pytest.mark.benchmark(
    group="NBody",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_NBody1000000(benchmark):
        benchmark(runNBody, 1000000)

@pytest.mark.benchmark(
    group="NBody",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_NBody10000000(benchmark):
        benchmark(runNBody, 10000000)
