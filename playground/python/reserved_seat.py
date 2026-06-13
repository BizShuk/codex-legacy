def reserved_seat(N, S, p):
    n = 0
    if p == 1:
        pass
    elif p == 2:
        pass
    elif p == 3:
        n = reserved_seat_3(N, S)
    else:
        pass

    return n


def reserved_seat_3(N, S):
    max = N * 3

    # init seats
    all_seats = []
    for i in xrange(0, N):
        all_seats.append({"A": 0,
                          "B": 0,
                          "C": 0,
                          "D": 0,
                          "E": 0,
                          "F": 0,
                          "G": 0,
                          "H": 0,
                          "I": 0,
                          "J": 0,
                          "K": 0})

    takens = S.strip().split(" ") if len(S.strip()) > 0 else []

    for seat in takens:
        _row = int(seat[0])
        _col = seat[1]
        all_seats[_row - 1][_col] = 1

    return max


if __name__ == "__main__":
    reserved_seat(3, ' 1A 2B ', 3)
    reserved_seat(1, '   1A  ', 3)
    reserved_seat(3, '  ', 3)
    reserved_seat(0, '  ', 3)
    pass
