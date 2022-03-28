S = input()
# 先頭のaの連続数を数える
l = len(S) - len(S.lstrip("a"))
# 末尾から数えたaの連続数を数える
r = len(S) - len(S.rstrip("a"))

# 末尾より先頭の方がaが多かったらaを足しても回文にできない
if l > r:
    print("No")
    exit()

# 差分のaを足した文字列を作る
T = "a" * (r - l) + S

if T == T[::-1]:
    print("Yes")
else:
    print("No")
