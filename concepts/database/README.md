# Database


- [果有人问你数据库的原理，叫他看这篇文章](http://blog.jobbole.com/100349/)
- [Architecture of a Database System](http://db.cs.berkeley.edu/papers/fntdb07-architecture.pdf)

# UUID vs GUID
mysql support UUID
detail check later when I get time

# NOSQL

### feature
- distributed
- schema-fee
- eventually consistent
- [CAP](https://blog.longwin.com.tw/2013/03/nosql-db-choose-cap-theorem-2013/)
    + Consistency
    + Availability
    + Partition tolerance

- [Mongo](database/mongo.md)


# SQL

### feature

- ACID , 更新資料的過程中，為保證事務（transaction）是正確可靠的，所必須具備的四個特性
    + Atomicity
    + Consistency
    + Isolation
    + Durability.



# crash solution
- undo
- redo






### SQL一般定义4个隔离级别

- 串行化(Serializable，SQLite默认模式）：最高级别的隔离。两个同时发生的事务100%隔离，每个事务有自己的『世界』。

- 可重复读（Repeatable read，MySQL默认模式）：每个事务有自己的『世界』，除了一种情况。如果一个事务成功执行并且添加了新数据，这些数据对其他正在执行的事务是可见的。但是如果事务成功修改了一条数据，修改结果对正在运行的事务不可见。所以，事务之间只是在新数据方面突破了隔离，对已存在的数据仍旧隔离。
举个例子，如果事务A运行”SELECT count(1) from TABLE_X” ，然后事务B在 TABLE_X 加入一条新数据并提交，当事务A再运行一次 count(1)结果不会是一样的。
这叫幻读（phantom read）。

- 读取已提交（Read committed，Oracle、PostgreSQL、SQL Server默认模式）：可重复读+新的隔离突破。如果事务A读取了数据D，然后数据D被事务B修改（或删除）并提交，事务A再次读取数据D时数据的变化（或删除）是可见的。
这叫不可重复读（non-repeatable read）。

- 读取未提交（Read uncommitted）：最低级别的隔离，是读取已提交+新的隔离突破。如果事务A读取了数据D，然后数据D被事务B修改（但并未提交，事务B仍在运行中），事务A再次读取数据D时，数据修改是可见的。如果事务B回滚，那么事务A第二次读取的数据D是无意义的，因为那是事务B所做的从未发生的修改（已经回滚了嘛）。
这叫脏读（dirty read）。


### 并发控制
read and write the same datas in the same timea

- 监控所有事务的所有操作
- 检查是否2个（或更多）事务的部分操作因为读取/修改相同的数据而存在冲突
- 重新编排冲突事务中的操作来减少冲突的部分
- 按照一定的顺序执行冲突的部分（同时非冲突事务仍然在并发运行）
- 考虑事务有可能被取消



### lock
悲观锁

原理是：

如果一个事务需要一条数据
它就把数据锁住
如果另一个事务也需要这条数据
它就必须要等第一个事务释放这条数据
这个锁叫排他锁。
但是对一个仅仅读取数据的事务使用排他锁非常昂贵，因为这会迫使其它只需要读取相同数据的事务等待。因此就有了另一种锁，共享锁。

共享锁是这样的：

如果一个事务只需要读取数据A
它会给数据A加上『共享锁』并读取
如果第二个事务也需要仅仅读取数据A
它会给数据A加上『共享锁』并读取
如果第三个事务需要修改数据A
它会给数据A加上『排他锁』，但是必须等待另外两个事务释放它们的共享锁。
同样的，如果一块数据被加上排他锁，一个只需要读取该数据的事务必须等待排他锁释放才能给该数据加上共享锁。



死锁

但是使用锁会导致一种情况，2个事务永远在等待一块数据：


两段锁协议（Two-Phase Locking Protocol)


更多类型的锁（比如意向锁，intention locks）和更多的粒度（行级锁、页级锁、分区锁、表锁、表空间锁）
