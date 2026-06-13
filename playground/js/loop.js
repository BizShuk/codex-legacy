Array.prototype.x = "testX"

var t = [1,2,3]
t.length =4
t[-1] = 5
t.p = "p"


function inheritedArray(){}
inheritedArray.prototype = Array.prototype


var newArr = new inheritedArray()
newArr.push(1,2,3);



console.log("for in(show all enumerable properties):");
for(var k in t) {
    console.log(k);

    if(t.hasOweProperty(k)) {
        console.log(k); // only show on its owe properties, not including which are inherited.
    }
}

console.log("for of(show only iterable part):");
for(var k of t) {
    console.log(k);
}

console.log("forEach(show defined index of array):");
t.forEach(function(i,e){
    console.log(i,e);
})





console.log("for in(show all enumerable properties):");
for(var k in newArr) {
    console.log(k);
}

console.log("for of(show only iterable part):");
for(var k of newArr) {
    console.log(k);
}

console.log("forEach(show defined index of array):");
newArr.forEach(function(i,e){
    console.log(i,e);
})



