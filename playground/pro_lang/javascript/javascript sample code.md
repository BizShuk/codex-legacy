javascript sample code

```
// return true if the property is not defined by Object
// retrun false by length , default function name
obj.hasOwnProperty(property);

```


global scope
```
	var i = "global"
	function f_scope(){
		console.log(i);	// undefined
	}
	{
		console.log(i);	// global
	}


	function sum(x, y) {
				// implied global
				result = x + y;
				return result;
	}
	sum(1,2);
	console.log(result);	// x+y 


	myglobal = "hello"; // antipattern
	console.log(myglobal); // "hello"
	console.log(window.myglobal); // "hello"
	console.log(window["myglobal"]); // "hello"
	console.log(this.myglobal); // "hello"


	function foo() {
		var a = b = 0;	// equal var a = ( b = 0 )
	}
	fool();
	console.log(a,b);	// undefined , 0
```


```
	var month = "06",
		year = "09";
		month = parseInt(month, 10);	// 6
		year = parseInt(year, 10);		// 9
		month = parseInt(month);	// 0
		year = parseInt(year);		// 0
```


why?
```

			// antipattern
			$(elem).data(key, value);
			// preferred
			$.data(elem, key, value);
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
