class Solution(object):
    def wordPattern(self, pattern, str):
        """
        :type pattern: str
        :type str: str
        :rtype: bool
        """
        str_arr = str.split(" ")
        if len(str_arr) != len(pattern):
            return False
        p_hash = {}
        t_hash = {}
        for i in xrange(0, len(pattern)):
            if pattern[i] not in p_hash:
                p_hash[pattern[i]] = str_arr[i]
            if str_arr[i] in t_hash and t_hash[str_arr[i]] != pattern[i]:
                return False
            t_hash[str_arr[i]] = pattern[i]
        a = ' '.join([p_hash[pattern[i]] for i in xrange(0, len(pattern))])

        return ' '.join(
            [p_hash[pattern[i]] for i in xrange(0, len(pattern))]) == str
