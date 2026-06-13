# Innod DB Lock

[15.7.3 Locks Set by Different SQL Statements in InnoDB](https://dev.mysql.com/doc/refman/8.0/en/innodb-locks-set.html)

### Terminology

[Shared Lock(S)]
[Exclusive Lock(X)]
[Intension Shared Lock(IS)], table lock for later S lock
=> SELECT ... FOR SHARE
[Internsion Exclusive Lock(IX)], table lock for later X lock
=> SELECT ... FOR UPDATE

[Record Lock], Lock index record, if no index, hidden clustered index created by default [Section 15.6.2.1, “Clustered and Secondary Indexes"](https://dev.mysql.com/doc/refman/8.0/en/innodb-index-types.html)

[Gap Lock], Lock on a gap between index records, or a lock on the gap before the first or after the last index record.
ex: SELECT c1 FROM t WHERE c1 BETWEEN 10 and 20 FOR UPDATE

only purpose is to prevent other transactions from inserting to the gap

disable by change isolation level to RC.
disable by searches and [index scans] .
used for foreign-key constraint checking and duplicate-key checking.

[Next-Key Locks] combination of a record lock on the index record and a gap lock on the gap before the index record.

[Predicate Locks],

non-indexed
RC , lock all, release once valid
RR , lock all, relase only tx finished.

By default, InnoDB operates in RR isolation level. In this case, InnoDB uses next-key locks for [searches] and [index scans], which prevents phantom rows [see Section 15.7.4, “Phantom Rows”](https://dev.mysql.com/doc/refman/8.0/en/innodb-next-key-locking.html)

[AUTO-INC Locks]

[Lock Comparibility]
A lock is granted to a requesting transaction if it is compatible with existing locks

### Consisten Non-Blocking Read

[Consistent non-blocking read]
A read operation that uses snapshot information to present query results based on a point in time, regardless of changes performed by other transactions running at the same time.
RR: time of first read
RC: time of each read
Exception: if update in earlier statement in the same transaction, later select will see the changes and as well as changes in other tx

DML won't be read from snapshot. If other tx commited few rows which might get deleted in this DML

for INSERT INTO ... SELECT, UPDATE ... (SELECT), and CREATE TABLE ... SELECT
by default, isolation level is RC

[Locking Read]
    query data and then insert or update related data within the same transaction, the regular SELECT statement does not give enough protection. Other transactions can update or delete the same rows you just queried.

    Bad:
            SELECT ... FROM  <=== what you select might be updated/deleted before this tx commited.
            INSERT ...
        
    Solution:
        [SELECT ... FOR SHARE, can read but not DML]
        [SELECT ... FOR UPDATE]


    Bad:
        SELECT counter_field FROM child_codes FOR SHARE;
        UPDATE child_codes SET counter_field = counter_field + 1;
        => read the same counter number, but failed to update due to deadlock
           use FOR UPDATE instead

    [SELECT LAST_INSERT_ID();]
    The SELECT statement merely retrieves the identifier information (specific to the current connection)


    [NOWAIT]
        A locking read that uses NOWAIT never waits to acquire a row lock. The query executes immediately, failing with an error if a requested row is locked.

    [SKIP LOCKED]
        A locking read that uses SKIP LOCKED never waits to acquire a row lock. The query executes immediately, removing locked rows from the result set.

    NOWAIT and SKIP LOCKED only apply to row-level locks.

[UPDATE ... WHERE ...] sets an exclusive next-key lock on every record the search encounters. However, only an index record lock is required for statements that lock rows using a unique index to search for a unique row.

[Phantom Rows](https://dev.mysql.com/doc/refman/8.0/en/innodb-next-key-locking.html)

Contention-Aware Transaction Scheduling (CATS) MySQL 8.0.20
    vs
First In First Out (FIFO) algorithm

### Deadlock

InnoDB tries to pick small transactions to roll back, where the size of a transaction is determined by the number of rows inserted, updated, or deleted
