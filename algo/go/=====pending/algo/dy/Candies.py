# Way 1 : tree-like path
# get lower peak and then go left and right by plus one
# overlap use max(from left , from right)

# Way 2
# from left to right , get go up candies
# from right to left , get go down candies


size = int(input())

x = [int(input()) for i in range(size)]
y = [0 for e in x]

i = 0
turn = -1
local_lower = []

while i+1 < size:
    if x[i] < x[i+1]:
        tmpturn =  1
    if x[i] ==x[i+1]:
        tmpturn =  0
    if x[i] > x[i+1]:
        tmpturn = -1

    if turn == -1 and tmpturn >=0:
        local_lower += [i]

    turn = tmpturn
    i+=1

if turn == -1 :
    local_lower += [i]


def goleft(x,y,i):
    while i-1 >=0 and x[i-1]>=x[i]:
        if x[i-1] > x[i]:
            y[i-1] = max(y[i]+1,y[i-1])
        if x[i-1] == x[i]:
            y[i-1] = max(y[i-1],1)
        i-=1

def goright(x,y,i,size):
    while i+1 < size and x[i] <= x[i+1]:
        if x[i+1] > x[i]:
            y[i+1] = max(y[i]+1,y[i+1])
        if x[i+1] == x[i]:
            y[i+1] = max(y[i+1],1)
        i+=1


for e in local_lower:
    y[e] = 1
    goleft(x,y,e)
    goright(x,y,e,size)

#print(y)
print(sum(y))