class Solution(object):

    def arrangeCoins(self, n):
        """
        :type n: int
        :rtype: int
        """

        i = 0
        add = 1

        while i+add <n:
            i += add
            add+=1

        return add-1 if i+add != n else add
