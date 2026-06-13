from functools import cmp_to_key

def s(c,d):
    
    print(c,d)
    print(c['a'],d['a'])
    
    
    if c['a'] > d['a'] or (c['a'] == d['a'] and c['b'] < d['b']):
        return 1
    if c['a'] < d['a'] or (c['a']==d['a'] and c['b']>d['b']):
        return -1
    return 1

a= [{'a':3,'b':4},{'a':4,'b':3},{'a':3,'b':5}]
a= sorted(a,key=cmp_to_key(s))
print(a)
