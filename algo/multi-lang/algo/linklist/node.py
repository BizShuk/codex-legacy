class Node(object):

    """Docstring for Node. """

    def __init__(self, val, next=None):
        self.next = next if next is not None else 0
        self.val = val

    def isCircle(self):
        pass

    def length(self):
        pass
