# Python basic

def fnA(a,b,c,d):
    pass
fnA(*[1,2,3,4])


### integer ###
`format(integer,'[x|o|d|b]')` , x for hex , o for octal , d for decimal , b for binary

### list ###
concate two list ,  `[] + []`
length , `len([])`
new , `[]`
list.sort(function) , sort with integer has problem ,  9 10 => 10 9
`[a::b]` , list start with a , and jump with b difference
`[a:b]` , list start with a and end before b
`[i for i in range(x)]`
"".join(list)
list(str)
list(tuple)


index only between  -n ~ n , -n-1 => throw error
### dictionary ###
new , {}
{i:j for i in range(x)}
dict


### string ###
import string
string.*
"".join(list)
"%s %s %.2f" %(a,b,c)
"%02d %.2f"

strip
index
replace
find
split
isalnum isalpha isdecimal isdigit islower is upper isnumeric isprintable isspace is title




### condtion ###
- 0 or "" or [] or {} == False
- 1 == True 
- "1" == False
- bool("123") == True
- bool([1]) == True
- bool([]) == False
- bool({1:1}) == True
- bool({}) == False


### build-in ###

use mostly:
- filter
- reduce
- map
- sorted(iterable,key=lambda x:(x[0],x[1],...),reverse=True) , if with different order use `-`
- input
- sum
- str
- type
- len

- enumerate
- iter with next , except StopIteration
- all , all of them are iterable
- any , any one of them are iterable
- math: abs
- chr(ascii_code) => char 
- bit manipulate , showing format(var,'[x|o|d|b]') , bit_length = len(bin(va).lstrip('-0b'))
- dir , return list of attribute of object. If no argument , return local scope attr

### operator ###
object.__add__(self, other)
object.__sub__(self, other)
object.__mul__(self, other)
object.__matmul__(self, other)
object.__truediv__(self, other)
object.__floordiv__(self, other)
object.__mod__(self, other)
object.__divmod__(self, other)
object.__pow__(self, other[, modulo])
object.__lshift__(self, other)
object.__rshift__(self, other)
object.__and__(self, other)
object.__xor__(self, other)
object.__or__(self, other)
- `//` == `%`
- `n**i` == n to power of i
- `>>`
- `<<`
- `**`
- `//`


### class ###


- `__init__` , object initializatioon
- `__doc__`  , class document
- `__str__`  , str()
- `__repr__` , repr()
- `__init__` , pythone 2 vs python 3
- `__main__` ,
- `__next__` ,

`super(parent,self).__init__()` , gerantee common parent only excute once , if no parent , it'll throw exception

```
class a(object):
    def __init__(self):
        super()
```