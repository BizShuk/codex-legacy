from __future__ import print_function

data = [[1,4,7,9],[10,2,5,8],[12,11,3,6]]
"""
    output:
        1 2 3
        4 5 6
        7 8
        9
        10 11
        12
"""
print(data)



rows = len(data)
cols = len(data[0]) if rows>0 else 0

if not rows or not cols:
    exit

for i in range(0,rows+cols-1):
    _i,_j = 0,i
    if i >= cols:
        _i,_j = i-cols+1,0
    def pslanting(data,i,j,rows,cols):
        _i,_j = i,j
        while _i < rows and _j < cols:
            yield data[_i][_j]
            _i += 1
            _j += 1


    for e in pslanting(data,_i,_j,rows,cols):
        print(e," ", end="")


    print("\n")

