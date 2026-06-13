# DEBUG

### invalid type for composite literal



### why does copy slice need twice cap()
```go
t := make([]byte, len(s), (cap(s)+1)*2)
copy(t, s)
s = t
```

### slice copy by value
`copy` build-in function

### slice always pass pointer
append slice to another variable will cause slice pieces
ex:
1,2,3,4
1,2,3,5   4
1,2,3,6   4 5
I'm not sure that it'll be gc or not. But it should be after end of function
check [here](https://play.golang.org/p/DUe3qME6c_)


### int to string 
use strconv.Itoa(123)

### Could not determine kind of name for C.xxxxxx
A: Need feader file , shared object file , ldconfig , pkg-config path(*.pc file in pkg_config_path )
