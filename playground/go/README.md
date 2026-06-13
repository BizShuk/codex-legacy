# Golang

### undoc
- excution priority , global variable -> init() -> main()

### pass by 
Go is always passed by Value

### Outline
- [Debug](debug.md)
- [Language Spec](language-spec.md)
- [Optimization](#optimization)
- [Testing](#testing)
- [Tools](tools.md)
- [Syntax](syntax.md)




### ref
- [Wiki](https://github.com/golang/go/wiki) , Working with Go ,Learning more about Go
- [IDE, Edit](https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins) , [Vim-go](https://github.com/fatih/vim-go)
- [git2go tutorial](https://blog.gopheracademy.com/advent-2014/git2go-tutorial/)
- [Go study](go/)

### Exercise
- [A tour of Go](https://tour.golang.org/welcome/1)
- [Go examples](https://gobyexample.com/)
- [Go Concurrency Visualizing](http://divan.github.io/posts/go_concurrency_visualize/)


### Language spec
- [Functions Types in Go](http://jordanorelli.com/post/42369331748/function-types-in-go-golang)
- [Language Specification](https://golang.org/ref/spec)


# Optimization
1. use type struct is 10 times excution time than basic unit
2. `make([]int,100000)` = make([]int,0,100000)


### inlining
`func funcA(a,b){ return 2*utils.Max(a,b) }
會被拆程
```
func funcA(){
    tmp :=a
    if b>a{
        tmp =b
    }
    return 2*tmp
}
```

不會產生新的stack 就不會overwrite register

### hot split
```
func A()
    for _,v := range vals{
        H(v)
    }
```
多建一大塊stack在heap跟stack中間,原本位置就讓給正常的用


### slice
don't neet to `make([]slice,0)` or `a:=[]slice{}`, default slice value is nil
so just `var a []slice`




# Testing

### Go test
1. file name end with \_test.go
2. func name will be `func Test<Tested_func_name>(t *testing.T)`
3. if call `t.Fail` or `t.Error` , each means failed
