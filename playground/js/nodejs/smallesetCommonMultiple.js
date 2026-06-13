function smallestCommons(arr) {

    arr.sort(function(a,b){ return a-b;});
    var small = arr[0];
    var big = arr[1];

    var b_primes = {};

    var smc_factors = {};

    while (big >= small) {
        var dividend = small;
        var divisor_start_from = 2;
        var dividend_factors = {};

        console.log("small;",small);

        for (var i=divisor_start_from ; i<=dividend ; i++){
            if (typeof b_primes[dividend] !== "undefined"){
                var all_factors = get_all_factors(dividend);

                console.log(i,dividend,dividend_factors,all_factors);

                Object.keys(all_factors).forEach(function(e,i){
                        dividend_factors[e] = dividend_factors[e] + all_factors[e] || all_factors[e];
                        return e;
                        });
                break;
            }

            if (dividend % i === 0){
                var old_dividend = dividend;
                dividend_factors[i]=1+dividend_factors[i] || 1;
                dividend = dividend / i;
                b_primes[old_dividend] = [i,dividend];
                i--;
            }

        }
        // console.log(b_primes,dividend_factors);
        Object.keys(dividend_factors).forEach(function(e,i){
            if (typeof smc_factors[e] == 'undefined' || dividend_factors[e]>=smc_factors[e]) smc_factors[e]=dividend_factors[e];
            return e;
        })
        small++;
    }

    function get_all_factors(num){
        if ( num === 1 ) return {};
        var all_factors = get_all_factors(b_primes[num][1]);
        all_factors[b_primes[num][0]] = 1 + all_factors[b_primes[num][0]] || 1;
        return all_factors;
    }


    var smc = 1;
    Object.keys(smc_factors).forEach(function(e,i){
        smc = smc * Math.pow(parseInt(e),smc_factors[e]);
        return e;
    });
    console.log(smc_factors,smc,smc_factors);
    return smc;
}

smallestCommons([13,1]);

