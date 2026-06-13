class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        str = str.strip()
        simbol = True if len(str) > 0 and str[0] in ["-", "+"] else False
        i = 1 if simbol == True else 0
        s = ""

        while 1:
            if i >= len(str):
                break
            c = ord(str[i])

            if (48 <= c and c <= 57):
                if s == "" and str[i] == "0":
                    pass
                else:
                    s += str[i]
            else:
                break
            i += 1

        if s == "":
            s = "0"

        s = "-" + s if simbol == True and str[0] == "-" else s
        s = int(s)
        if s > 2147483647:
            s = 2147483647
        if s < -2147483648:
            s = -2147483648
        return s
