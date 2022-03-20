## daydream
## https://atcoder.jp/contests/abs/tasks/arc065_a

text = input()

text = text.replace("eraser", "")
text = text.replace("erase", "")
text = text.replace("dreamer", "")
text = text.replace("dream", "")

if len(text) == 0:
    print("YES")
else:
    print("NO")
