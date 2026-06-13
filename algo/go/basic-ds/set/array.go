package set

// [Pattern]: [Array Set] use array present set
// index: 0, 1, 2, 3, 4, 5, 6, 7
// value: 7,-4, 1, 1,-1, 3, 0,-3
// =>
// Set1: 1 - 2
//         \ 3 - 5
// Set2: 4
// Set3: 7 - 0 - 6

// value represent:
//      negative: set count
//      positive: predecessor

// FindSet(elem): 照出該element的root
// [Hint]: 拉平leaf這樣就可以O(1)找出root

// UnionSet(x,y): 合併sets
// [Hint]: root互相銜接, 高度小的接高度多的=>高度不變
