a, b = map(str, input().split())

# 手法 a, bをreverseさせると桁数を気にせずに良くなり、うまくいく
a = a[::-1]
b = b[::-1]

length = min(len(a), len(b))

for i in range(length):
    if int(a[i]) + int(b[i]) > 9:
        print("Hard")
        exit()

print("Easy")
