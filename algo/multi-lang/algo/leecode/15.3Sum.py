class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        ret =[]
        
        ln = len(nums)
        
        for i in range(ln):
            e = nums[i]
            h = {}
            
            for e2 in nums[i+1:ln]:
                revs = -e-e2 
                reee = [e,e2,revs]
                
                if  revs in h:
                    reee.sort()
                    if reee not in ret:
                        ret += [reee]
                h[e2] = 1
                
        return ret
 

class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        ret =[]
        
        ln = len(nums)
        nums.sort()
        
        for i in range(ln):
            if i>0 and nums[i-1] == nums[i]: 
                continue
            
            e = nums[i]
            h = {}
            
            for e2 in nums[i+1:ln]:
                revs = -e-e2 
                reee = [e,revs,e2]
                
                if revs in h :
                    if reee not in ret:
                        ret += [reee]
                h[e2] = 1
                
        return ret
                
