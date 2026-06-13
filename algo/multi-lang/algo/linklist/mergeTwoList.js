/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var mergeTwoLists = function(l1, l2) {
    var ret = new ListNode(0);
    var tmp = ret;
    while (l1) {
        if (l2 && l2.val < l1.val) {
            tmp.next = l2;
            l2 = l2.next;
        } else {
            tmp.next = l1;
            l1 = l1.next;
        }

        tmp = tmp.next;
    }

    if (l2) {
        tmp.next = l2;
    }


    return ret.next;

};
