# 最大公約数の求め方
# ユークリッドの互除法

# 2つの値のGCD
def calculateGCD(num1, num2):
    while num1 >= 1 and num2 >= 1:
        if num1 > num2:
            num1 = num1 % num2
        else:
            num2 = num2 % num1

    if num1 >= 1:
        return num1
    else:
        return num2


# 値がリストで与えられた場合のGCD
def calculateGCDFromList(lst):
    gcd = 0
    for i in range(len(lst) - 1):
        if i == 0:
            a = lst[i]
            b = lst[i + 1]
            while a >= 1 and b >= 1:
                if a > b:
                    a = a % b
                else:
                    b = b % a
        else:
            a = lst[i + 1]
            b = gcd
            while a >= 1 and b >= 1:
                if a > b:
                    a = a % b
                else:
                    b = b % a

        if a >= 1:
            gcd = a
        else:
            gcd = b

    return gcd


n = int(input())
lst = list(map(int, input().split()))

print(calculateGCDFromList(lst))
