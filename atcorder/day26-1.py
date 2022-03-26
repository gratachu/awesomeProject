S = input()
a, b = map(int, input().split())

tmp=list(S)
tmp[a-1] = S[b-1]
tmp[b-1] = S[a-1]

S = "".join(tmp)
print(S)