# Mocking in Go

### Package Function

```go
var localFunc = package.Func

func () {
    localFunc() // not package.Func
}
```

### Package Component

> Create a struct than contain a interface component
> => when testing, you can assign mock component when
> constructing it

### Always Create interface for external library

> If you have a struct, create an interface for it.
> External user can use it for their mocking on its functions
