from __future__ import print_function
class Tree(object):
    """docstring for tree"""
    def __init__(self, root):
        super(tree, self).__init__()
        self.root = root


class Node(object):
    """docstring for node"""
    def __init__(self, value):
        super(Node, self).__init__()
        self.value = value
        self.left = None
        self.right= None

    def inOrder(self):
        def findNext(node):
            stack = []
            while node :

                stack.append(node)
                if node.left:
                    node = node.left
                else:
                    node = node.right

            return stack

        stack = findNext(self)
        while len(stack) > 0:
            c = stack.pop()
            print(c.value,end="")
            if len(stack) > 0:
                c = stack.pop()
                print(c.value,end="")
                if c.right:
                    stack += findNext(c.right)
        print("")

    def preOrder(self):

        stack = [self]
        while len(stack) > 0:
            node = stack.pop()
            print(node.value,end="")

            if node.right:
                stack += [node.right]

            if node.left:
                stack += [node.left]
        print("\n")

    def postOrder(self):
        def findNext(node):
            stack = []
            while node :

                stack.append(node)
                if node.left:
                    node = node.left
                else:
                    node = node.right

            return stack

        stack = findNext(self)

        while len(stack)>0:
            c = stack.pop()
            print(c.value,end="")
            if len(stack) > 0:
                top = stack[len(stack)-1]
                if c == top.left:
                    stack += findNext(top.right)
        print("")



    def toString(self):
        queue = [self]

        while len(queue) > 0:
            c,level_queue = None,[]

            while True:
                if len(queue)==0:
                    queue += level_queue
                    break
                c = queue.pop(0)
                print(c.value,end="")
                if c.left:
                    queue.append(c.left)
                if c.right:
                    queue.append(c.right)

#t = createBtree([1,2,3,4,5,6,7,8,9])
def createBtree(a):
    node_list = [None]
    if len(a)==0:
        return node_list

    for e in a:
        node_list.append(Node(e))

    len_a = len(a)+1
    for i in range(1,len_a):

        node_list[i].left = node_list[i*2]  if i*2 < len_a else None
        node_list[i].right= node_list[i*2+1]  if i*2+1 < len_a else None

    return node_list[1]





















