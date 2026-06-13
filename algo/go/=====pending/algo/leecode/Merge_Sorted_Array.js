/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function(nums1, m, nums2, n) {

    var i = 0,
        j = 0;
    nums1.length = m;
    nums2.length = n;
    while (n != 0) {
        if (i < m + j && nums1[i] < nums2[j]) {
            i++;
        } else {
            nums1.splice(i, 0, nums2[j]);
            j++;
            i++;
            n--;
        }
    }




};
