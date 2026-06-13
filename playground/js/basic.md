# Basic


- [event delegate]
- [bind]
    multiple bind will get first bind object, last of binding obj will not work.
- [call by value vs call by reference]
- [call & apply] , call faster but need to know exactly size of arguments 
- [Closure]()
- [Ajax](#ajax)


### reference
- [js tips](http://ejohn.org/apps/learn/)
- [js using patterns](http://shichuan.github.io/javascript-patterns/#function-patterns)
- [极客标签](http://www.gbtags.com/)
- [码农周刊](http://weekly.manong.io/issues/)
- [前端周刊](http://www.feweekly.com/issues)
- [极客头条](http://geek.csdn.net/)
- [Startup News](http://news.dbanotes.net/)
- [Hacker News](https://news.ycombinator.com/news)
- [InfoQ](http://www.infoq.com/)
- [w3cplus](http://www.w3cplus.com/)
- [Stack Overflow](http://stackoverflow.com/)
- [Atp](http://atp-posts.b0.upaiyun.com/posts/)
- [web開發筆試面試](http://mianshiti.diandian.com/)




### boolean

true:
- 1
- true
- "anyword"

false:
- 0
- false
- ""

null is object
undefined is undefined


###     call vs apply
        function theFunction(name, profession) {
            alert("My name is " + name + " and I am a " + profession + ".");
        }
        theFunction("John", "fireman");
        theFunction.apply(undefined, ["Susan", "school teacher"]);
        theFunction.call(undefined, "Claude", "mathematician");


### diff == vs ===
    "=="    =>  會自動作型別轉換, ex: false == "" == 0 weak type comparsion please check [boolean](#boolean)
    "==="   =>  絕對相等 type與value 皆須相同


### 優化手段
    for (var i = size; i < arr.length; i++) {}
    for 循环每一次循环都查找了数组 (arr) 的.length 属性，在开始循环的时候设置一个变量来存储这个数字，可以让循环跑得更快：
    for (var i = size, length = arr.length; i < length; i++) {}

    yahoo 14條優化手段
        http://inspiregate.com/internet/seo/62-yahoo-14-rules-to-accelerate-access-to-the-site.html
    高性能網站建構指南
        http://div.io/topic/371


### delegating(ie8以下不支援) and bubbling

### event trigger
trigger start from dom root (html -> body ->element) to child and bubble back

    當window接觸到event 會從window往下傳(delegating) 到了目標後會先觸發bubble event 然後才是delegated event 在往上傳(bubbling)
    e.preventDefault() => stop default event
        e.retrunValue = false , for ie
    e.stopPropagation() => stop bubbling
        e.cancelBubble = true , for ie

    jQuery
        delegate
            $("ul").delegate("li", "click", function(){
                $this).hide();
            });
            // 將ul下面的li都加上click event 即使是新增的



### Memory leak
- [js info](http://javascript.info/tutorial/memory-leaks)
- [MSDN](https://msdn.microsoft.com/en-us/library/ms976398.aspx)
- [js mem leak](http://blogger.gtwang.org/2014/01/javascript-memory-leak-patterns.html)

issue:

will cause leak or not?
```
    function A(){}
    A.prototype.x = function(){
        var y = copy of A
        return y
    }
    var a = new A();
    a = a.x()

```

### function arguments passed by 
Arguments are passed by Value
Objects are passed by Reference


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


### === Class ===
In JavaScript, class methods are not bound by default.


### string
```
str = "joe:x:100:100:see";
str.split(":")  => ["joe","x",100,100,"see"]
```

### number
All numbers in JavaScript are 64-bit (double-precision) floating point numbers.


### bits op
negative number right shift bits with 1 and postive with 0


# Browser
### eventhandler(e) 
e.preventDefault => prevent default event like a link or a form submit button



### call by value vs call by reference
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

### Closure
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



### DOM ###
Document Objectc Model

- `document.write("<content>")`

- `document.getElementById(<id>)` , elem
- `document.getElementsByTagName(<tag>)` , elems
- `document.getElementsByClassName(<className>)` , elems
- `document.querySelectorAll(<selector>)` , elems
- `document.createElemnt(<tag>)` , elem
- `elem.style.<style_attr>`
- 
- elem.innerHTML =  new html content   Change the inner HTML of an element
- elem.value = 
- elem.attribute = new value   Change the attribute value of an HTML element
- elem.setAttribute(attribute, value)  Change the attribute value of an HTML element
- elem.style.property = new style  Change the style of an HTML element

navigation:
- firstChild
- lastChild
- parentNode
- nextSibling
- previousSibling

[event](#event)
    

navigator:


### event ###
`elem.addEventListener(event,function,useCapture)` , useCapture => trigger when event capturing , false trigger when event bubbling.

attach different function to the same event without overide olds

`elem.onclick`

particular event list:
- onchange , An HTML element has been changed
- onclick , The user clicks an HTML element
- onmouseover , The user moves the mouse over an HTML element
- onmouseout , The user moves the mouse away from an HTML element
- onkeydown , The user pushes a keyboard key
- onload , The browser has finished loading the page

### Ajax ###
Asynchronouse Javascript And Xml

pure ajax
```
// XMLHttpRequest.DONE


var xmlhttp = new XMLHttpRequest();
xmlhttp.open("<http method>","<url>",async);
xmlhttp.onreadystatechange = function() {
    xmlhttp.readyState 
    xmlhttp.status     // http status
    xmlhttp.statusText // for error code
}
xmlhttp.send(dataobj)
```

### Object-Orieted

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




### special cases

```javascript
    function printing() {
       console.log(1);
       setTimeout(function() { console.log(2); }, 1000);
       setTimeout(function() { console.log(3); }, 0);
       console.log(4);
    }
    printing(); // result: 1 4 3 2
```

```
    (function() {
        var a = b = 5;  // a: local , b: global
        c = d = 6;      // c,d: global
    })();
    console.log(b,c,d); // 5 6 6 , a 為local 會炸
```

spread operation:
```
function fn() {

    console.log(Array.from(arguments));
}

a = [1, 2, 3];
fn(...a);
```


```
    console.log(typeof null);           // object
    console.log(typeof {});             // object
    console.log(typeof []);             // object ,but [] instanceof Array
    console.log(typeof undefined);      // undefined
    console.log(typeof new Array());    // object
    console.log(typeof function(){});   // function
    console.log(typeof new Date())      // object


    instanceof , check prototype chain
    new String() instanceof String
```


```
        NaN   ==  NaN   =>  false   ,特殊狀況NaN不屬於任何值 沒有值=>永遠不會相等
        NaN   === NaN   =>  false   ,特殊狀況NaN不屬於任何值 沒有值=>永遠不會相等
        true  ==  '1'   =>  true
        true  ==  '5'   =>  false
        false ==  '0'   =>  true
        true  === '1'   =>  false
        true  === '5'   =>  false
        false === '0'   =>  false
```

```
    console.log(myglobal); // "hello"
    console.log(window.myglobal); // "hello"
    console.log(window["myglobal"]); // "hello"
    console.log(this.myglobal); // "hello"
```

```
    var month = "06",
        year = "09";
        month = parseInt(month, 10);    // 6
        year = parseInt(year, 10);      // 9
        month = parseInt(month);    // 0
        year = parseInt(year);      // 0
```


why?
```

            // antipattern
            $(elem).data(key, value);
            // preferred
            $.data(elem, key, value);
```

```
0.1 +0.2 === 0.3  // false
function multFloats(a,b){
  var atens = Math.pow(10,String(a).length - String(a).indexOf('.') - 1), 
      btens = Math.pow(10,String(b).length - String(b).indexOf('.') - 1); 
  return (a * atens) * (b * btens) / (atens * btens); 
}
```

why?
```
    var scareMe = function () {
                alert("Boo!");
                scareMe = function () {
                    alert("Double boo!");
                };
            };
            // 1. adding a new property
            scareMe.property = "properly";
            // 2. assigning to a different name
            var prank = scareMe;
            // 3. using as a method
            var spooky = {
                boo:scareMe
            };
            // calling with a new name
            prank(); // "Boo!"
            prank(); // "Boo!"
            console.log(prank.property); // "properly"
            // calling as a method
            spooky.boo(); // "Boo!"
            spooky.boo(); // "Boo!"
            console.log(spooky.boo.property); // "properly"
            // using the self-defined function
            scareMe(); // Double boo!
            scareMe(); // Double boo!
            console.log(scareMe.property); // undefined
```


```
var foo = 1;

if (True){
    var foo =3;
}
console.log(foo); // 3
```