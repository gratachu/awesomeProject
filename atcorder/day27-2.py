# https://atcoder.jp/contests/abc238/tasks/abc238_b

N = int(input())
A = list(map(int, input().split()))

cut = [0]

angle = 0

for i in range(N):
  angle += A[i]
  cut.append(angle)

for i in range(N+1):
  cut[i] %= 360

cut.sort()

cut.append(360)

ans = 0

for i in range(1, N+2):
  ans = max(ans, cut[i] - cut[i - 1])
  
print(ans)