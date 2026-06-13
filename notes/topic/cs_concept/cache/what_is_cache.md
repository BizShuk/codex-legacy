# What is Cache?

Cache is a way to increase data retrival rate. Such as memory, DNS, CDN(AWS CloudFront), and HTTP Cache are caches. It helps to improve app performance, reduce DB cost, reduce loading, eliminate DB hotspots, increse read throughput.

### Design Considerations

- [TTL], time to live with range, against data staleness
- [Hit Rate], [Miss Rate](penalty)
- [Cache Penetration] => [Fake Cache data] [Limited User Access]
- pre-fetch [Refresh-Ahead] [Scripting] [Write-Through]
- [Eviction], when memory is over filled. [TTL] [Use time] []
- [Thundering herd], `Dog piling`, multiple applications hit cache and miss, and go to database.

### Dedicated Tools

- Redis
- MemCached
