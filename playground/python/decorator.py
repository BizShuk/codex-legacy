def a(func):
    c1,c2 = 1,2

    def b(*args):
        c3,c4 =3,4
        print(c1,c2,c3,c4)
        func(c1,c2)

    return b

@a
def d(w1,w2):
    print(w1,w2)

d(8,9)  
