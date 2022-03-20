# uneven number
# https://atcoder.jp/contests/abc136/tasks/abc136_b

n = int(input())
cnt = 0 

for i in range(1, n+1):
  print(str(i))
  if not len(str(i)) % 2 == 0:
    cnt += 1


print(cnt)