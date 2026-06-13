// 比較特殊狀況的condition


// null => false
// undefined => false
// 0 or number => true
// "" or "string" => true
// [] or [1,2,3] => true
// {} or {"1":2} => true


var compareMethod = [];

// compareMethod.push((e) => console.log(e,'with condition !e',!e) );
// compareMethod.push((e) => {
//     console.log(e,'with condition typeof e != "undefined"', typeof e != "undefined");
// });
// compareMethod.push((e) => {
//     console.log(e,'with condition typeof e != null',typeof e != null);
// });


// Useful
compareMethod.push((e) => {
    console.log(e,'with condition !e', !e);
});

compareMethod.push((e) => { // only work for Null and Undefined
    console.log(e,'with condition null == e',null == e);
});

compareMethod.push((e) => { // work for Null, Undefined, and NaN
    console.log(e,'with condition !e && 0!=e',!e &&  0!=e);
});


var showConditions = () => {

    // only undefined and null should get true
    var elems = require('./testdata.js')().data;
    elems.forEach((e) => {
        compareVar(e);
    });
}

var compareVar = (elem) => {
    compareMethod.forEach((md) => {
        md(elem);
    });
    console.log("");
}

showConditions();