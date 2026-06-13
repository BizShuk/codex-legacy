"""
    coin change with infinite coins
"""
def coin_change(amount,coins):
    Ncoins = len(coins)
    coins = sorted(coins)
    for i in range(amount):
        for j in range(Ncoins):
            coin = coins[j]
            if i - coin < 0 and i+1 == coin:
                w[i][j]=1
            if i - coin >=0 :
                w[i][j] = sum(w[i-coin][0:j+1])
    return sum(w[amount-1])