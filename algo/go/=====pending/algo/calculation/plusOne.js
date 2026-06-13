/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function(digits) {


    digits[digits.length - 1]++;
    for (var i = digits.length - 1; i > 0; i--) {
        digits[i - 1] += Math.floor(digits[i] / 10);
        digits[i] = digits[i] % 10;
    }

    if (digits[0] >= 10) {
        digits.unshift(Math.floor(digits[i] / 10));
        digits[i + 1] = digits[i + 1] % 10;
    }
    return digits;

};
