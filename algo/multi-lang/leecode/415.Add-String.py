class Solution(object):
    def addStrings(self, num1, num2):
        """
        :type num1: str
        :type num2: str
        :rtype: str
        """
        strmap = {i:int(i) for i in "0123456789"}
        n1l,n2l = len(num1)-1,len(num2)-1
        plus=0
        ret = ""
        
        while n1l >= 0 or n2l >=0:

            print(n1l,n2l)
            c1 = num1[n1l] if n1l >= 0 else "0"
            c2 = num2[n2l] if n2l >= 0 else "0"
            
            plus = strmap[c1] + strmap[c2] + plus
            ret = str(plus%10) +ret
            plus = plus / 10
            n1l-=1
            n2l-=1
            
        while plus!=0:
            ret = str(plus%16) + ret
            plus = int(plus/16)
        return ret
            
            
