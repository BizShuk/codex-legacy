# Mysql vs Postgre

### Difference

- Postgres is an object-relational database, while MySQL is a purely relational database. This means that Postgres includes features like table inheritance and function overloading, which can be important to certain applications.
- Postgres also adheres more closely to SQL standards.
- Postgres handles concurrency better than MySQL for multiple reasons.
    Postgres implements [Multiversion Concurrency Control] ([MVCC]) without read locks Postgres supports parallel query plans that can use multiple CPUs/cores Postgres can create indexes in a non-blocking way (through the CREATE INDEX CONCURRENTLY syntax), and it can create partial indexes (for example, if you have a model with soft deletes, you can create an index that ignores records marked as deleted) Postgres is known for protecting data integrity at the transaction level. This makes it less vulnerable to data corruption.
- MySQL has had a reputation as an extremely fast database for read-heavy workloads, sometimes at the cost of concurrency when mixed with write operations.

### From articles

- Postgres, advertises itself as “the most advanced open-source relational database in the world”
- Postgres performance was more balanced - reads were generally slower than MySQL, but it was capable of writing large amounts of data more efficiently, and it handled concurrency better.
