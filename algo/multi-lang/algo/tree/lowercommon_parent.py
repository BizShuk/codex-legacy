import tree

btrees = [
    [1,2,3,4,5,6,7],
    [7,6,5,4,3,2,1],
    [1,2,3],
    [9,7,5,3,1,2,4,6,8]
]

sNodes = [
    [2,5],
    [1,7],
    [2,3],
    [8,6]
]


# lowerest common parent
def findLCP(root,n1,n2):
    if root is None:
        return None
    
    if root.value == n1 or root.value == n2:
        return root

    else:
        l = findLCP(root.left ,n1,n2)
        r = findLCP(root.right,n1,n2)

        if l and r:
            return root
        
        if l:
            return l
        else:
            return r

        

for i in range(len(btrees)):
    print(btrees[i],sNodes[i])
    root = tree.createBtree(btrees[i])
    w= findLCP(root,sNodes[i][0],sNodes[i][1])
    print(w.value)
