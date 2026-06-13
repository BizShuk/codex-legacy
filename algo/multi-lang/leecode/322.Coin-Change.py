class Solution(object):
    def coinChange(self, coins, amount):
        """
        :type coins: List[int]
        :type amount: int
        :rtype: int
        """
        cl = len(coins)
        dp = [-1 for i in range(amount+1)]
        dp[0] = 0
        
        for i in range(min(coins),amount+1):
            m=[dp[i-c] for c in coins if i-c >=0 and dp[i-c] != -1]
                
            dp[i] = min(m)+1 if len(m)>0 else -1
        return dp[amount]
    
    def coinChange_2(self, coins, amount):
        """
        :type coins: List[int]
        :type amount: int
        :rtype: int
        """
        cl = len(coins)
        dp = [[-1 for i in range(cl+1)] for j in range(amount+1)]
        for i in range(cl+1):
            dp[0][i] = 0
            
        for i in range(cl-1,-1,-1):
            c = coins[i]
            for j in range(c,amount+1):
                m = [v for v in dp[j-c] if v!= -1]
                
                dp[j][i] = min(m) +1 if bool(m) else -1
        
        m = [v for v in dp[amount] if v != -1]
        return min(m) if bool(m) else -1
