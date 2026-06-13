import random


class abtest():
    """TODO: Docstring for abtest.
    :returns: TODO

    """

    def __init__(self):
        self.a = 0
        self.b = 0
        self.c = 0

    def rand(self):
        perc = random.random()
        if perc < 0.2:
            self.a += 1
        elif perc < 0.6:
            self.b += 1
        else:
            self.c += 1

    def show(self):
        total = self.a + self.b + self.c
        print self.a / (total * 1.0)
        print self.b / (total * 1.0)
        print self.c / (total * 1.0)


i = 0
ins = abtest()
while i < 1000000:
    ins.rand()
    i += 1

ins.show()
