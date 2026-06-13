# CAP theorem

aka `Brewer's theorem`

### Consistency

Even you connect to different app/node/operator, you still get the same result in last time you update it. Some approaches are eventuall consistent, like replica, async queue. It's weak consistency.

It's very hard and impossible to maintain strong consistency globally.
And it comes with cost if more stronger consistency.

### Availability

The client or user be able to access the data at all times, even if it is not consistent

### Partition Tolerance

This is different from the DB partition. This is network partition.

Even nodes are disconnected, the operation still sync once nodes connect back.

## Example

### Mongo, CP

Setup: Primary, Replica *2

Situation: If primary is out, two replicas will elect the primary out(check [Consensus Algo]). During the partition moment, system is not available until primary is elected.

MongoDB system behaves as a consistent system(Not strong consistency if query replica) and compromises on availability during a partition.

### Cassandra, AP

Cassandra is P2P system, consist of multiple nodes. It keeps multiple data replicas in seperate nodes.

Situation: When patition occurs, replicas are not updated. The data is still available, but just inconsistent. Cassandra provides [Eventual Consistency]. Divergent data will stay inconsistent until we update it.

[Consistent Hashing]:How does Cassandra choose node to store data?

### ElasticSeach

TODO: <https://discuss.elastic.co/t/elasticsearch-and-the-cap-theorem/15102/7>

### AWS DynamoDB

TODO: Check CAP with DynamoDB
