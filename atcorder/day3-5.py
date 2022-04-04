import re


N, w = map(int, input().split())
cheese = [list(map(int, input().split())) for _ in range(N)]

cheese.sort(reverse=True)

# 美味しさ
result = 0

for i in range(len(cheese)):
    amount = cheese[i][1]
    taste = cheese[i][0]

    # 使っていい量が0の時
    if w <= 0:
        break

    # チーズの量が使っていい量より少なかった場合
    if amount <= w:
      plus_taste = amount * taste
      result += plus_taste
      w -= amount
    elif amount > w:
      plus_taste = w * taste
      result += plus_taste
      w -= w

print(result)