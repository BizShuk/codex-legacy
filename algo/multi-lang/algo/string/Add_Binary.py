class Solution(object):
    def addBinary(self, a, b):
        """
        :type a: str
        :type b: str
        :rtype: str
        """
        i, j = len(a) - 1, len(b) - 1

        ret = ""
        plus = 0
        while True:
            ta, tb = 0, 0

            if i >= 0:
                ta = int(a[i])
            if j >= 0:
                tb = int(b[j])

            total = (ta + tb + plus)
            plus = total / 2
            remain = total % 2
            ret = str(remain) + ret
            i -= 1
            j -= 1
            if i < 0 and j < 0:
                break
        while plus > 0:
            remain = plus % 2
            ret = str(remain) + ret
            plus = plus / 2

        return ret
