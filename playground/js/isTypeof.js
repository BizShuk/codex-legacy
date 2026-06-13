var testdata = require('./testdata.js')()
var elems = testdata.data;
var ins = testdata.instance;



elems.forEach((e) => {
    console.log(e,typeof e);
});
ins.forEach((e) => {
    console.log(e,typeof e);
});