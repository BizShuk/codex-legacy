class Solution(object):

    def countAndSay(self, n):
        """
            :type n: int
            :rtype: str
        """
            if n <= 1:
                return '1'

            ret = ""
            last_str = self.countAndSay(n - 1)
            print n - 1, last_str
            i, end, count, last_word = 0, len(last_str), 0, ''
            while i <= end:
                if i == end:
                    break

                if last_word == last_str[i]:
                    count += 1
                else:
                    if last_word != '':
                        ret += str(count) + last_word
                        last_word = last_str[i]
                        count = 1
                        i += 1

                        if last_word != '':
                            ret += str(count) + last_word
                            return ret
