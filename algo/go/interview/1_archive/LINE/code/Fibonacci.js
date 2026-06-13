var bigint = function(n){
    this.digits = [];
    this.digits.push(n%10);
    n = Math.floor(n/10);
    while (n>=1) {
        this.digits.push(n%10);
        n = Math.floor(n/10);
    }
}

bigint.prototype.toString = function() {
    return this.digits.reverse().join("");
}

// array
bigint.prototype.add = function(num1, num2) {
    if (!(num1 instanceof bigint && num2 instanceof bigint)) throw "Not valid Inputs";
    var resultbigint = new bigint(0);


    let result=[], arrA=num1.digits, arrB=num2.digits;
    let i = 0, carry=0, sum=0, a, b, al=arrA.length, bl=arrB.length;

    while(i<al || i<bl) {
        a = arrA[i] || 0;
        b = arrB[i] || 0;
        sum = a+b+carry;
        result.push(sum%10);
        carry = Math.floor(sum/10);
        i++;
    }
    while( carry >=1) {
        result.push(carry%10);
        carry = Math.floor(carry/10);
    }
    resultbigint.digits = result;
    return resultbigint;
}
bigint.add = bigint.prototype.add;



let fibonacci_contructor = () => {
    let i = 1,
        fv = { 0:new bigint(0), 1:new bigint(1) };
    // recurrsive
    // return function (n) {
    //     // console.log(typeof n,n);
    //     if (typeof n !== "number" || n < 0) return false;
    //     if ( n in fv ) return fv[n];
    //     console.log(n);
    //     fv[n] = fibonacci(n-1) + fibonacci(n-2);
    //     return fv[n];
    // }

    return function (n) {
        if (typeof n != "number" || n < 0) return false;
        i = n;

        while( fv[i]==null ){
            i--;
        }

        while( i++<n ) {
            fv[i] = bigint.add(fv[i-1],fv[i-2]);
        }

        return fv[n];
    }
};

let fibonacci = fibonacci_contructor();
let answer = fibonacci(8181);
console.log(answer.toString());
answer = fibonacci(5);
console.log(answer.toString());
