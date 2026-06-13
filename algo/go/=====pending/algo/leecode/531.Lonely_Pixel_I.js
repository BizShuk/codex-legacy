/**
 *  * @param {character[][]} picture
 *   * @return {number}
 *    */
var findLonelyPixel = function(picture) {
    let r_len = picture.length,
    c_len = picture[0].length,
    r_e=[];
    c_e=[];

    for ( let i = 0 ; i < r_len ; ++i ){
        r_e.push(0);
    }
    for ( let i = 0 ; i < c_len ; ++i ){
        c_e.push(0);
    }

    picture.forEach(function(rv,rk){
        rv.forEach(function(cv,ck){
            if (cv === 'B') {
                r_e[rk]++;
                c_e[ck]++;
            }
        });
    });


    let count = 0;   
    for ( let i = 0 ; i < r_len ; ++i ){
        if ( r_e[i] !== 1 ) continue;

        let row = picture[i];
        let Bi = row.indexOf('B');

        if ( c_e[Bi] === 1 ){ 
            count++;
            continue;
        }

    }
    return count;
};

var ret_rcount = function(arr,match){
    let count = 0;
    arr.forEach(function(v,i){
        if (v === match) count++;
    });
    return count;
};
