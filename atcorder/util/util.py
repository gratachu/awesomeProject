# 使いそうなメソッドメモ

## 桁数を求める
n = 123
print(len(str(n)))

## 約数を列挙する
divisors = [i for i in range(1, n + 1) if n % i == 0]
# 約数の個数を求める
divis = len([i for i in range(1, n + 1) if n % i == 0])
