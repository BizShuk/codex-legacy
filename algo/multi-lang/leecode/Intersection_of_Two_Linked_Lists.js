/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} headA
 * @param {ListNode} headB
 * @return {ListNode}
 */
var getIntersectionNode = function(headA, headB) {

    var p1 = headA,
        p2 = headB;
    var a = true,
        b = true;
    if (p1 == null || p2 == null) return null;

    while (p1 != null && p2 != null) {

        if (p1 == p2) return p1;
        p1 = p1.next;
        p2 = p2.next;

        if (p1 === null && a) {
            p1 = headB;
            a = false;
        }
        if (p2 === null && b) {
            p2 = headA;
            b = false;
        }

    }
    return null;
};
