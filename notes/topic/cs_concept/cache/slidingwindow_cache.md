# Sliding Window cache

Some data only needs to keep for 1 hour. If just using single value to address it, it's not enough.

### Solution 1

Using Linked List to append each instance with timestamp on a list.

pros: it can check every request whether they are fulfilled.
cons: resources is depended on the size of queue.

### Solution 2

Keep records in DB and use a process to update cache value periodically. Ex: `select count(*) from video_read where timestamp > 111111111`. Those data can be clean once it past configured hours.

pros: limited resource
cons: not real-time update.

### Solution 3, Solution 1 improvement

Adjust time granularity from every request to every second/minute/hour
