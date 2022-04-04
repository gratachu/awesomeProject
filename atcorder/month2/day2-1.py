import collections


n = int(input())
lst = list(map(int, input().split()))

# 100 + 400のパターンと200 + 300のパターンを計算する
counter = collections.Counter(lst)

a = counter[100] * counter[400]
b = counter[200] * counter[300]

print(a + b)
