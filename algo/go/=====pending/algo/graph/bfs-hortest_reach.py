def findpath(N,h,start):
    ret = [-1 for i in range(N) ]
    ret[start] =0

    stack = [start]
    while len(stack)>0:
        c = stack.pop(0)
        for e in h[c]:
            if ret[e] == -1:
                stack.append(e)
                ret[e] = ret[c]+6
    return ret[0:start]+ret[start+1:]

cases = int(input())

for case in range(cases):
    N,E= list(map(int,input().split()))

    h = {i:[] for i in range(N)}
    for i in range(E):
        e = list(map(int,input().split()))
        h[e[0]-1].append(e[1]-1)
        h[e[1]-1].append(e[0]-1)

    start = int(input())-1
    w = findpath(N,h,start)
    print(*w)




