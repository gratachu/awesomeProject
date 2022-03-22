number = input()

# 加減算の結果は7になればいい
expect_result = 7

for i in range(2 ** (len(number) - 1)):
    tmp = int(number[0])
    ans = number[0]

    for j in range(len(number) - 1):
        if (i >> j) & 1:
            tmp -= int(number[j + 1])
            ans += "-"
        else:
            tmp += int(number[j + 1])
            ans += "+"

        ans += number[j + 1]

    if tmp == expect_result:
        break

print(ans + "=7")
