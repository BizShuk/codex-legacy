class Solution(object):
    def fizzBuzz(self, n):
        """
        :type n: int
        :rtype: List[str]
        """

        ret = []

        i=1
        while i <= n :
            astr = ""

            if i%3 == 0 :
                astr += "Fizz"
            if i%5 == 0 :
                astr += "Buzz"
            if i%3 != 0 and i%5!=0:
                astr = str(i)

            ret.append(astr)
            i+=1
        return ret
