var isPalindrome = function(s) {
    t = s.replace(/[\s\W]+/g,"").toLowerCase()
    if (t == t.split("").reverse().join("")) {
        return true;
    }

    return false;
};


console.log(isPalindrome("aab"));
console.log(isPalindrome("efe"));
console.log(isPalindrome("a."));
console.log(isPalindrome("ab@a"));
console.log(isPalindrome("c#dc"));

