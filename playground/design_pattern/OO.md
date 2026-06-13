# Object-Oriented
4 features of OO

### encapsulation 

|       |class    |package  |subclass |anywhere |
|-------|---------|---------|---------|---------|
|public |    Y    |    Y    |    Y    |    Y    |
|protect|    Y    |    Y    |    Y    |    N    |  
| none  |    Y    |    Y    |    N    |    N    |
|private|    Y    |    N    |    N    |    N    |

final : 成員 (member) 宣告時若使用關鍵字 final 修飾，表示該成員為常數 (constant) ，屬性 (field) 或區域變數 (local variable) 都不能重新給值，而方法 (method) 繼承 (inherit) 後不可被改寫 (override) 。

### abstraction

### inheritance
我們都知道就是父子關係(IS-A relationship),子類別會繼承父類的方法和屬性,繼承本身並沒有問題,但常會見到誤用的情形!就是開發者忘了IS-A的概念,為了reuse,造成了功能型的繼承!若要reuse且非IS-A關係,應用Composition來達成,把共同的部份移至composed class,把要做的事delegate另一個物件完成!

### polymorphism
延申自繼承(Inheritance)或介面(Interface),指的就是不同型態的物件,定義相同的操作介面,由於被呼叫者(Callee)有著相同介面,呼叫者並不用指定特定型別,只需針對介面進行操作,實際執行的物件則在runtime決定,藉此增加程式碼的彈性。 


### Generic
傳入的參數可以動態型別
 types to-be-specified-later that are then instantiated when needed for specific types provided as parameters


- [C++](c_cpp/generic.cpp)

### Overloading
same method with different parameters

```
func show(a int)
func show(b string)
```

Dont Support: Go


### Resource Sharing
intently to get right resource, you need to lock resource.

##### Race condition
```
var i int = 0
func addone() int {
    return i++
}
```
=>
```
T1:1 , get from i = 0
T2:1 , get from i = 0
T1:2 , get from i = 1
T2:3 , get from i = 2
```