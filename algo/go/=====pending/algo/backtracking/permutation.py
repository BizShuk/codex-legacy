import itertools  
import itertools  


datas = [1,2,3,4,5,6,7]

times = [1,5,8,9,2,3,6,3]

a = itertools.permutations(datas,4)

max_num = None
for a,b,c,d in a:
    h = 10*a +b
    m = 10*c +d
    if h >24 or m > 60 or (h == 24 and m > 0) :
            continue
    if max_num is None or max_num < 100*h+m:
        max_num = 100*h+m

print(max_num)
