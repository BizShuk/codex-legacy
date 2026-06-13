/**
 * @param {number} n - a positive integer
 * @return {number} - a positive integer
 */
var reverseBits = function(n) {
    var i=0, ret =0;
    while(i<32){
        ret = (ret<<1) | (n&1);
        n >>=1;
        i++;
    }
    return ret>>>0;
};

// Not working , because all numbers in JavaScript are 64-bit (double-precision) floating point numbers. 
var reverseBits_or = function(n) {
    console.log(n,n.toString(2));
    var tmp = n > 0;

    n = n >>> 16 | (n << 16);
    n = ((n & 0xff00ff00) >>> 8) | ((n & 0x00ff00ff) << 8);
    n = ((n & 0xf0f0f0f0) >>> 4) | ((n & 0x0f0f0f0f) << 4);
    n = ((n & 0xcccccccc) >>> 2) | ((n & 0x33333333) << 2);
    n = ((n & 0xaaaaaaaa) >>> 1) | ((n & 0x55555555) << 1);
    if (tmp && n<0) {
        n = -n;
    }
    return n;               
}

cases=1
a=reverseBits(cases);
console.log("cases:",cases,"answer:",a,"assert:",2147483648);
cases=2
a=reverseBits(cases);
console.log("cases:",cases,"assert:",a);
cases=2147483548
a=reverseBits(cases);
console.log("cases:",cases,"assert:",a);
cases=4294967295
a=reverseBits(cases);
console.log("cases:",cases,"answer:",a,"assert:",  4294967295);





