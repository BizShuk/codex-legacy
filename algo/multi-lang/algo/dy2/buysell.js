function buy_sell(elems){
    var all_buysell = []
    if(elems.length <= 1) return 0;
    var min = max = -1;

    for(var i=0; i<elems.length; ++i) {
        if (i+1<elems.length && elems[i] < elems[i+1] ) {
            if (min == -1){
                min = elems[i];
            }
            continue;
        }

        if (min>=0 && (i+1 == elems.length || elems[i] > elems[i+1]) ) {
            max = elems[i]
            all_buysell.push([min,max]);
            min = max = -1;
        }
    }
    console.log(all_buysell);
    return all_buysell;
}


buy_sell([3,4,1,2,0,77,2,45]);


