# buy mutiple time

# buy once only or twice

# backward way for max

# buy or self at each time  ,
def maxProfit(stocks):
    profit,end,s=0,len(stocks),stocks
    maxlast = 0

    def nextProfit():
        tradeprofit = 0
        buyTime = []
        tmpmax = maxlast
        ti = end
        while ti > 0:
            ti-=1
            if tmpmax < s[ti]:
                tmpmax = s[ti]
            if tmpmax - s[ti]>0:
                tradeprofit = tmpmax-s[ti]
                break
        return tradeprofit,ti,tmpmax

    while True:
        n,end,maxlast = nextProfit()
        profit += n
        if end <= 0:
            break

    return profit


cases = int(input())

for case in range(cases):
    s,data = int(input()),list(map(int,input().strip().split()))
    w= maxProfit(data)
    print(w)