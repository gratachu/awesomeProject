# K-th Common Divisor
# https://atcoder.jp/contests/abc120/tasks/abc120_b

a, b, k = map(int, input().split())

lst = []
for i in range(1, min(a, b) + 1):
    if a % i == 0 and b % i == 0:
        lst.append(i)
lst.reverse()

print(lst[k - 1])
