# times = list(map(int, input().split()))

# t_time = [times[0], times[1]]
# a_time = [times[2], times[3]]

# # 時間判定
# if t_time[0] > a_time[0]:
#     print("Aoki")
# # 時間が一緒
# elif t_time[0] == a_time[0]:
#     if t_time[1] == a_time[1] or t_time[1] < a_time[1]:
#         print("Takahashi")
#     else:
#         print("Aoki")
# # その他
# else:
#   print("Aoki")


N = int(input())
lst = list(map(int, input().split()))

result = 0

for i in range(2001):
    if i in lst:
        continue

    result = i
    break

print(result)

