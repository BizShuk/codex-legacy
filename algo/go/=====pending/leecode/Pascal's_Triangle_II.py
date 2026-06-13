def getRow(self, rowIndex):
    """
    :type rowIndex: int
    :rtype: List[int]
    """
    ret = [1]
    for i in xrange(0, rowIndex):
        for j in xrange(0, len(ret) - 1):
            ret[j] += ret[j + 1]
        ret.insert(0, 1)
    return ret
