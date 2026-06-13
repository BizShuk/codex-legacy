# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):

    def levelOrderBottom(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """

        stack = []

        if root is None:
            return []

        ret = []
        cur = root
        stack.append(cur)

        while len(stack) > 0 :

            lstack = []

            for i in range(len(stack)):
                poped = stack.pop(0)
                lstack.append(poped.val)
                if poped.left:  stack.append(poped.left)
                if poped.right: stack.append(poped.right)

            ret.insert(0,lstack)
        return ret
