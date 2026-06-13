class Solution(object):
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        interval = numRows if numRows <= 2 else numRows * 2 - 2
        len_s = len(s)
        ret = ""

        for i in xrange(0, numRows):
            peer = False
            if i % interval == 0 or i % interval == numRows - 1:
                peer = True

            tmp_i = i
            zip_num = interval - (i % interval) - i

            while True:
                if tmp_i >= len_s:
                    break
                ret += s[tmp_i]

                if tmp_i + zip_num >= len_s:
                    break
                if not peer:
                    ret += s[tmp_i + zip_num]

                tmp_i += interval

        return ret
