n = int(input())
cnt = 0

for i in range(1, n + 1):
    if i % 2 != 0:
        if len([j for j in range(1, i + 1) if i % j == 0]) == 8:
            cnt += 1

print(cnt)
