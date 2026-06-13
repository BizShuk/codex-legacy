class Solution(object):
    def readBinaryWatch(self, num):
        """
        :type num: int
        :rtype: List[str]
        """

        return [ "%d:%02d"%(i,j) 
                for i in range(12) for j in range(60)
                if (bin(i)+bin(j)).count('1') == num ]
