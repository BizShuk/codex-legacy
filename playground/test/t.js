

function e(i) {

    return function(){
        console.log(i);

    }
}


for (var i = 10 - 1; i >= 0; i--) {
    setTimeout(e(i),0);
}