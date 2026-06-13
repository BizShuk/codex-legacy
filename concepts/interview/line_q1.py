import sys

class Fibonacci(object):

    """Docstring for Fibonacci. """

    def __init__(self):
        """TODO: to be defined1. """
        self.result = {
                0:0,
                1:1
        }

    def f(self,n):
        if n in self.result:
            return self.result[n]

        i = 0 
        while i <= n:

            if i <= 1 or n in self.result:
                i+=1
                continue
            
            self.result[i] = self.result[i-1] + self.result[i-2]
            i+=1

        return self.result[n]

    # maximum recursion depth exceeded
    def f2(self,n):
        if n in self.result:
            return self.result[n]
        return self.f2(n-1) + self.f2(n-2)


F = Fibonacci()




result = F.f(8181)
print(result)

for i in range(100):
    result = F.f(i)
    print(result)

print(sys.getrecursionlimit())
