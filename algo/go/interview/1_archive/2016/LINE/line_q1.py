import sys


# Find out Nth fibonacci number
class Fibonacci(object):

    def __init__(self):
        self.result = {
            0: 0,
            1: 1
        }
    # [Solution]: Incremental approach with cache for next time.

    def fib_incre_cache(self, n):
        if n in self.result:
            return self.result[n]

        i = 0
        while i <= n:
            if i <= 1 or n in self.result:
                i += 1
                continue

            self.result[i] = self.result[i-1] + self.result[i-2]
            i += 1

        return self.result[n]

    # [Solution]: [Time exceeding] Recursive way from top to bottom.
    def fib_recur(self, n):
        if n in self.result:
            return self.result[n]
        return self.f2(n-1) + self.f2(n-2)


F = Fibonacci()


result = F.fib_incre_cache(8181)
print(result)

for i in range(100):
    result = F.fib_incre_cache(i)
    print(result)

print(sys.getrecursionlimit())
