# ES7


### Iterators and Generators
可以暫停跟恢復的function




### Destructuring assignment
check here , https://developer.mozilla.org/zh-TW/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment

    var a, b, rest;
    [a, b] = [10, 20];
    console.log(a); // 10
    console.log(b); // 20

    [a, b, ...rest] = [10, 20, 30, 40, 50];
    console.log(a); // 10
    console.log(b); // 20
    console.log(rest); // [30, 40, 50]

    ({a, b} = {a: 10, b: 20});
    console.log(a); // 10
    console.log(b); // 20

    // 實驗性（尚未標準化）
    ({a, b, ...rest} = {a: 10, b: 20, c: 30, d: 40});

    const { a: [{foo: f}] } = obj; // f = 123
    如果用{}當assignment 需要有const, let, var或者整個expression要背()包住


### Array Comphre

    var numbers = [1, 2, 3];
    var letters = ['a', 'b', 'c'];

    var cross = [for (i of numbers) for (j of letters) i + j];
    // ["1a", "1b", "1c", "2a", "2b", "2c", "3a", "3b", "3c"]

    var grid = [for (i of numbers) [for (j of letters) i + j]];
    // [
    //  ["1a", "1b", "1c"],
    //  ["2a", "2b", "2c"],
    //  ["3a", "3b", "3c"]
    // ]

    [for (i of numbers) if (i > 1) for (j of letters) if(j > 'a') i + j]
    // ["2b", "2c", "3b", "3c"], the same as below:

    [for (i of numbers) for (j of letters) if (i > 1) if(j > 'a') i + j]
    // ["2b", "2c", "3b", "3c"]

    [for (i of numbers) if (i > 1) [for (j of letters) if(j > 'a') i + j]]
    // [["2b", "2c"], ["3b", "3c"]], not the same as below:

    [for (i of numbers) [for (j of letters) if (i > 1) if(j > 'a') i + j]]
    // [[], ["2b", "2c"], ["3b", "3c"]]


    var str = 'abcdef';
    var consonantsOnlyStr = [for (c of str) if (!(/[aeiouAEIOU]/).test(c)) c].join(''); // 'bcdf'
    var interpolatedZeros = [for (c of str) c + '0' ].join(''); // 'a0b0c0d0e0f0'


    var numbers = [1, 2, 3, 21, 22, 30];
    var doubledEvens = [for (i of numbers) if (i % 2 === 0) i * 2];
    console.log(doubledEvens); // logs 4,44,60


### block scoped variable
- let
- const
- 