// TODO:
//  string comparsion

var Comparsion=(function (){

    function num_compare( arr1 , arr2 ){

        if( arr1.length !=0 && arr2.length !=0 ){
            var a1 = parseInt(arr1.shift());
            var a2 = parseInt(arr2.shift());

            if( ! isNaN(a1) && ! isNaN(a2) ){
                
                if( a1 > a2 ){
                    return 1;  
                }else if( a1 < a2 ){
                    return -1;
                }else{
                    return num_compare( arr1 , arr2 );
                }
                
            }else if( (isNaN( a1 ) && isNaN( a2 )) || a1 == a2 ){
                return num_compare( arr1 , arr2 );
            }else if( isNaN( a2) ){
                return 1;   
            }else{
                return -1;
            }
        }else if( arr1.length == 0 ){
            return -1;
        }else if( arr2.length ==0 ){
            return 1;
        }else{
            return 0;
        }

    }

    function str_compare( arr1 , arr2 ){
        
    }

    function testcase(){
        function num_compare_output(arr1,arr2,result){
            console.log( arr1 , arr2 , result);
            console.log( num_compare( arr1 , arr2 ));
        }

        var arr1=[2,2,3];
        var arr2=[2,2,4];


        num_compare_output([1,2,3],[2,3,4],-1);       // -1
        num_compare_output([2,2,3],[2,3,4],-1);       // -1
        num_compare_output([1,2,3],[1,3,4],-1);       // -1
        num_compare_output(["a",2,3],[2,3,4],-1);     // -1
        num_compare_output([1,"a",4],[2,3,4],-1);     // -1
        num_compare_output([2,1,4],[2,1,5],-1);       // -1
        num_compare_output([2,1],[2,1,5],-1);         // -1
        num_compare_output(["a",2,"3"],[undefined,1,5],1);     // -1
        num_compare_output([],[2],-1);                 // -1

        num_compare_output(arr1.slice(),arr2.slice(),-1);
        console.log("original arr1 and arr2 :",arr1,arr2);

    }


    return {
        num_compare:num_compare,
        str_compare:str_compare,
        testcase:testcase
    };
}())


Comparsion.testcase();






