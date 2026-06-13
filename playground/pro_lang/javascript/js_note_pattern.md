
### Object orieted

##### # inherited
parent constructor => child constructor


funciton() vs new function()
    ex:

    function person(){
        console.log(1);
        this.a = function(){    console.log(2); }
        this.b = function(){    console.log(3); }
        return 123;
    }

    person();   // output: 1 , return "123"
    var tmp = new person(); // output: 1 , return person object
    tmp.a();    // output: a
    tmp.b();    // output: b


Closure
    var nodes = document.getElementsByTagName('button');
    for (var i = 0; i < nodes.length; i++) {
       nodes[i].addEventListener('click', (function(i) {
          return function() {
             console.log('You clicked element #' + i);
          }
       })(i));
    }
    You clicked element #+nodes.length  // 不是 正確的i

    solution:
        function handlerWrapper(i) {
           return function() {
              console.log('You clicked element #' + i);
           }
        }

        var nodes = document.getElementsByTagName('button');
        for (var i = 0; i <nodes.length; i++) {
           nodes[i].addEventListener('click', handlerWrapper(i));
        }



global vs local
    (function() {
        var a = b = 5;
        c = d = 6;
    })();
    console.log(b,c,d); // 5 6 6 , a 為local 會炸

Object
    add method 並避免重複寫
        String.prototype.repeatify = String.prototype.repeatify || function(times) {}

    this的運用
        var fullname = 'John Doe';
        var obj = {
            fullname: 'Colin Ihrig',
            prop: {
                fullname: 'Aurelio De Rosa',
                getFullname: function() {
                    return this.fullname;
                }
            }
        };
        console.log(obj.prop.getFullname());
        var test = obj.prop.getFullname;
        console.log(test());

        // Aurelio De Rosa
        // John Doe

    call vs apply
        function theFunction(name, profession) {
            alert("My name is " + name + " and I am a " + profession + ".");
        }
        theFunction("John", "fireman");
        theFunction.apply(undefined, ["Susan", "school teacher"]);
        theFunction.call(undefined, "Claude", "mathematician");


call by value vs call by reference
    ex:

    var num=10,name="jason",obj1={value:"first"},obj2={value:"second"},obj3=obj2;
    console.log(num,name,obj1,obj2,obj3);   // 10,"jason",{value:"first"},{value:"second"},{value:"second"}
    function change(num,name,obj1,obj2){
        num = num*10;
        name = "Liam";
        obj1 = obj2;
        obj2.value = "new"
    }
    change(num,name,obj1,obj2);
    console.log(num,name,obj1,obj2,obj3);   // 10,"jason",{value:"first"},{value:"new"},{value:"new"}

