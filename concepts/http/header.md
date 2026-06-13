# Http Header


### Header properties
- Cache-Control, e.g. no-cache,max-age=30
    - no-store
    - no-cache
    - public
    - private
    - max-age
- Content-Length
- Content-Type
    - application/json
    - application/xml
    - text/xml
    - text/plain
    - text/html
- Content-Language
    - en
- Content-Encoding
    - gzip
- Server 
    - openresty
    - nginx
    - Microsoft-IIS/8.5
- Accept , 
- Accept-Charset , e.g. UTF-8 
- Accept-Encoding , e.g. gzip , deflate
- Accept-Language , e.g. zh-TW,zh-CN;q=0.6,en-US;q=0.2
    - q , weight
    - languages
- Authorization , [basic](https://en.wikipedia.org/wiki/Basic_access_authentication)
    - use `user:passwd@` before domain ,e.g. http://user:passwd@shuk.info
    - use header `Authorization: Basic base64(<user>:<passwd>)`
- X-..... , 非標準使用Header
    - X-Powered-By , what kind of web service tech
    - X-Version , web service version
- X-Frame-Options , Policy for that page is embeded to another one
    + DENY
    + SAMEORIGIN
    + ALLOW-FROM [uri]


### Cache
[ ]check flow
1. send request:  
check cache expire time 
    - unexpired , use cache
    - expired , use cache data to `ETag` and `Last-Modified`
2. get response  
    - 200 , save cache headers to cache
    - 304 , using cache


- ETag , 內容驗證權杖
- Last-Modified , 時間驗證權杖
- If-None-Match , 上次ETag時間
- If-Modified-Since , 上次Last-Modified時間


### Out of date
-   Pragma ,  useless after http1.1
    - no-cache
- Expires , useless after http1.1
    - `<time>`


