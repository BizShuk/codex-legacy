class Solution(object):
    def toHex(self, num):
        """
        :type num: int
        :rtype: str
        """
        ret = ""
        if num == 0: return "0"
            
        hmap = [i for i in '0123456789abcdef']
        
        while num != 0 and len(ret) < 8:    # num ==0 for positive integer , len(ret) for negative integer
            a = num & 15
            num = num >>4
            ret = hmap[a] + ret

        return ret
                
