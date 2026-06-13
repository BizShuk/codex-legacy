class Solution(object):
    def compareVersion(self, version1, version2):
        """
        :type version1: str
        :type version2: str
        :rtype: int
        """
        a = version1.split(".")
        b = version2.split(".")

        i = 0
        while 1:
            if i >= len(a) and i >= len(b):
                break

            v1 = int(a[i]) if i < len(a) else 0
            v2 = int(b[i]) if i < len(b) else 0
            if v1 != v2:
                return 1 if v1 > v2 else -1

            i += 1

        return 0
