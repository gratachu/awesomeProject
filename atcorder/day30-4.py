from bisect import bisect_right


n, k = map(int, input().split())
scores = [list(map(int, input().split())) for _ in range(n)]


sumOfScores = []
for i in scores:
    sumOfScores.append(sum(i))

sumOfScores.sort(reverse=True)

for j in range(1, n + 1):
    best = sum(scores[j]) + 300
    idx = bisect_right(sumOfScores, best)
    if n - k + 1 <= idx:
        print("Yes")
    else:
        print("No")
