class Solution(object):
    def findAnagrams(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: List[int]
        """

        pl = len(p)
        sl = len(s)

        sd = {}
        pd = {}
        for i in p:
            pd[i] = pd[i]+1 if i in pd else 1

        ret = []

        i = 0
        while i < sl:
            c = s[i]

            if i < pl:
                sd[c] = sd[c]+1 if c in sd else 1
            else:
                sd[c] = sd[c]+1 if c in sd else 1
                sd[s[i-pl]] -= 1

            all_matched = True


            for k in pd:
                if k not in sd or sd[k]!=pd[k]:
                    all_matched = False
                    break

            for k in sd:
                if (k not in pd or sd[k]!=pd[k]) and sd[k]!=0:
                    all_matched = False
                    break

            if all_matched:
                ret.append(i-pl+1)

            i+=1

        return ret

            
    def findAnagrams_2(self,s,p):
        d = {}

        for i in range(len(s)-len(p)+1):
            ss = "".join(list(s[i:i+len(p)]))
            ss.sort()
            if ss in d:
                d[ss] += [i]
            else:
                d[ss] = [i] 


        pstr = "".join(list(p).sort())

        return d[pstr] if pstr in d else []
        
