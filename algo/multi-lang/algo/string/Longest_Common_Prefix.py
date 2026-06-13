def longestCommonPrefix(self, strs):
    """
    :type strs: List[str]
    :rtype: str
    """
    if len(strs) == 0:
        return ""
    size = len(strs)
    lens = []
    for i in xrange(0, size):
        lens.append(len(strs[i]))

    i = 0
    lprefix = ""
    while 1:
        thischar = None
        for j in xrange(0, size):
            if i >= lens[j]:
                return lprefix
            if thischar is None or thischar == strs[j][i]:
                thischar = strs[j][i]

            if i >= lens[j] or strs[j][i] != thischar:
                return lprefix

        lprefix += thischar if thischar is not None else ""
        i += 1

    return lprefix
