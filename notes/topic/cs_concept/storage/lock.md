# DB Lock

### Pessimistic Lock

Lock everytime using resources. It almost equals to DB Transaction depends on Isolation level.

### Optimistic Lock

Assume no one will update. But if someone tries to update, the person/query will check whether the data has changed since last query. By version/[CAS]

Suitable for non-intense resource.

ABA issue: Fetch in first A, but update in second A.
Loopback issue: Retry until successful[Spin Lock]. High cost if intensive update.

Example: `Update [table] Set column1=[new value] Where column1=[old value]`

### Comapre and Swap, [CAS]

Part of Optimistic lock.

### Shared Lock

Similiar to [Read-Write Lock]
如果一个事务只需要读取数据A
它会给数据A加上『共享锁』并读取
如果第二个事务也需要仅仅读取数据A
它会给数据A加上『共享锁』并读取
如果第三个事务需要修改数据A
它会给数据A加上『排他锁』，但是必须等待另外两个事务释放它们的共享锁。
同样的，如果一块数据被加上排他锁，一个只需要读取该数据的事务必须等待排他锁释放才能给该数据加上共享锁。

### Update Lock

如果一个事务需要一条数据
它就把数据锁住
如果另一个事务也需要这条数据
它就必须要等第一个事务释放这条数据
这个锁叫排他锁。
但是对一个仅仅读取数据的事务使用排他锁非常昂贵，因为这会迫使其它只需要读取相同数据的事务等待。因此就有了另一种锁，共享锁。

### Dead Lock

但是使用锁会导致一种情况，2个事务永远在等待一块数据：

### Two-Phase Lock

### More

更多类型的锁（比如意向锁，intention locks）和更多的粒度（行级锁、页级锁、分区锁、表锁、表空间锁）
