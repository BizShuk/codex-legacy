# Clover

Clover is user migration project based on region. User have a store country which is mapping into  a IDC.

The Clover is targeting the EU user moving from SG in to GCP(useast2a). The existing migration is working but there is some enhancement.

1. Consistency. The original design is kind of multi-thread programming, very hard to debug logically. The migration potentially exists a rollback issue(race condition). By adding the redis CAS command to prevent the overwrite(optimistic locking, keep migration stage always move from previous stage to next)
2. Newly joined migration data is not existing the aligned IDC. The data migration source was only from one IDC to another. But for refund data, it is in the other IDC in the VGeo location.
3. Complexity is high. The older design is based on one-to-one IDC. There is no clear thinking about multiple IDCs interactions and configurations.

These leads to re-design the architecture of single set migration system.

The Clover migration was started 5/4. It was facing a lot of issues.

- Infinite loop
- Wrong data operations
- MQ hive accumulation
- C&S bug on re-insertion in new cycle
- dbatman performance issue
- DB performance issue(split shard by double the existing shards)

In the end of migration, it archives 120 users/s, 100k qps in target IDC.
wallet_shard_db around avg 60% CPU.
