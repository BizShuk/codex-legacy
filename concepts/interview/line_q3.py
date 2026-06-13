def dicesteps(n):
    if n < 0:
        return 0
    if n == 0 or n == 1:
        return 1
    return dicesteps(n-1) + dicesteps(n-2) + dicesteps(n-3) + dicesteps(n-4) + dicesteps(n-5) + dicesteps(n-6)


def dicesteps2(n):
    if n <= 0:
        return 0
    ret = {0:1,-1:0,-2:0,-3:0,-4:0,-5:0,-6:0}

    for i in range(1,n+1):
        ret[i] = ret[i-1] + ret[i-2] + ret[i-3] + ret[i-4] + ret[i-5] + ret[i-6] 



    return ret[n]

ret = dicesteps2(1)
print(ret,'\n')

ret = dicesteps2(2)
print(ret,'\n')

ret = dicesteps2(3)
print(ret,'\n')

ret = dicesteps2(4)
print(ret,'\n')

ret = dicesteps2(5)
print(ret,'\n')

ret = dicesteps2(6)
print(ret,'\n')

ret = dicesteps2(610)
print(ret,'\n')

