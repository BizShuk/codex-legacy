package special

// [Question]: Egg dropping problem: Given 2 eggs and a 100-floor building, find the minimum egg drops required in worst the case to find the first floor at which the egg breaks.
// => 14

// [Hint]: We only have two eggs. It means first one will be used to identify the range of floor, the second one to identify which floor is exactly.
// [Hint]: It increases counter every drop test. When tests go higher of floor, the range of the floor should be 1 smaller than previous.
// [Hint]: Assume minimul number of tests is N.
// (12+1) * 12 / 2  = 78
// (13+1) * 13 / 2  = 91
// (14+1) * 14 / 2  = 15*7 = 105
// (14+2) * 13 / 2  = 13*8 = 104
// (14+3) * 12 / 2  = 17*6 = 102
// (14+4) * 11 / 2  = 11*9 =  99 => answer
