## shift only
# from tkinter import N


# n = input()
# elements = list(map(int, input().split()))
# count = 0

# while all(e % 2 == 0 for e in elements):
#     elements = [e / 2 for e in elements]
#     count += 1

# print(count)

## coins
## https://atcoder.jp/contests/abs/tasks/abc087_b

# A = int(input())
# B = int(input())
# C = int(input())
# s = int(input())
# count = 0

# for a in range(A + 1):
#     for b in range(B + 1):
#         for c in range(C + 1):
#             if a * 500 + b * 100 + c * 50 == s:
#                 count += 1

# print(count)

## some sums
## https://atcoder.jp/contests/abs/tasks/abc083_b

maxNum, minRange, maxRange = map(int, input().split())
elements = []

for n in range(maxNum + 1):
    s = sum(list(map(int, str(n))))
    if minRange <= s <= maxRange:
      elements.append(n)

print(sum(elements))
