from tree import *
import pytest
import time

@pytest.mark.benchmark(
    group="Binary Tree",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_TreeM20(benchmark):
        res = benchmark(tree, -20)

@pytest.mark.benchmark(
    group="Binary Tree",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Tree0(benchmark):
        res = benchmark(tree, 0)

@pytest.mark.benchmark(
    group="Binary Tree",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Tree2(benchmark):
        res = benchmark(tree, 2)

@pytest.mark.benchmark(
    group="Binary Tree",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Tree4(benchmark):
        res = benchmark(tree, 4)

@pytest.mark.benchmark(
    group="Binary Tree",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Tree8(benchmark):
        res = benchmark(tree, 8)

@pytest.mark.benchmark(
    group="Binary Tree",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Tree10(benchmark):
        res = benchmark(tree, 10)

@pytest.mark.benchmark(
    group="Binary Tree",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Tree15(benchmark):
        res = benchmark(tree, 15)

@pytest.mark.benchmark(
    group="Binary Tree",
    min_time=0.1,
    max_time=10,
    min_rounds=5,
    timer=time.time,
    disable_gc=True,
    warmup=False
)
def test_Tree20(benchmark):
        res = benchmark(tree, 20)
