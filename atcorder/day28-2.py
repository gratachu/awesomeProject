from math import dist


def solve():
    ans = 0
    for i in range(N):
        for j in range(i):
            l = dist(XY[i], XY[j])
            ans = max(ans, l)

    return ans


N = int(input())
XY = [list(map(int, input().split())) for _ in range(N)]

print(solve())
