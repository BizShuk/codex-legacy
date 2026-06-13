var tree = require('./tree.js')


var min
min = tree.findMin([1,2,3,4,5,6,7])
console.log(min);
min = tree.findMin([7,1,2,3,4,5,6])
console.log(min);
min = tree.findMin([])
console.log(min);
min = tree.findMin([-4,-5,-6,-7,0,-1,-2,-3])
console.log(min);
