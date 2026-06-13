class Person:
    def __init__(self, firstName, lastName, idNumber):
        self.firstName = firstName
        self.lastName = lastName
        self.idNumber = idNumber
    def printPerson(self):
        print("Name:", self.lastName + ",", self.firstName)
        print("ID:", self.idNumber)

    @classmethod
    @staticmethod
    @abstractmethod
    def display():

class Student(Person):
    def __init__(self, firstName, lastName, idNumber , scores):
        super(self.__class__ , self).__init__(firstName , lastName , idNumber)
        self.scores = scores
    def calculate(self):
        grade = sum(self.scores) / len(self.scores)
        if 90 <= grade and grade <= 100:
            grade = "O"
        elif 80 <= grade and grade < 90:
            grade = "E"
        elif 70 <= grade and grade < 80:
            grade = "A"
        elif 55 <= grade and grade < 70:
            grade = "P"
        elif 40 <= grade and grade < 55:
            grade = "D"
        else:
            grade = "T"
        return grade   

# input from file
f = open('oo.input' , 'r')
input = f.readline

# test code
line = input().split()
firstName = line[0]
lastName = line[1]
idNum = line[2]
numScores = int(input()) # not needed for Python
scores = list( map(int, input().split()) )
s = Student(firstName, lastName, idNum, scores)
s.printPerson()
print("Grade:", s.calculate())
