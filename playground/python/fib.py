
class c():
    def __init__(self) :
        self.h = {0:0,1:1}
        self.max = 1
    def fib(self,n):
        if n in self.h: return self.h[n]
        
        for i in range(self.max+1,n+1):
            self.h[i] = self.h[i-1]+self.h[i-2]
            
        return self.h[n]
    
Fib = c()


for i in range(3):
    r = Fib.fib(int(input()))
    print(r)






