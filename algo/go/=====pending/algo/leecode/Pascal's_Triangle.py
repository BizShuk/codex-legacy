def generate(self, numRows):
    """
    :type numRows: int
    :rtype: List[List[int]]
    """
    if numRows == 0:
        return []
    if numRows == 1:
        return [[1]]
    ret = self.generate(numRows - 1)
    last_row = ret[len(ret) - 1]
    new_row = []
    for i in xrange(0, len(last_row) + 1):
        a = last_row[i - 1] if i - 1 >= 0 else 0
        b = last_row[i] if i < len(last_row) else 0
        new_row.append(a + b)
        ret.append(new_row)
    return ret
