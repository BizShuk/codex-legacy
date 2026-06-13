# Lock in Java

[Violatile]: Read/Write on main memory only

[synchronized]: Get value from main memory when it enters the block and flush the result into memory when exit the block. It can be use on methods/blocks. Used on static method is acquiring lock on class, non-static method is on object.

When to use Volatile over Synchronized modifiers can be summed up into this: Use Volatile when you variables are going to get read by multiple threads, but written to by only one thread. Use Synchronized when your variables will get read and written to by multiple threads.

### java concurrent package

Achieved by mutex.

```java
List list = new ArrayList<>();
Map map = new HashMap<>();
list = Colllections.synchronizedList(list);
map = Collections.syncrhonizedMap(map);
```
