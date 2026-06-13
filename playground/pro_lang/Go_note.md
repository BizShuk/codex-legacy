Go note

### installation

by binary
1. download tar file and untar it
2. add bin to path

by source file
1.





### go lang doc
- ref. [manual](https://gobyexample.com/structs)
- ref. [user doc](http://astaxie.gitbooks.io/build-web-application-with-golang/content/en/03.2.html)
- ref. [link](http://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.1.html)

[Go example](https://gobyexample.com/)

###note:
	access uri with double / , ex: http://xxx.xxx.xxx.xxx//shuk_go , it'll send http status code 301 first


syntax:


		var var_name int;			=>		var var_name type;
		function

		func (class) func_name return_type {
		}

		if got class , then only used by object

example:

		package main
		import "fmt"
		This person struct type has name and age fields.
		type person struct {
		    name string
		    age  int
		}
		func main() {
		    fmt.Println(person{"Bob", 20})

		    fmt.Println(person{name: "Alice", age: 30})

		    fmt.Println(person{name: "Fred"})

		    fmt.Println(&person{name: "Ann", age: 40})

		    s := person{name: "Sean", age: 50}
		    fmt.Println(s.name)

		    sp := &s
		    fmt.Println(sp.age)

		    sp.age = 51
		    fmt.Println(sp.age)
		}

