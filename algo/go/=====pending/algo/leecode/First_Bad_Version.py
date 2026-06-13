# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):


class Solution(object):
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """

        min, max = 1, n
        while True:
            mean = (min + max) / 2
            if min == max:
                return mean
            if isBadVersion(mean):
                max = mean
            else:
                min = mean + 1
