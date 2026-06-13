

// for IE8
if (!Array.prototype.indexOf) {
    Array.prototype.indexOf = function(what, i) {
        i = i || 0;
        var L = this.length;
        while (i < L) {
            if (this[i] === what) return i;
            ++i;
        }
        return -1;
    }
}



// ways to check integer
// ref: http://stackoverflow.com/questions/14636536/how-to-check-if-a-variable-is-an-integer-in-javascript
//function (data) {
//    if (data === parseInt(data, 10)) return true;
//    if(Number.isInteger(Number(data))) return true;
//    return false;
//}