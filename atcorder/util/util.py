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