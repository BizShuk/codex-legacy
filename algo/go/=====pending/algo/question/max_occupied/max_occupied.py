
data = [
    ['n1' ,4,8],
    ['n2' ,3,5],
    ['n3' ,1,5],
    ['n4' ,4,5],
    ['n5' ,1,9],
    ['n6' ,5,8],
    ['n7' ,3,7],
    ['n8' ,4,6],
    ['n9' ,5,8],
    ['n10',4,6],
    ['n11',5,6],
    ['n11',9,10],
    ['n11',9,10],
    ['n11',9,10],
    ['n11',9,10],
    ['n11',9,10],
    ['n11',9,10],
    ['n11',9,10],
    ['n11',9,10],
]



def max_occupied(data):
    s = {}
    d = {}

    allelement = set()

    for e in data:
        st,et = e[1],e[2]
        s[st] = s.get(st,0) + 1
        d[et] = d.get(et,0) + 1
        allelement.add(st)
        allelement.add(et)


    allelement = list(allelement)
    allelement.sort()

    count = 0
    ret =  []
    for e in allelement:
        count += s[e] if e in s else 0
        count -= d[e] if e in d else 0

        ret.append(count)

    max_people = max(ret)
    max_days_index = [i for i in range(len(ret)) if ret[i]==max_people]
    print(max_people)


    max_days = []
    for i in max_days_index:
        curday = allelement[i]
        nextday = allelement[i+1]

        while curday < nextday:
            max_days.append(curday)
            curday+=1

    print(max_days)
    print(ret)

    return max(ret)


max_occupied(data)



