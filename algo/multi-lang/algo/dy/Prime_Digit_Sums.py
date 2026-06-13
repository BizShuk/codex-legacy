#!/bin/python3

import sys


q = int(input().strip())

primes = {i:1 for i in [2,3,5,7,11,13,17,19,23,29,31,37,41,43] }

def getPrimesUnder1000():
    ret = {}
    for i in range(1,1000):
        w = "%03d"%(i)
        i1 , i2 , i3 = int(w[0]) , int(w[1]) , int(w[2])
        sum3 = i1 + i2 + i3
        if sum3 in primes:
            if i1 not in ret:
                ret[i1] = {}
            if i2 not in ret[i1]:
                ret[i1][i2] = {}
            ret[i1][i2][i3] = 1
    return ret


def getPrimesUnder100000():
    ret = {}
    for i in range(1,100000):
        w = "%05d"%(i)
        alldigits = [int(w[0]) , int(w[1]) , int(w[2]) , int(w[3]) , int(w[4])]
        sumall = sum(alldigits)
        if sumall in primes or sumall-alldigits[0] in primes or sumall-alldigits[4] in primes or alldigits[0] + alldigits[1] + alldigits[2] in primes or alldigits[1] + alldigits[2] + alldigits[3] in primes or alldigits[2] + alldigits[3] + alldigits[4] in primes :
            tmpret = ret
            for i in range(0,5):
                if alldigits[i] not in tmpret:
                    tmpret[alldigits[i]] = {}
                tmpret = tmpret[alldigits[i]]
    return ret

#primesUnder1000 = getPrimesUnder1000()
primesUnder100000 = getPrimesUnder100000()


def checkRules(s,n):
    len_s = len(s)

    if len_s >= 2:
        w = sum(s[-2:]) + n
        if w not in primes:
            return False
    if len_s >= 3:
        w = sum(s[-3:]) + n
        if w not in primes:
            return False
    if len_s >= 4:
        w = sum(s[-4:]) + n
        if w not in primes:
            return False

    return True

def getNextPossible(s):
    c = 2 if len(s) >= 2 else len(s)
    w = primesUnder1000
    for i in range(c,0,-1):
        w = w[s[len(s)-i]]

    return list(w.keys())

def getNextPossible5(s):
    c = 4 if len(s) >= 4 else len(s)
    w = primesUnder100000
    for i in range(c,0,-1):
        w = w[s[len(s)-i]]

    return list(w.keys())

for a0 in range(q):
    n = int(input().strip())
    # your code goes here

    if n <5 :
        print(0)
        continue

    stack = [[2,3,4,5,6,7,8,9]]
    nowstr = [1]

    countall = 0

    while len(stack) != 0:
        if len(nowstr) < n:
            possiblenext = getNextPossible5(nowstr)
            #print(possiblenext)
            stack.append(possiblenext)
        else:
            countall += 1
            #print(nowstr)
            nowstr.pop()

        #print(stack)
        while len(stack) != 0:
            tmps = stack.pop()

            if len(tmps) == 0:
                if len(nowstr) >0:
                    nowstr.pop()
                continue
            p = tmps.pop()
            #print(nowstr,p)
            stack.append(tmps)
            if checkRules(nowstr,p):
                nowstr.append(p)
                break
    print(countall%1000000007)


















