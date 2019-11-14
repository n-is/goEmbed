def fib(n):

        if n < 0:
                return -1

        if n == 0:
                return 0
        elif n == 1:
                return 1

        n0, n1 = 0, 1

        for _ in range(2, n+1):
                tmp = n0 + n1
                n0 = n1
                n1 = tmp

        return n1

