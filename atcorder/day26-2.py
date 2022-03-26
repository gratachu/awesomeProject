# https://atcoder.jp/contests/abc243/tasks/abc243_b

n = int(input())
a = list(map(int, input().split()))
b = list(map(int, input().split()))

contain = 0
pos_and_contain = 0

for i in a:
    if i in b and i != b[a.index(i)]:
        contain += 1

    if i == b[a.index(i)]:
        pos_and_contain += 1

print(pos_and_contain)
print(contain)
