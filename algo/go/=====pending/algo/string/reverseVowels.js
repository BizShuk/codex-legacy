/**
 * @param {string} s
 * @return {string}
 */
var reverseVowels = function(s) {
    var vowel = ["a", "e", "i", "o", "u", "A", "E", "I", "O", "U"];
    var ret = s.split("");
    var front = 0;
    var tail = ret.length - 1;


    while (1) {
        while (vowel.indexOf(ret[front]) < 0 && front < tail) {
            front++;
        }

        while (vowel.indexOf(ret[tail]) < 0 && front < tail) {
            tail--;
        }

        if (front < tail) {
            var tmp = ret[tail];
            ret[tail] = ret[front];
            ret[front] = tmp;
        } else {
            break;
        }
        front++;
        tail--;

    };
    return ret.join("");

};
