import collections


N = int(input())
A = list(map(int, input().split()))

c = collections.Counter(A)

print(c.most_common()[-1][0])
