N = int(input())
H = list(map(int, input().split()))

# 配列を初期化
dp = [None] * N

# 1つ目の足場の時は必要な体力は0
dp[0] = 0

# 2つ目以降を配列に入れる
for i in range(1, N):
    # 2つ目の足場の時は1つ目から移動するしか方法がないから単純に計算するだけ
    if i == 1:
        dp[i] = abs(H[i - 1] - H[i])
    if i >= 2:
        # 1つ前の足場から移動するパターン
        v1 = dp[i - 1] + abs(H[i - 1] - H[i])

        # 2つ前の足場から移動するパターン
        v2 = dp[i - 2] + abs(H[i - 2] - H[i])

        # 消費体力の少ない方なので少ない方をとる
        dp[i] = min(v1, v2)

print(dp[N - 1])
