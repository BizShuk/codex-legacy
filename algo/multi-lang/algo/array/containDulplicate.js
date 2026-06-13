/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
    var h = {};

    for (var i in nums) {
        if (h[nums[i]]) return true;
        h[nums[i]] = 1;
    }
    return false
};
