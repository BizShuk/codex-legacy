

function myReplace (str, before, after){
    var after_formated = "";
    for (var i=0 ; i<before.length && i<after.length ; i++){
        if (before[i].match(/[A-Z]/g)){
            after_formated += after[i].toUpperCase();
        }else{
            after_formated += after[i].toLowerCase();
        }    
    }   

    str = str.replace(before, after_formated);
    console.log(str);
    return str;
}



myReplace("A quick brown fox jumped over the lazy dog", "jumped", "leaped");
myReplace("He is Sleeping on the couch", "Sleeping", "sitting");
