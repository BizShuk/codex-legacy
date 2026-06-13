a = [1,2]
b = [3,4]
"""
    print how many ways to go , assume a on bottom left of b
"""

def coordinate_path(a,b):
    w = [[0,0,0] for i in range(a[0],b[0]+1)]
    w[0][0]=1
    #print(w) , use w = [[0]*(b[1]-a[1]+1)]*(b[0]-a[0]+1) got wrong answer

    for i in range(0,b[0]+1-a[0]):
        for j in range(0,b[1]+1-a[1]):

            w[i][j] += w[i-1][j] if i-1>=0 else 0
            w[i][j] += w[i][j-1] if j-1>=0 else 0

    return w[b[0]-a[0]][b[1]-a[1]]

def coordinate_path_recursive(a,b):
    if a[0]>b[0] or a[1] > b[1]:
        return 0
    if a[0]==b[0] and a[1] == b[1]:
        return 1
    count = 0

    a[0]+=1
    count += coordinate_path_recursive(a,b)
    a[0]-=1
    a[1]+=1
    count += coordinate_path_recursive(a,b)
    a[1]-=1

    return count


def coordinate_path_recursive_w(a,b,path):
    if a[0]>b[0] or a[1] > b[1]:
        return []
    if a[0]==b[0] and a[1] == b[1]:

        return [path + [[a[0],a[1]]]]

    path.append([a[0],a[1]])
    a[0]+=1
    _ret1 = coordinate_path_recursive_w(a,b,path)
    a[0]-=1
    a[1]+=1
    _ret2 = coordinate_path_recursive_w(a,b,path)
    a[1]-=1
    path.pop()
    return _ret1 + _ret2



print(coordinate_path(a,b))
print(coordinate_path_recursive(a,b))
t = coordinate_path_recursive_w(a,b,[])

for i in range(len(t)):
    print(t[i])

