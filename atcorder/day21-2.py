## kagami mochi
## https://atcoder.jp/contests/abs/tasks/abc085_b

n = int(input())
lst = []

for i in range(n):
    e = int(input())
    lst.append(e)

print(len(set(lst)))
