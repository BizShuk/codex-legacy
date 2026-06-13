/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function(nums) {
    var count = 0,
        cur_num = "";
    for (var i in nums) {
        if (count == 0) {
            cur_num = nums[i];
        }
        if (cur_num == nums[i]) {
            count++;
        } else {
            count--;
        }
    }
    return cur_num;
};


console.log(majorityElement([1, 2, 3, 2, 2, 1, 2, 3, 2]));
console.log(majorityElement([1, 1, 1, 1, 1, 3, 3, 2, 4]));
console.log(majorityElement([1, 2, 2]));
console.log(majorityElement([3]));
