# Creates a binary tree of depth 'd'
def make_tree(d):

    if d > 0:
        d -= 1
        return (make_tree(d), make_tree(d))
    return (None, None)


def check_tree(node):

    (l, r) = node
    if l is None:
        return 1
    else:
        return 1 + check_tree(l) + check_tree(r)


def tree(n, min_depth=4):

        if (n < 0):
                return 0

        max_depth = max(min_depth + 2, n)
        stretch_depth = max_depth + 1

        t = make_tree(stretch_depth)
        check_tree(t)

        long_lived_tree = make_tree(max_depth)

        mmd = max_depth + min_depth
        for d in range(min_depth, stretch_depth, 2):
                iterations = 2 ** (mmd - d)
                cs = 0
                for i in range(iterations):
                        cs = cs + check_tree(make_tree(d))

        check_tree(long_lived_tree)

        return 0
