class Solution(object):
    def repeatedSubstringPattern(self, str):
        """
        :type str: str
        :rtype: bool
        """
        
        repeat = ""
        sl = len(str)
        
        suc = False
        
        for i in range(sl):
            
            repeat += str[i]
            
            rl = len(repeat)
            for j in range(i+1,sl,rl):
                if repeat != str[j:j+rl]:
                    suc = False
                    break
                suc = True
            
            if suc == True:
                break
            
        return suc
        
        
        
        
