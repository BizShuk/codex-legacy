class Solution(object):
    def isIsomorphic(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        if len(s) != len(t):
            return False
        s_hash = {}
        t_hash = {}

        for i in xrange(0, len(s)):
            if s[i] in s_hash and s_hash[s[i]] != t[i]:
                return False
            s_hash[s[i]] = t[i]

        for i in xrange(0, len(t)):
            if t[i] in t_hash and t_hash[t[i]] != s[i]:
                return False
            t_hash[t[i]] = s[i]

        return ''.join([s_hash[s[i]] for i in xrange(0, len(s))]) == t
