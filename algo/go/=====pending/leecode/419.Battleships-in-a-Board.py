class Solution(object):
    def countBattleships(self, board):
        """
        :type board: List[List[str]]
        :rtype: int
        """
        
        ll = len(board)
        lr = len(board[0])
        
        total = 0
        for i in range(ll):
            row = board[i]
            for j in range(lr):
                A = True if i   > 0  and board[i-1][j] == 'X' else False
                B = True if j   > 0  and board[i][j-1] == 'X' else False
                C = True if i+1 < ll and board[i+1][j] == 'X' else False
                D = True if j+1 > lr and board[i][j+1] == 'X' else False
                
                N = True if board[i][j] == 'X' else False
                #print(N,A,B,C,D)
                if N :
                    total += 1
                if N and ( A ^ B ):
                    total -= 1
                if N and (( A and B ) or ( C and D )):
                    return False
                
                    
        return total
                
