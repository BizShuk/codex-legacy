/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var swapPairs = function(head) {
    var fh = new ListNode(0);
    var tmpp;
    fh.next = head;
    var p = fh;
    while (p.next && p.next.next) {
        tmpp = p.next.next.next;
        p.next.next.next = p.next;
        p.next = p.next.next;
        p.next.next.next = tmpp;
        p = p.next.next;
    }
    return fh.next;
};
