/**
 *  * @param {number[]} nums
 *   * @param {number} k
 *    * @return {number}
 *     */
var findPairs = function(nums, k) {
    let nums_len = nums.length;
    let count = 0;
    if (k < 0) return count;

    let e = {};
    let ret = new Set();
    for ( let i = 0 ; i < nums_len ; ++i) {
        let n=nums[i];
        if ( (n + k) in e ) ret.add(n);
        if ( (n - k) in e ) ret.add(n-k);  

        e[n] = 1;

    }
    return ret.size;

};
