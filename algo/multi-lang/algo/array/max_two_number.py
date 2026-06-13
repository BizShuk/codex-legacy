data = [
    [1,2,3,4,5,6,7,8,0],
    [],
    [1],
    [2,2],
    [3,2,2]
]

def max_two_number(data):
    a1 = None
    a2 = None
    for e in data:
        if a1 is None or a1 < e:
            a2 = a1
            a1 = e
            continue
        if a2 is None or a2 < e:
            a2 = e
            continue
    return a2 or a1

for e in data:
    print(max_two_number(e))

