# Service Discovery

- Service register itself to service discovery
- healthcheck to maintain the node on service discovery

### Server-side

> App servers register to service discovery. Load Balancer request server list from service discovery. Client just send queries to load balancer.

pros: simple and easy for client
cons: one more layer

### Client-side

> App servers register to service discovery. Client request server list from service discovery and select one of them to query.

pros: peer-to-peer
cons: coupling some logics inside client library
