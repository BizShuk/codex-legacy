from functools import reduce
from random import random

def max(visits, time):
    return reduce(lambda num, t: num - 1 if time > t else num, visits[1],
        reduce(lambda num, t: num + 1 if time > t else num, visits[0] , 0))
        
GUESTS = 30
starts = [int(random() * 24) for i in range(GUESTS)]
ends = [int(random() * (24 - start) + start) for start in starts]
visits = [sorted(starts), sorted(ends)]


print(visits)

for t in range(24):
    num = max(visits, t)
    if num != 0:
        print("%2d 時訪客：%2d 位" % (t, num))
