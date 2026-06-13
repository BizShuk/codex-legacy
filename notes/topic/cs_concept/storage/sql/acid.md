# ACID

### Atomicity

Any transaction updated multiple rows is treated as a single unit, either update all, or all failed.

### Consistency

### Isolation

- [Read Uncommitted], 最低级别的隔离，是读取已提交+新的隔离突破。如果事务A读取了数据D，然后数据D被事务B修改（但并未提交，事务B仍在运行中），事务A再次读取数据D时，数据修改是可见的。如果事务B回滚，那么事务A第二次读取的数据D是无意义的，因为那是事务B所做的从未发生的修改（已经回滚了嘛）。
这叫脏读（[Dirty Read]）。

- [Read Committed]，Oracle、PostgreSQL、SQL Server default, 可重复读+新的隔离突破。如果事务A读取了数据D，然后数据D被事务B修改（或删除）并提交，事务A再次读取数据D时数据的变化（或删除）是可见的。
这叫不可重复读 [Non-Repeatable Read]。

- [Repeatable Read]，Mysql default, 每个事务有自己的『世界』，除了一种情况。如果一个事务成功执行并且添加了新数据，这些数据对其他正在执行的事务是可见的。但是如果事务成功修改了一条数据，修改结果对正在运行的事务不可见。所以，事务之间只是在新数据方面突破了隔离，对已存在的数据仍旧隔离。
举个例子，如果事务A运行”SELECT count(1) from TABLE_X” ，然后事务B在 TABLE_X 加入一条新数据并提交，当事务A再运行一次 count(1)结果不会是一样的。
这叫幻读 [Phantom Read]。

- [Serializable]，SQLite default, 最高级别的隔离。两个同时发生的事务100%隔离，每个事务有自己的『世界』。
-

### Durability

When server is down and back again. The data still remain the same.

[WAL]
