class Solution(object):
    def check_point(self, i, j, board):
        cur = board[i][j]
        count = 0
        for tmp_i in xrange(0, 9):
            if board[tmp_i][j] == cur:
                count += 1
        for tmp_j in xrange(0, 9):
            if board[i][tmp_j] == cur:
                count += 1
        i = i / 3
        j = j / 3

        for tmp_i in xrange(i * 3, i * 3 + 3):
            for tmp_j in xrange(j * 3, j * 3 + 3):
                if board[tmp_i][tmp_j] == cur:
                    count += 1
        return count == 3

    def isValidSudoku(self, board):
        """
        :type board: List[List[str]]
        :rtype: bool
        """

        for i in xrange(0, len(board)):
            for j in xrange(0, len(board[i])):
                if board[i][j] == ".":
                    continue

                if not self.check_point(i, j, board):
                    return False

        return True
