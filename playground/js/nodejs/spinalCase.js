function spinalCase(str) {
    // "It's such a fine line between stupid, and clever."
    // --David St. Hubbins

    var uppercase_break = /(\w)([A-Z])/g;
    var breakword = /[_ ]+/g;

    str = str.replace(uppercase_break,"$1 $2").toLowerCase();
    str = str.replace(breakword,"-");


    return str;
}

spinalCase('This Is Spinal Tap');

