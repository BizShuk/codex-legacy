The integer V is strictly between integers U and W if U < W < V < W or if U > V > W.
Gives a non-null zero-indexed array A consisting of N integers A. A pair of exponents (P, Q), where 0 <P <Q <If none of the values ​​in the array are strictly between the values ​​A [P] and A [Q ], Then N has the adjacent value. 

For example, in array A like this: 

A [0] = 0
A [1] = 3
A [2] = 3
A [3] = 7
A [4] = 5
A [5] = 3
A [6] = 11
A [7] = 1 

The following index pairs have adjacent values: 
(0,7), (1,2), (1,4),
(1,5), (1,7), (2,4),
(2,5), (2,7), (3,4),
(3,6), (4,5), (5,7). 

For example, index 4 and index 5 have contiguous values ​​because there is no strict value between A [4] = 5 and A [5] = 3 in array A; the only such value may be number 4, which Does not exist in the array.

Given two indices P and Q, their distance is defined as abs (A [P] -A [Q]) where abs (X) = X for abs X) = -X for X < For example, the distance between indexes 4 and 5 is 2 because (A [4] - A [5]) = (5 - 3) = 2.

Write a function:
def solution (A)

Given a non-zero-space index array A consisting of N integers, return the minimum distance between indices with adjacent values ​​in this array. If the minimum distance is greater than 100,000,000, the function should return -1. If there is no adjacent index, the function should return -2.