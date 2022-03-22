# 移動階数計算
def calFloors(lst):
  start_floor = 1
  ans = 0
  for f in lst:
    m = abs(f - start_floor)
    ans += m
    start_floor = f
  return ans

N = int(input()) # ログ
logs = []

for i in range(N):
  n = int(input())
  logs.append(n)

print(calFloors(logs))
