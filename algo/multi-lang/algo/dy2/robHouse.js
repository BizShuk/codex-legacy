/**
 * @param {number[]} nums
 * @return {number}
 */
var rob = function(nums) {
    if (nums.length == 0) return 0;

    var max = [nums[0]];

    for (var i = 1; i < nums.length; i++) {
        var last2 = (i - 2 >= 0) ? max[i - 2] : 0;
        max[i] = (max[i - 1] > nums[i] + last2) ? max[i - 1] : nums[i] + last2;
    }
    return max[nums.length - 1];

};
