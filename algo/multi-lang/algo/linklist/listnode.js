function ListNode(val) {
    this.val = val;
    this.next = null;
    this.prev = null;
}

// reverse single list
ListNode.prototype.Reverse = function() {
    var prev,next,cur;
        cur = this.Copy();

    while (cur!==null) {
        next = cur.next;
        cur.next = prev;
        prev = cur;
        cur = next;
    }

    return prev;
}

Reverse_r = function(head) {
    
}


/**

 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} head
 * @param {number} m
 * @param {number} n
 * @return {ListNode}
 */

ListNode.prototype.ReverseBetween = function(m, n) {
    var newlist = this.Copy();

    if (newlist == null) { return null; }

    var dummy = new ListNode(0);
    dummy.next = newlist;
    pre = dummy;
    for (var i = 0, len = m-1; i < len; i++) {
        pre = pre.next;
    }

    start = pre.next;
    then = start.next;

    for (var i = 0, len = n-m; i < len; i++) {
        start.next = then.next;
        then.next = pre.next;
        pre.next = then;
        then = start.next;
    }


    return newlist;
};


ListNode.prototype.toArray = function() {
    var cur = this;
    var val = [];
    while (cur != null) {
        val.push(cur.val);
        cur = cur.next;
    }
    return val;
}

ListNode.prototype.Copy = function() {
    var cur,newhead
        cur = this;
        newhead = new ListNode(cur.val);
        newhead_cur = newhead;

        while (1) {
            if (cur.next===null) {
                newhead_cur.next = null;
                break;
            }
            newhead_cur.next = new ListNode(cur.next.val);
            newhead_cur = newhead_cur.next
            cur = cur.next;

        }
        return newhead;
}



// @param array arr
// @return root
function Create_List(arr) {
    if (arr.length==0) {
        return null
    }

    var cur,next,root
        next = null
    for (var i = arr.length-1; i >= 0 ; i--){ 
        cur = new ListNode(arr[i]);
        cur.next = next;
        next = cur;
    }

    return cur
}

module.exports = {
    ListNode : ListNode,
    Create_List : Create_List,
}
