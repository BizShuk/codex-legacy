const MY_FAV = 7;

// this will throw an error - Uncaught TypeError: Assignment to constant variable.
// MY_FAV = 20;

// MY_FAV is 7
console.log('my favorite number is: ' + MY_FAV);

// trying to redeclare a constant throws an error
// Uncaught SyntaxError: Identifier 'MY_FAV' has already been declared
// const MY_FAV = 20;

// the name MY_FAV is reserved for constant above, so this will fail too
// var MY_FAV = 20;

// this throws an error too
// let MY_FAV = 20;

if (MY_FAV === 7) {
    // this is fine and creates a block scoped MY_FAV variable
    // (works equally well with let to declare a block scoped non const variable)
    let MY_FAV = 20;

    // MY_FAV is now 20
    console.log('my favorite number is ' + MY_FAV);

    // this gets hoisted into the global context and throws an error
    // var MY_FAV = 20;
}

// MY_FAV is still 7
console.log('my favorite number is ' + MY_FAV);
