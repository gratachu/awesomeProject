# https://atcoder.jp/contests/abc243/tasks/abc243_a
from tkinter import E


V, A, B, C = map(int, input().split())

while True:
    if V < A:
        print("F")
        exit()
    V -= A

    if V < B:
        print("M")
        exit()
    V -= B

    if V < C:
        print("T")
        exit()
    V -= C
