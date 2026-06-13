import  random

class node(object):

    """Docstring for node. """

    def __init__(self,n):
        """TODO: to be defined1. """
        self.val = n #random.randrange(0,100,1)
        self.left = None
        self.right = None

    def getLevel(self,n):
        if n == 0:
            return [self]
        
        left = self.left.getLevel(n-1) if self.left else []
        right = self.right.getLevel(n-1) if self.right else []

        return left+right
    
    def __str__(self):
        return self.val,'L' if self.left else '_','R' if self.right else '_'

root = node(0)
n1 = node(1)
n2 = node(2)
n3 = node(3)
n4 = node(4)
n5 = node(5)
n6 = node(6)

root.left  = n1
root.right = n2
n1.left  = n3
n1.right = n4
n2.left  = n5
n2.right = n6


v = root.getLevel(0)
print(v)
v = root.getLevel(1)
print(v)
v = root.getLevel(2)
print(v)

print(root.__doc__)
