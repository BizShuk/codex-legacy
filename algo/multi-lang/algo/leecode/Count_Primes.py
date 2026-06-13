class Solution(object):
    def countPrimes(self, n):
        """
        :type n: int
        :rtype: int
        """

        x = [1] * n
        count = 0
        for i in xrange(0, n - 1):
            if i == 0 or x[i] == 0:
                continue

            count += 1
            cur = i + 1
            j = i + cur
            while 1:
                if j >= n - 1:
                    break

                x[j] = 0
                j += cur

        return count
