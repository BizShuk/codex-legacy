/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function(s) {
    // I:1 V:5 X:10 L:50 C:100 D:500 M:1000
    var map = {
        I: 1,
        V: 5,
        X: 10,
        L: 50,
        C: 100,
        D: 500,
        M: 1000
    };

    var tmp_value = 0;
    var last_char = "";
    var ret = 0;
    var i = 0;

    while (1) {
        if (i >= s.length) {
            ret += tmp_value;
            break;
        }

        if (last_char == s[i]) {
            tmp_value += map[s[i]];
        }

        if (last_char != s[i]) {
            if (last_char !== "" && map[last_char] < map[s[i]]) {
                tmp_value = tmp_value * (-1);
            }
            ret += tmp_value;
            last_char = s[i];
            tmp_value = map[s[i]];
        }

        i++;
    }


    return ret;

};
