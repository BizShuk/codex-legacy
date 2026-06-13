class Solution(object):
    def trailingZeroes(self, n):
        """
        :type n: int
        :rtype: int
        """

        ret = 0

        i=1
        while True:
            division = 5**i
            if division > n :
                break;
            ret += int(n/division)
            i+=1

        return ret
