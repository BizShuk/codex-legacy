class Solution(object):
    def findTheDifference(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: str
        """
        ret = 0
        
        for i in s : ret ^= ord(i)
        for i in t : ret ^= ord(i)
        
        return chr(ret)
        
