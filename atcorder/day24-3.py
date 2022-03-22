# # 素数判定法
# def isPrime(num):
#     limit = int(num**0.5) + 1  # numの平方根を出してそれが割り切れなければ素数という手法をとる
#     for i in range(2, limit):
#         if num % i == 0:
#             return False
#     return True


# # 約数列挙
# def allDivisor(num):
#     divisors = []
#     limit = int(num**0.5) + 1
#     for i in range(2, limit):
#         if num % i == 0:
#             divisors.append(i)  # 割った数を追加
#             divisors.append(num // i)  # 商を追加

#     return list(set(divisors))


# N = 286
# divisors = allDivisor(N)

# result = []
# for d in divisors:
#     if isPrime(d):
#         result.append(d)

# print(result)

# 入力
N = 286

# 素因数分解
Answer = []
LIMIT = int(N**0.5)
for i in range(2, LIMIT + 1):
    while N % i == 0:
        N /= i
        Answer.append(i)

if N >= 2:
    Answer.append(int(N))

# 出力
print(*Answer)
