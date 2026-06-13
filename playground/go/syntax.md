# Syntax

### Slice

- [Slice trick](https://github.com/golang/go/wiki/SliceTricks)

use append like this
```
a = append(a[:i], append(make([]T, j), a[i:]...)...)
a = append(a, make([]T, j)...)
a = append(a[:i], append([]T{x}, a[i:]...)...)
```
Above will cause memory garbage. It could be avoid by following
```
s = append(s, 0)
copy(s[i+1:], s[i:])
s[i] = x
```



### type struct
```
cases := []struct { 
    in, want string    
}{
    {"Hello, world", "dlrow , olleH"},
        {"Hello, world", "dlrow , olleH"},
        ...
}
```

##### type alias
`type <new type> <existing type>>`



### for

##### iteration
```
for _,c := range cases {
    // _ = index    
    // c = value

}
```

##### while
```
for a<b{

}
```


### Print Variable
- %v , value => only property value
- %+v , plus value => with key and property value
- %#c , shaped value => go format
- %s , string
- %d , digit

