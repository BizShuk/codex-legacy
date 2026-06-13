function translate(str) {
      
    if (str[0].match(/[aeiou]/)) {
        str += "way";
    }else{
        var newstr ="";
        var tmpstr ="";
        
        for (var i=0; i<str.length ; i++){
            if ( str[i].match(/[aeiou]/) ){
                newstr = str.substr(i);
                break;
            }else{
                tmpstr += str[i];
            }
        }
        str = newstr + tmpstr + "ay";
    }
    
    console.log(str);      
    return str;
}

translate("consonant");
translate("california");
translate("paragraphs");
translate("glove");
translate("algorithm");
translate("eight");
