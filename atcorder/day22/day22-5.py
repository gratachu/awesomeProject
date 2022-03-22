# Half and Half
# https://atcoder.jp/contests/abc095/tasks/arc096_a

a, b, c, x, y = map(int, input().split())

if a + b <= c * 2:
    print(a * x + b * y)
else:
    if x <= y:
        print(min(2 * c * y, 2 * c * x + b * (y - x)))
    else:
        print(min(2 * c * x, 2 * c * y + a * (x - y)))
