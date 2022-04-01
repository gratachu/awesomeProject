n, w = map(int, input().split())
cheese = [list(map(int, input().split())) for _ in range(n)]

ans = 0

cheese.sort(reverse=True)
ans = 0

for a, b in cheese:
    d = min(w, b)
    ans += d * a
    w -= d

print(ans)