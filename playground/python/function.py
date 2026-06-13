# Fibonacci


### Closure + Nested function + callback
def A_closure(n):
    a,b=0,1
    def c():
        nonlocal a,b
        a,b = b,a+b
        print(b)
        return b
    return c

### Coroutine + Generator
def B_coroutine():
    a,b=0,1
    while 1:
        a,b = b,a+b
        print(b)
        yield b



def decorator(func):
   return func

@decorator
def some_func():
    pass
# above === below
def decorator(func):
    return func

def some_func():
    pass

some_func = decorator(some_func)






t1 = A_closure(3)
t2 = A_closure(5)

t1()
t1()
t1()
t1()
t2()
t2()



t1 = B_coroutine()
t2 = B_coroutine()
next(t1)
next(t1)
next(t1)
next(t1)
next(t1)
next(t2)
next(t2)
next(t2)
