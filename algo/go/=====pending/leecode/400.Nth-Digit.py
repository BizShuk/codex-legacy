class Solution(object):
    def findNthDigit(self, n):
        """
        :type n: int
        :rtype: int
        """
        
        i = 0
        while True:
            c = 9 * (10**i) * (i+1)
            if n - c < 0 :
                
                break
            n -= c
            i+=1
        
        newNum = 10**i  -1 
        newNum += int((n-1)/(i+1))+1   #  該數有i+1個digits
        reverseCount = i - (n-1)%(i+1)
        
        for j in range(reverseCount):
            newNum = int(newNum/10)
        return newNum%10
