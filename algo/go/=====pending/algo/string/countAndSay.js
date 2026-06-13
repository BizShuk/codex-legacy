/**
 * @param {number} n
 * @return {string}
 */
var countAndSay = function(n) {
    if (n == 1) return "1";


    var num = countAndSay(n - 1);
    var count = 1,
        last_elem = "",
        ret = "";

    for (var i = 0, len = num.length; i < len; i++) {

        if (last_elem === num[i]) {
            count++;
        }

        if (last_elem && last_elem !== num[i]) {
            ret += count + last_elem;
            last_elem = "";
            count = 1;
        }

        last_elem = num[i];
    }

    ret += count + last_elem;

    return ret;
};

console.log(countAndSay(1));
console.log(countAndSay(2));
console.log(countAndSay(3));
console.log(countAndSay(4));
