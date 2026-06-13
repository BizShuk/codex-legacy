
var list = require('../linklist/listnode.js')


var isPalindrome = function(head) {
    var tmp = head;

    function check(node) {
        if (node==null) {
            return true
        }

        isPalin = check(node.next) && (node.val == tmp.val);
        tmp = tmp.next;

        return isPalin;
    }


    return check(head)
};




/*
case1 = list.Create_List(["a","b","a"])
case2 = list.Create_List(["a","a","b"])
case3 = list.Create_List(["a","d","c"])
console.log(case1);
console.log(case2);
console.log(case3);


console.log(isPalindrome(case1));
console.log(isPalindrome(case2));
console.log(isPalindrome(case3));

console.log("value");
console.log("value");
console.log("value");


case5 = list.Create_List([0,1,2,3,4,5,6,7,8,9,10])
case6 = case5.Copy();



case7 = case5.Reverse();
console.log(case7.toArray());

case8 = case5.ReverseBetween(4,7);
console.log(case8.toArray());
*/



case11 = list.Create_List([5]);
t = case11.ReverseBetween(1,1);
console.log("[5],1,1\n",t.val);


case11 = list.Create_List([1,2,3,4,5]);
t = case11.ReverseBetween(1,5);
console.log("[1,2,3,4,6],1,5\n",t);


t = case11.ReverseBetween(1,1);
console.log("[1,2,3,4,5],1,1\n",t.val);

t = case11.ReverseBetween(5,5);
console.log("[1,2,3,4,5],5,5\n",t.val,t.next.val,t.next.next.val,t.next.next.next.val,t.next.next.next.next);


t = case11.ReverseBetween(2,3);
console.log("[1,2,3,4,5],2,3\n",t.val,t.next.val,t.next.next.val,t.next.next.next.val);



case11 = list.Create_List([1,2,3]);
t = case11.ReverseBetween(2,3);
console.log("[1,2,3],2,3\n",t);

t = case11.ReverseBetween(1,2);
console.log("[1,2,3],1,2\n",t);



case11 = list.Create_List([5,3]);
t = case11.ReverseBetween(1,2);
console.log("[5,3],1,2\n",t);
