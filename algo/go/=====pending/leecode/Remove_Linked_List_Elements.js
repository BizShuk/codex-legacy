/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} head
 * @param {number} val
 * @return {ListNode}
 */
var removeElements = function(head, val) {
    var p = head;
    var last = head;
    while (p !== null) {

        if (p.val == val) {
            if (p === head) head = p.next;
            last.next = p.next;
        } else {
            last = p;
        }
        p = p.next;
    }
    return head;
};
