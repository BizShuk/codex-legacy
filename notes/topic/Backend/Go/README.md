# Go note

### array vs slice

- cap
- len
- append

##### array unawared copy

pass an array to function will result a new copy of array

##### Memory Leak problem

slice a existing array or a slice will cause underline array keep unnecessary elements as long as slice keep.

### type

```pointer receiver
func (p *int) method() { }
```

```value receiver
func (p int) method() { }
```

### Mutex

```go
type T Struct {
    lock Sync.Mutex
}
var t T
t.lock.Lock()
t.lock.Unlock()
```

### init function

when app runs, func init will be run ahead of main func.
It could have multiple init functions.

### fmt.Print

This will take Stringer interface(String method) to print.
However, if there is error interface(Error method) implement. error interface has higher priority.

## Project Layout

<https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/>
[Repo](https://github.com/eliben/modlib)

<https://github.com/golang-standards/project-layout>

## method goroutine

<https://stackoverflow.com/questions/36121984/how-to-use-a-method-as-a-goroutine-function>
