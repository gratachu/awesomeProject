# 使いそうなメソッドメモ

## 桁数を求める
n = 123
print(len(str(n)))

## 約数を列挙する
divisors = [i for i in range(1, n + 1) if n % i == 0]
### 約数の個数を求める
divis = len([i for i in range(1, n + 1) if n % i == 0])


## 全探索で範囲を超えそうな時に平方根で範囲を半分に絞る
for i in range(1, int(n**0.5) + 1):
    print(i)

## 当てはまらない条件の時にはfor文をskipしたい時
for i in range(n):
    # nがiで割り切れない時は飛ばしたい
    if n % i != 0:
        continue

## for文でiを3桁にゼロ埋めをしてリストにする
for i in range(1000):
    num = [c for c in f"{i:03}"]

## 文字列を任意の桁数でゼロ埋めする 例は3桁
str(n).zfill(3)

# 2次元配列をワンライナーで作る
list_in_list = [list(map(int, input().split()))[1:] for _ in range(10)]

# 最小がわからない時の対処法 => いったん無限代数で置き換えておく
ans = float("inf")

# bit全探索基本
# 入力
N, S = map(int, input().split())
A = list(map(int, input().split()))

# 全パターンを探索：(1 << N) は 2 の N 乗
answer = "No"
for i in range(0, 1 << N):
    partsum = 0
    for j in range(0, N):
        # (i & (1 << j)) != 0LL の場合、i の 2 進法表記の下から j+1 桁目が 1
        # (1 << j) は Python では「2 の j 乗」を意味します
        if (i & (1 << j)) != 0:
            partsum += A[j]
    if partsum == S:
        answer = "Yes"
        break

# 出力
print(answer)
