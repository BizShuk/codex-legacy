# Distributed Lock

In case of race condition, we need locks for resources. Resources usually won't interfere other resources. So the lock should be bind only the resource we need.

Some cases different, Multiple instances would like to elect master node. All nodes will be using the same lock to let only one node become master node.

There are few concepts required in distributed lock system

- TTL, to prevent node drop off and resources is still locking.
- Unique ID of lock
- Lock Manager?
- ZooKeeper?
- Priority, => resolve deadlock
- consensus algo

TODO: Distributed Lock system
