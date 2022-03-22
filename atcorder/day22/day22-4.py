# Digits in Multiplication
# https://atcoder.jp/contests/abc057/tasks/abc057_c

n = int(input())
result = 10

for a in range(1, int(n**0.5) + 1):
    if n % a != 0:
        continue

    b = n // a

    lenA = len(str(a))
    lenB = len(str(b))

    if max(lenA, lenB) < result:
        result = max(lenA, lenB)

print(result)
