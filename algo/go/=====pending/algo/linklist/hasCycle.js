/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} head
 * @return {boolean}
 */
var hasCycle = function(head) {
    if (head == null) return false;
    var s1 = head;
    var s2 = head;

    while (s1 !== null && s2 !== null && s2.next !== null) {
        s1 = s1.next;
        s2 = s2.next.next;
        if (s1 == s2) return true;
    }

    return false;

};
