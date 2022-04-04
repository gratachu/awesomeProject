n, a, b = map(int, input().split())
p, q, r, s = map(int, input().split())

height = q - p + 1
width = s - r + 1

for i in range(p, q + 1):
    ans = []
    for j in range(r, s + 1):
        if i - j == a - b or i + j == a + b:
            ans.append("#")
        else:
            ans.append(".")

    print(*ans, sep="")
