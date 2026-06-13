/**
 * @param {number} n
 * @return {boolean}
 */
var isPowerOfTwo = function(n) {
    if(n<=0) return false
        for (var i=1;i<n;i*=2){
            if (n%i !== 0 ){
                return false;
            }

        }
    return true;
};

// trick n&(n-1) ==0

console.log("case:",0,"result:"isPowerOfTwo(0);
console.log("case:",1,"result:"isPowerOfTwo(1);
console.log("case:",2,"result:"isPowerOfTwo(2);
console.log("case:",3,"result:"isPowerOfTwo(3);
console.log("case:",4,"result:"isPowerOfTwo(4);
