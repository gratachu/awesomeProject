N = int(input())
S = input()

result = 0

for i in range(1000):
    sec_num = [c for c in f"{i:03}"]

    tmp = 0
    for c in S:
        if c == sec_num[tmp]:
            tmp += 1
            if tmp == 3:
                result += 1
                break
print(result)