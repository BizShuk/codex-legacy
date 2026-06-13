var sum = 0;

// 0: node , 1: js path , 2: start for parameters
for (var i = 2, len = process.argv.length; i < len; i++) {
    sum += Number(process.argv[i]);
}

console.log(sum);
