n = int(input())

if n >= 42:
    n += 1
    cnt = str(n).zfill(3)
    print("AGC{cnt}".format(cnt=cnt))
else:
    cnt = str(n).zfill(3)
    print("AGC{cnt}".format(cnt=cnt))
