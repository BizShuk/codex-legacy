/**
 *  * @param {number} x
 *  * @return {boolean}
 */
var isPalindrome = function(x) {

    if (x<0||(x!==0&&x%10===0)) return false;

    var last=0;

    while (parseInt(x)>last){
        last = last*10 +parseInt(x%10);
        x = x/10;
    }

    return (x===0 || parseInt(x)===last || parseInt(x*10)===last);
};  





console.log(2034114302,isPalindrome(2034114302));
console.log(isPalindrome(123));
console.log(isPalindrome(121));
console.log(isPalindrome(11));
console.log(isPalindrome(-2147483648));
console.log(isPalindrome(123));
console.log(isPalindrome(121));
console.log(isPalindrome(9123123));
console.log(isPalindrome(0));
console.log(isPalindrome(10));

