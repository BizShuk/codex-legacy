/**
 *  * @param {number[]} nums1
 *   * @param {number[]} nums2
 *    * @return {number[]}
 *     */
var intersect = function(nums1, nums2) {
    var h = {};
    var ret = [];
    if (nums1.length < nums2.length) {
        a = nums1;
        b = nums2;
    } else {
        a = nums2;
        b = nums1;
    }

    for (var i in a) {
        (a[i] in h) && (h[a[i]]++) || (h[a[i]] = 1);
    }

    for (var i in b) {
        if (h[b[i]] > 0) {
            h[b[i]]--;
            ret.push(b[i]);
        }

    }


    return ret;
};



console.log(intersect([1], [1]));
