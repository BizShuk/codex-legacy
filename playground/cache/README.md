# Cache

[Document from Amazon](https://aws.amazon.com/caching/implementation-considerations/)



### The Thundering Herd aka dog piling
when a key expired, identical query send by many users in the same time. This also happens in adding a new cache node.

##### solution
preswamp, 預先取得 15分鐘expired 就在14分鐘先取 可以有不同的策略 最後1分鐘有read則更新之類的

    1. Write a script that performs the same requests that your application will. If it's a web app, this script can be a shell script that hits a set of URLs.
    
    2. If your app is set up for lazy caching, cache misses will result in cache keys being populated, and the new cache node will fill up.
    
    3. When you add new cache nodes, run your script before you attach the new node to your application. Because your application needs to be reconfigured to add a new node to the consistent hashing ring, insert this script as a step before triggering the app reconfiguration.
    
    4. If you anticipate adding and removing cache nodes on a regular basis, prewarming can be automated by triggering the script to run whenever your app receives a cluster reconfiguration event through Amazon Simple Notification Service (Amazon SNS).


### Strategy

##### Naming Convention
A hierarchal structure separated by colons is a common naming convention for KEYs, such as object:type:id

##### lazy caching
read cache first. If miss, read from storage and set the cache

##### write back
modify data in cache and mark it as dirty. when it is evicted or after a period of time, write this back to storage.

##### write throught
Every time modify the cache, modify the storage as well

#####  Russian doll caching
[The performance impact of "Russian doll" caching](https://signalvnoise.com/posts/3690-the-performance-impact-of-russian-doll-caching)



### eviction policy

- volatile-lru eviction policy
- random-based eviction policy
- TTL (time to live) , don't use the same time period for all key => stagger the re-query time window