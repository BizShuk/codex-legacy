# Transaction

### Annotation

@EnableTransactionManagement

## @Transactional

### isolation

- DEFAULT
Use the default isolation level of the underlying database

- READ_COMMITTED
dirty reads are prevented; non-repeatable reads and phantom reads can occur.

- READ_UNCOMMITTED
a transaction may read data that is still uncommitted by other transactions.

- REPEATABLE_READ
dirty reads and non-repeatable reads are prevented; phantom reads can occur.

- SERIALIZABLE
dirty reads, non-repeatable reads, and phantom reads are prevented.

- `Dirty Reads`: Transaction "A" writes a record. Meanwhile, Transaction "B" reads that same record before Transaction A commits. Later, Transaction A decides to rollback and now we have changes in Transaction B that are inconsistent. This is a dirty read. Transaction B was running in READ_UNCOMMITTED isolation level so it was able to read Transaction A changes before a commit occurred.

- `Non-Repeatable Reads`: Transaction "A" reads some record. Then Transaction "B" writes that same record and commits. Later Transaction A reads that same record again and may get different values because Transaction B made changes to that record and committed. This is a non-repeatable read.

- `Phantom Reads`: Transaction "A" reads a range of records. Meanwhile, Transaction "B" inserts a new record in the same range that Transaction A initially fetched and commits. Later Transaction A reads the same range again and will also get the record that Transaction B just inserted. This is a phantom read: a transaction fetched a range of records multiple times from the database and obtained different result sets (containing phantom records).

### mode

- aspectj
- proxy

### proxy-target-class

- true, class-based proxies
- false, standard JDK interface-based proxies

### propagation

- REQUIRED
target method cannot run without an active tx

- REQUIRES_NEW
a newtxhas to start every time the target method is called

- MANDATORY ??? difference with REQUIRED???
target method requires an active tx to be running

- SUPPORTS
the target method can execute irrespective of atx

- NOT_SUPPORTED
the target method doesn’t require the transaction context to be propagated.

- NEVER
the target method will raise an exception if executed in a transactional process.

- NESTED

### timeout

### noRollbackFor

a rollback should not be issued if the target method raises this exception.

### rollbackFor

Default is rollbackFor=RunTimeException.class

### implmentation

- org.springframework.jdbc.datasource.DataSourceTransactionManager
- PlatformTransactionManager

TransactionTemplate

### What is Transaction manual work?
