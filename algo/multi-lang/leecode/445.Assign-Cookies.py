class Solution(object):
    def findContentChildren(self, g, s):
        """
        :type g: List[int]
        :type s: List[int]
        :rtype: int
        """
        g.sort(lambda a,b: a-b )
        s.sort(lambda a,b: a-b )

        i,j,ret=0,0,0
        while True:
            if i >= len(g) or j >= len(s):
                break

            gv,sv = g[i],s[j]

            if gv <= sv :
                ret+=1
                i+=1
                j+=1

            if gv > sv :
                j+=1

        return ret

            

            

            

            

            

            

            

            

            

            

            

            

        

        

        

        

        

        

        
