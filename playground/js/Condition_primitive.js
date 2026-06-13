// 判斷es5 primitive variable condition

var elems3 = require('./testdata.js')().data;
var elems4 = require('./testdata.js')().data;
console.log(elems3);

elems3.forEach((c1) => {
    elems4.forEach((c2) => {
        console.log(c1,c2,c1 == c2);
    })
})