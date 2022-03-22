formula = input()
digits = formula.split("+")

numbers = []

for i, d in enumerate(digits):
  n = d.count("<") * 10 + d.count("/")
  numbers.append(n)

print(sum(numbers))