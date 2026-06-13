# Interview Questions
"""
    Union of n arrays with x elements. Output common members contained in at least 2 arrays. Explain the complexity of the algorithm use
"""


datas = [
        [1, 2, 3],
        [4, 5, 6],
        [3, 4, 5],
        [7, 8, 9],
        [0, 0, 0],
        [6, 2, 7]
]


def union_Narray(arrs, times):
    h, ret = {}, []

    for group_name in range(len(arrs)):
        edge = arrs[group_name]

        for node in set(edge):
            h[node] = h.get(node, []) + [group_name]

    return [k for k, v in h.items() if len(v) >= times]


result = union_Narray(datas, 2)
print(result)
