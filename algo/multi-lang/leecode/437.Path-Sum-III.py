# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def pathSum(self, root, sum):
        """
        :type root: TreeNode
        :type sum: int
        :rtype: int
        """
        
        ret = 0
        
        if not root :
            return 0
            
        if sum ==  root.val:
            ret += 1
        
        ret += self.pathSum(root.left,sum) if root.left else  0
        ret += self.pathSum(root.right,sum) if root.right else  0
        
        ret += self.findPath(root.left,sum-root.val) if root.left else  0
        ret += self.findPath(root.right,sum-root.val) if root.right else  0
        
        return ret 
        
    def findPath(self,root,sum):
        
        ret = 0
        if root.val == sum:
            ret += 1
            
        if root.left is None and root.right is None:
            return ret
        
        ret += self.findPath(root.right,sum-root.val) if root.right else  0    
        ret += self.findPath(root.left ,sum-root.val) if root.left else  0
        
        return ret
        
        
        
        
        
        
        
        
        
        
        
        
        
        
