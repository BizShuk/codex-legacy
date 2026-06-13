# Load Balance

[Pattern]: Load Balance Strategy

### Round Robin

Distribute requests evenly across all servers.

### Sticky Session

Client will connect to specific server only instead of random one.
When using web socket, this might be useful. There are some drawbacks as well

- Uneven distribution, it depends on algo to distribute connections to servers.
- When user sticks to specific server, it's more easier to overload one of servers in the same group.
