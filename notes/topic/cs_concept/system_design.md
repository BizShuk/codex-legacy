# System Design

It covers multiple areas.

### Availability

1. More instances and using load-balancer
2. Store more copies in multiple nodes
3. Create replicas for each node
4. Backup outside of region
5. Circuit break

### Scalability

Aspects: Application, System(Infra), Team

1. Leverage cloud solution to scale in/out
2. Make code/process more independant
3. Create idempotent service
4. Master-Slave, Replica
5. Cache to mitigate loading of DB
6. CDN mitigate loading of network traffic and reduce latency
7. setup TTL for content
8.

### Application Tuning

- N+1 problem, The N+1 query problem happens when the data access framework executed N additional SQL statements to fetch the same data that could have been retrieved when executing the primary SQL query.
- Dependant sequential call/object
- Multi-level nested query/iteration (Multiple enumerations)
