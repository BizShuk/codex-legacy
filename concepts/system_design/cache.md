# Cache 

### tools
- [Memcached](https://en.wikipedia.org/wiki/Memcached)
- [Redis]()


# Policy
[Cache Replacement Policy](https://en.wikipedia.org/wiki/Cache_replacement_policies)

### LRU 
Least Recently Used , add a byte to each elements ( age byte ) and replace the minimum byte element.

### MRU
Most Recently Used , add a byte to each elements ( age byte ) , and replace the maxmum byte element. Update if there is the same one  as well.