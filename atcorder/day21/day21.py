## card game for two
## https://atcoder.jp/contests/abs/tasks/abc088_b
n = int(input())
cards = list(map(int, input().split()))
cards.sort(reverse=True)

alice = 0
bob = 0

for i in range(len(cards)):
    if i % 2 == 0:
        alice += cards[i]
    else:
        bob += cards[i]

print(alice - bob)
