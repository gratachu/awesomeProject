# https://atcoder.jp/contests/abc242/tasks/abc242_a
a, b, c, x = map(int, input().split())

if a>=x:
  print(float(1))
if b >= x > a:
  print(float(c) / float(b-a))
if b < x:
  print(float(0))