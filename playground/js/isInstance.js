// 確認各種instance


var testdata = require('./testdata.js')();

var elems = testdata.data;

var instances = testdata.instance;

elems.forEach((e)=>{
    instances.forEach((ins) => {
        if ( elems.indexOf(e) < elems.length-3  // 不是自訂object,就不用跑自己class比較
            && instances.indexOf(ins) > instances.length-3) return;
        console.log(e,ins,e instanceof ins);
    });
    console.log();
})

