# database

# TODO: clean up

# document db

### pros

one to many
better for analytics by recording when event occurs.
schema flexibility
better performance due to locality

### cons

nested item, not a problem if not too deep
additional work to keep denormalized data consistent

### term

document refernece = foreign key

schema-on-read = dynamic checking in PL
    if many different objects use the same data
    if you have no control of structure of data, maybe external.
schema-on-write= static, like update mysql schema

# data locality

Fast access, not liimited in Document DB. If only small portion of data is needed, keep document small.

- multi-table index cluster tables in Oracle
- column-family concept in bigtable

# Relational db

### pros

more support on many-to-many and many-to-one

### Mysql

Alter table copy entire table and it takes minutes or hours. Only few milliseconds in other RDB.

### Mongo

aggregation pipeline

# Graph-Like Data Models

- vertices (nodes/entities)
- edges(relationships/arcs)

Hash
LSM-Tree
SSTable
B-Tree

Recovery Strategy

- WAL
- COR
