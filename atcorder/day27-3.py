# https://atcoder.jp/contests/abc237/tasks/abc237_b

H, W = map(int, input().split())
board = []

for i in range(H):
  b = list(map(int, input().split()))
  board.append(b)

ans = []

for j in range(W):
  l = []
  for k in range(H):
    l.append(board[k][j])
  
  ans.append(l)

for a in ans:
  print(" ".join([str(_) for _ in a]))