class MinStack(object):
    def __init__(self):
        """
        initialize your data structure here.
        """

        self.min = []
        self.stack = []

    def push(self, x):
        """
        :type x: int
        :rtype: void
        """

        self.stack.append(x)
        if len(self.min) == 0 or self.min[len(self.min) - 1] >= x:
            self.min.append(x)

    def pop(self):
        """
        :rtype: void
        """

        p = self.stack.pop()
        if p == self.min[len(self.min) - 1]:
            self.min.pop()

    def top(self):
        """
        :rtype: int
        """
        return self.stack[len(self.stack) - 1]

    def getMin(self):
        """
        :rtype: int
        """
        return self.min[len(self.min) - 1] if len(self.min) > 0 else None

# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()
