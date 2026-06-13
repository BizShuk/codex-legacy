list_a = [1,2,3,4,5,6,7,8,"",(1,2),9,]
print(list_a,len(list_a))


# iter 需要用for , 才不會有except , need __next__
iter_list_a = iter(list_a)

for i in iter_list_a:
    print(i)

print("index method:",list_a.index(""))


try:
    print(next(iter_list_a))
except StopIteration as e:
    print("end of iter",e)
finally:
    print("here finally")


enum_list_a = enumerate(list_a)
print(list(enum_list_a))
for i,j in enum_list_a:
    print(i,j)


print("all iter_list_a",all(iter_list_a));
print("all enum_list_a",all(enum_list_a));
print("any iter_list_a",any(iter_list_a));
print("any enum_list_a",any(enum_list_a));




class a(object):

    """Docstring for a. """

    def __init__(self):
        """TODO: to be defined1. """
        print("a init")

class b(a):

    """Docstring for b. """

    def __init__(self):
        """TODO: to be defined1. """
        a.__init__(self)
        print("b init")
        self.abc=123

ins_b = b()
print("dir of ins_b:attribute",dir(ins_b))
print("dir of empty list:",dir(list()))
print("dir of local scope:",dir())

w = [1,2,3,4,5,6]
print(*w)







print(list_a)
list_a[1:3] = [4,5]
print(list_a)
list_copy = list_a.copy()
list_a[:] = [5,4,3,2,1]
print(list_a)
print(list_copy)


# set & | ^ -  len 



# namedtuple
from collections import *

def testf(a,b,c,d):
    print(a,b,c,d)
def testg(*a):
    print(*a)
    print(a[0],a[1],a[2],a[3])

testf(*(1,2,3,4))
testg(1,2,3,4)

i = 0 
while i <10:
    i+=1
    print(i)
else:
    print("out of while:",i)
