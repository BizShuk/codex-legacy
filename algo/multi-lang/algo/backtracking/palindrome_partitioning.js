/**
 *
 *  * @param {string} s
 *
 *   * @return {string[][]}
 *
 *    */

var partition = function(s) {   
    var tokens = [] ;
    ret = palindrome(tokens,s,0);
    console.log("result:",ret);
    return ret;
};


function palindrome(tokens,s,i){
    if (i>=s.length){
        return [tokens];
    }

    var p1=[],_s;
    for (var j = 1, len = s.length-i; j <= len; j++) {
        _s = s.substring(i,j+i);
        if (_s == _s.split("").reverse().join("")){
            p1 = p1.concat(palindrome(tokens.concat([_s]),s,i+j));
        }
    }

    return p1;
}

partition("aab")
partition("efe")










