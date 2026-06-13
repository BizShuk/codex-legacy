# Interview Questions
"""
    Union of n arrays with x elements. Output common members contained in at least 2 arrays. Explain the complexity of the algorithm use
"""


datas= [
        [1,3,2],
        [4,5,6],
        [3,4,5],
        [7,8,9],
        [0,0,0],
        [6,2,7]
]


def union_Narray(arrs,times):
    h,ret = {},[]
    

    for i in range(len(arrs)):
        e = arrs[i]
        
        for ie in set(e):
            h[ie] = h.get(ie,[])+ [i] 
        #print(h)

    ret = [ k for k,v in h.items() if len(v)>= times]

    print(ret)

    return ret

union_Narray(datas,2)
