s = input()

ans = 0

# 与えられる文字数の間に入る/入らないなので
# 2の文字数-1乗を計算しなければならない
for i in range(2 ** (len(s) - 1)):
    tmp = 0
    for j in range(len(s)):
        if i >> j & 1:
            ans += tmp * 10 + int(s[j])
            tmp = 0
        else:
            tmp = tmp * 10 + int(s[j])
    ans += tmp

print(ans)