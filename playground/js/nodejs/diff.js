var fs = require('fs');
var path = require('path');
var readline = require('readline');



var file_a = ['a\n','b\n','c\n','d\n','e\n','f\n','g\n','h\n','i\n','j\n','k\n','o\n'];
var file_b = ['b\n','c\n','d\n','e\n','k\n','h\n','i\n','j\n','k\n','w\n'];


var diff = (function(root){
    var diff = {};
    diff.lcs_count = function(a,b){
        var result = 0;
    
        if( a.length == 0 || b.length == 0 ){
        }else if( a[0] == b[0] ){    
            result = 1 + this.lcs_count( a.slice(1) , b.slice(1) );
        }else{
            result = Math.max(  this.lcs_count( a , b.slice(1)) , 
                                this.lcs_count( a.slice(1) , b ) )
        }

        return result;
    }

    diff.lcs_line = function(a,b){
        var result = 0;
        var lines = "";

        if( a.length == 0 ){

            lines = b.length==0 ? lines : lines + color_add( b.join() );
        }else if( b.length == 0 ){

            lines = a.length==0 ? lines : lines + color_del( a.join() );
        }else if( a[0] == b[0] ){    

            lines += a[0] + this.lcs_line( a.slice(1) , b.slice(1) , lines);
        }else{

            var a_longest = this.lcs_count( a , b.slice(1) , lines );
            var b_longest = this.lcs_count( a.slice(1) , b , lines );

            if( a_longest > b_longest ){    
                lines = lines + color_add(b[0]) + this.lcs_line( a , b.slice(1) , lines );
            }else{                          
                lines = lines + color_del(a[0]) +this.lcs_line( a.slice(1) , b , lines );
            }
    
        }

        return lines;
    };

    var color_add = function( v ){
        return "\x1b[32m+" + v + "\x1b[m";  
    };

    var color_del = function( v ){
        return "\x1b[31m-" + v + "\x1b[m";  
    };

    return diff;
})()
    
var result = diff.lcs_count( file_a , file_b );
console.log( result );

var result_line = diff.lcs_line( file_a , file_b );
console.log( result_line )


