# Enter your code here. Read input from STDIN. Print output to STDOUT


def removeCalc(coordinates):
    leasttime = len(coordinates) / 2
    leasttime += 1 if len(coordinates) % 2 == 1 else 0
    print leasttime


turns = int(raw_input())
i = 0
while (i < turns):
    num_points = int(raw_input())
    j = 0
    p_coordinates = []
    while (j < num_points):
        _coor = raw_input()
        p_coordinates.append(_coor.split(" "))
        j += 1
    print p_coordinates
    removeCalc(p_coordinates)

    i += 1
