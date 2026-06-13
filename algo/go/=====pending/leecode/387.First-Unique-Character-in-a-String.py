class Solution(object):
    def firstUniqChar(self, s):
        """
        :type s: str
        :rtype: int
        """
        
        
        sl = len(s)
        min = None
        for c in string.ascii_lowercase:
            
            if s.count(c) == 1 and (s.find(c)<min or min is None):
                min = s.find(c)
                
                
                
        return min if min is not None else -1
            
