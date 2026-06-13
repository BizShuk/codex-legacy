# Struct Conversion

Only if underlying data structure are the same. [Reference](https://www.sohamkamani.com/golang/type-assertions-vs-type-conversions/#why-is-it-an-assertion)

```go
    type person struct {
        name string
        age int
    }

    type child struct {
        name string
        age int
    }

    type pet {
    name string
    }

    func main() {
        bob := person{
            name: "bob",
            age: 15,
            }
    babyBob := child(bob)
    // "babyBob := pet(bob)" would result in a compilation error
        fmt.Println(bob, babyBob)
    }
```
