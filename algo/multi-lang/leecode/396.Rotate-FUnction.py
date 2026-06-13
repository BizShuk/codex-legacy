class Solution(object):
    def maxRotateFunction(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        totalSum = sum(A)
        Al = len(A)
        
        F = {0:0}
        
        for i in range(Al):
            F[0] += i * A[i]
        
        maxNum = F[0]
            
        for i in range(Al-1,0,-1):
            F[Al-i] = F[Al-i-1] + totalSum - Al * A[i]
            maxNum = max(maxNum,F[Al-i])
            
        return maxNum

    # TLE
    def maxRotateFunction_TLE(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        
        Rsum = {}
        rotated = 0
        max = None
        l = len(A)
        
        for i in range(l):
            mmax = None
            for j in range(l):
                if j in Rsum:
                    Rsum[j] += i * A[rotated%l]
                else:
                    Rsum[j]  = i * A[rotated%l]
                
                if mmax is None or mmax < Rsum[j]:
                    mmax = Rsum[j]
                
                rotated -= 1

            max = mmax
            rotated += 1
            
        return 0 if max is None else max
