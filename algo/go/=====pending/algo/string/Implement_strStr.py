def strStr(haystack, needle):
    """
        :type haystack: str
        :type needle: str
        :rtype: int
    """
    if needle == "":
        return 0
    kmp = []
    ki = 0
    for i in xrange(0, len(needle)):
        if needle[i] != needle[ki]:
            ki = 0

        if i != 0 and needle[i] == needle[ki]:
            ki += 1

        kmp.append(ki)

    i = 0
    j = 0
    print haystack
    print needle
    print kmp

    while 1:
        print i, j

        if j >= len(needle):
            return i - j

        if i >= len(haystack):
            break

        if haystack[i] != needle[j] and j > 1:
            j = kmp[j - 1]

        if haystack[i] == needle[j]:
            j += 1

        i += 1
    return -1


print strStr("a", "a")
print strStr("mississippi", "issip")
