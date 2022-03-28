N = int(input())
H = list(map(int, input().split()))

pos = H[0]

for i in range(1, N):
    if pos < H[i]:
        pos = H[i]
    else:
        break

print(pos)
