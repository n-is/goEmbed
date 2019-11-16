from concat import *
import pytest
import time

def test_concat():
        nums = {"first": "3", "last":"4"}
        assert concat(nums) == "34"
        name = {"first": "Nischal", "last": "Nepal"}
        assert concat(name) == "NischalNepal"
        greet = {"first": "Hello", "last":"World!!"}
        assert concat(greet) == "HelloWorld!!"

def genString(n):
        greet = {"first": randomword(n), "last":randomword(n)}
        return [greet]

@pytest.mark.benchmark(
    group="String",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_concatNums(benchmark):
        nums = {"first": "3", "last":"4"}
        res = benchmark(concat, nums)
        assert res == "34"

@pytest.mark.benchmark(
    group="String",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_concatName(benchmark):
        name = {"first": "Nischal", "last": "Nepal"}
        res = benchmark(concat, name)
        assert res == "NischalNepal"

@pytest.mark.benchmark(
    group="String",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_concatGreet(benchmark):
        greet = {"first": "Hello", "last":"World!!"}
        res = benchmark(concat, greet)
        assert res == "HelloWorld!!"

@pytest.mark.benchmark(
    group="String",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Concat10(benchmark):
        res = benchmark.pedantic(concat, genString(10))

@pytest.mark.benchmark(
    group="String",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Concat100(benchmark):
        res = benchmark.pedantic(concat, genString(100))

@pytest.mark.benchmark(
    group="String",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Concat1000(benchmark):
        res = benchmark.pedantic(concat, genString(1000))

@pytest.mark.benchmark(
    group="String",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Concat10000(benchmark):
        res = benchmark.pedantic(concat, genString(10000))

@pytest.mark.benchmark(
    group="String",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Concat100000(benchmark):
        res = benchmark.pedantic(concat, genString(100000))

@pytest.mark.benchmark(
    group="String",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Concat1000000(benchmark):
        res = benchmark.pedantic(concat, genString(1000000))

@pytest.mark.benchmark(
    group="String",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Concat10000000(benchmark):
        res = benchmark.pedantic(concat, genString(10000000))

@pytest.mark.benchmark(
    group="String",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Concat100000000(benchmark):
        res = benchmark.pedantic(concat, genString(100000000))
