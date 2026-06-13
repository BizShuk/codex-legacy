class Solution(object):
    def reconstructQueue(self, people):
        """
        :type people: List[List[int]]
        :rtype: List[List[int]]
        """

        people.sort(key=lambda x: (-x[0],x[1]))
        ret = []
        for e in people:
            ret.insert(e[1],e)
            
        return ret
                    
                
