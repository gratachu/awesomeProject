s1 = input()
s2 = input()

s1_cnt = s1.count("#")
s2_cnt = s2.count("#")

if s1_cnt == 1 and s2_cnt == 1:
    if s1[0] == "#" and s2[1] == "#":
        print("No")
    elif s1[1] == "#" and s2[0] == "#":
        print("No")
    else:
        print("Yes")
else:
    print("Yes")
