def solve():
    N, X = map(int, input().split())
    pre = {0}

    for _ in range(N):
        a, b = map(int, input().split())
        cur = set()
        for x in pre:
            cur.add(x + a)
            cur.add(x + b)
        pre = cur
    return X in cur


print("Yes" if solve() else "No")

