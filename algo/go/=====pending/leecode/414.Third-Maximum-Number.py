class Solution(object):
    def thirdMax(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        a = None
        b = None
        c = None
        
        
        for i in range(len(nums)):
            s = nums[i]
            
            if s in [a,b,c]: continue
        
            if a is None or a < s:
                tmp = a 
                a = s
                s = tmp
                
            if b is None or b < s:
                tmp = b
                b = s
                s = tmp
                
            if c is None or c < s:
                tmp = c
                c = s
                s = tmp
        
        
        if c is not None : return c
        if a is not None: return a
        if b is not None: return b
                
                
