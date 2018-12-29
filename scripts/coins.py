from itertools import permutations


def calc(x):
    return x[0] + x[1] * x[2] ** 2 + x[3] ** 3 - x[4]

coins = (2, 3, 5, 7, 9)

s = [i for i in permutations(coins) if calc(i) == 399]
print(s[0])
