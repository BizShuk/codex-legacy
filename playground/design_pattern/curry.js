function add() {
    var elements = [],
        howmany = 3;


    function fn() {
        var el = elements.length;

        for (var i = 0, len = arguments.length; i < len && i + el < howmany; i++) {
            elements.push(arguments[i])
        }

        if (elements.length >= howmany) {
            return elements.reduce((a, b) => a + b);
        }

        return fn

    }
    return fn(...arguments);
}

r = add(1, 2, 3);
console.log(r);
r = add(1, 2)(3);
console.log(r);
r = add(1)(2, 3);
console.log(r);
r = add(1)(2)(3);
console.log(r);
