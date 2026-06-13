class Solution(object):
    def getHint(self, secret, guess):
        """
        :type secret: str
        :type guess: str
        :rtype: str
        """

        secret_hash = {}
        A = 0
        B = 0
        for i in xrange(0, len(secret)):
            c = secret[i]
            secret_hash[c] = secret_hash[c] + 1 if c in secret_hash else 1
            if secret[i] == guess[i]:
                A += 1

        for c in guess:
            if c in secret_hash and secret_hash[c] > 0:
                B += 1
                secret_hash[c] -= 1
        B = B - A
        return "{}A{}B".format(A, B)
