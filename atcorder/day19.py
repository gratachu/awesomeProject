# a = int(input())
# b, c = map(int, input().split())
# s = input()

# print("{} {}".format(a+b+c, s))

#####

# a, b = map(int, input().split())
# p = a * b
# if p % 2 == 0:
#   print("Even")
# else:
#   print("Odd")

#####
strs = list(input())
ball = 0

for s in strs:
  if s == "1":
    ball += 1

print(ball)


