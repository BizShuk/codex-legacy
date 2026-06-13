# Data Center Replication

How to replicate data across multiple data centers (DCs)?

This question is opposite of data collision in hash table. And DC is located in some geo regions. Some might be close to one or another. The replication should depend on business as well. EX: Close region has higher priority to replicate.

From data collision, we can know some approaches.

[Open Hashing]

1. [Separate Chaining], This doesn't fulfill the replication requirement.

[Closed Hashing], [Open Addressing]

2. [Linear Probing]
3. [Quadratic Probing]
4. [Double hashing]

[Consisten Hashing] is also a good approach similiar to Linear Probing. It provides more scalability/availability and less migration effort.
