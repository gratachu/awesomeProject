s = input()

cnt = 0
ans = 0
a = ["A", "C", "G", "T"]
for i in s:
    if i in a:
        cnt += 1
        ans = max(ans, cnt)
    else:
        cnt = 0

print(ans)
