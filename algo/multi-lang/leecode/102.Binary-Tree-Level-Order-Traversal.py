# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def levelOrder(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """

        stack,ret = [],[]
        cur = root
        stack.append(cur)

        while cur and len(stack)>0:

            lstack = []
            for i in range(len(stack)):
                cur = stack.pop(0)
                lstack.append(cur.val)
                if cur.left:  stack.append(cur.left)
                if cur.right: stack.append(cur.right)

            ret.append(lstack)

        return ret
