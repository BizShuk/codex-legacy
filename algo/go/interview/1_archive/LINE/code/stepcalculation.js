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
    while(carry >=1) {
        result.push(carry%10);
        carry = Math.floor(carry/10);
    }
    resultbigint.digits = result;
    return resultbigint;
}
bigint.add = bigint.prototype.add;



let stepsCalculator = function (stepChoose) {
    if (stepChoose < 1) throw "No valid inputs";
    var stepChoose = stepChoose,
        stepWays = {
            1:new bigint(1)
        },
        curIndex = 1;

    return function (stepCount) {

        if (stepCount < 1) return 0;
        if (stepCount<=curIndex) return stepWays[stepCount];

        let i=1, sum;

        while (curIndex++<stepCount) {
            sum=new bigint(0);

            for (i = 1; i <= stepChoose; i++) {

                if (i==curIndex){
                    sum = bigint.add(sum,new bigint(1));
                } else if (stepWays[curIndex-i]!=null) {
                    sum = bigint.add(sum, stepWays[curIndex-i]);
                } else {
                    continue;
                }

            }
            stepWays[curIndex] = sum;
        }

        return stepWays[stepCount];
    }

}


let cal1 = stepsCalculator(6);

calwp = function (n) {
    console.log(cal1(n).toString());
}

calwp(610);
