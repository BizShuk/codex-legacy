/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function(n) {
    var ways = [1, 2];
    for (var i = 2; i < n; i++) {
        ways.push(ways[i - 1] + ways[i - 2]);
    }
    return ways[n - 1];
};
