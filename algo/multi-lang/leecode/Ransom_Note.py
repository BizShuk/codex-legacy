class Solution(object):
    def canConstruct(self, ransomNote, magazine):
        """
        :type ransomNote: str
        :type magazine: str
        :rtype: bool
        """

        str_dict = {}
        for c in ransomNote:
            if c in str_dict:
                str_dict[c] += 1
            else:
                str_dict[c] = 1

        for c in magazine:
            if c in str_dict:
                str_dict[c] -= 1

        for key in str_dict:
            if str_dict[key] > 0:
                return False
        return True
