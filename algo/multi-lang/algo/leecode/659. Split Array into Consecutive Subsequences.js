/**
# 659. Split Array into Consecutive Subsequences
You are given an integer array sorted in ascending order (may contain duplicates), you need to split them into several subsequences, where each subsequences consist of at least 3 consecutive integers. Return whether you can make such a split.

Example 1:

    Input: [1,2,3,3,4,5]
    Output: True
    Explanation:
    You can split them into two consecutive subsequences :
    1, 2, 3
    3, 4, 5

Example 2:

    Input: [1,2,3,3,4,4,5,5]
    Output: True
    Explanation:
    You can split them into two consecutive subsequences :
    1, 2, 3, 4, 5
    3, 4, 5

Example 3:

    Input: [1,2,3,4,4,5]
    Output: False
    Note:
    The length of the input is in range of [1, 10000]

 * @param {number[]} nums
 * @return {boolean}
 */
var isPossible = function(nums) {

    if (nums.length <= 0) return false;

    var last_number=nums[0],i=1;
    var sequence_set=[[nums[0]]];


    var cur_num,same_value_set;
    while(i<nums.length){
        cur_num = nums[i];
        same_value_set = [];
        last_number = null;

        while(i<nums.length){
            if (last_number == null || last_number == nums[i]) {
                same_value_set.push(nums[i]);
                last_number = nums[i];
            } else {
                break;
            }

            i++;
        }

        var j = sequence_set.length - 1;

        same_value_set.forEach((e) => {
            //console.log(j,sequence_set,same_value_set)
            var last_set = sequence_set[j];
            if (j >= 0 && last_set[last_set.length-1] == e - 1) {
                last_set.push(e);
                j--;
            } else {
                sequence_set.push([e]);
            }
        });

    }
    //console.log(sequence_set);
    // check sequence_set for each elements length, all should big or equal thant 3
    for( var i in sequence_set) {
        var e = sequence_set[i];
        if( !(e instanceof Array) || e.length < 3) return false;
    }
    return true;
};













// public boolean isPossible(int[] nums) {
//     Map<Integer, Integer> freq = new HashMap<>(), appendfreq = new HashMap<>();
//     for (int i : nums) freq.put(i, freq.getOrDefault(i,0)+1);
//     for (int i : nums) {
//         if (freq.get(i) == 0) continue;
//         else if (appendfreq.getOrDefault(i,0) > 0) {
//             appendfreq.put(i, appendfreq.get(i)-1);
//             appendfreq.put(i+1, appendfreq.getOrDefault(i+1,0)+1);
//         }
//         else if (freq.getOrDefault(i+1,0) > 0 && freq.getOrDefault(i+2,0) > 0) {
//             freq.put(i+1, freq.get(i+1) - 1);
//             freq.put(i+2, freq.get(i+2) - 1);
//             appendfreq.put(i+3, appendfreq.getOrDefault(i+3,0) + 1);
//         }
//         else return false;
//         freq.put(i, freq.get(i) - 1);
//     }
//     return true;
// }