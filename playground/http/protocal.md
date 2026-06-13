# http protocal

POST /luajit.lua HTTP/1.1
Host: 192.168.2.52
Content-type: text/plain
Content-length: 10
Connection: Close
status=mymessage;


head="POST http://twitter.com/statuses/update.json HTTP/1.1\n
Host: twitter.com\n
Authorization: Basic myname:passwordinbs64\n
Content-type: application/x-www-form-urlencoded\n
Content-length: 10\n
Connection: Close\n
status=mymessage";

# Https
[ref]
- [傳輸層安全-認識SSL](http://blog.yogo.tw/2009/11/ssl.html)
- [How Https works](http://robertheaton.com/2014/03/27/how-does-https-actually-work/)


### HTTPS Handshake
1. Hello , client say hello to server including the various cipher suites and maximum SSL version that it supports.
2. Certificate Exchange , server say hello to client with certificate, public key, cipher suite, client have, and SSL version.
3. Key Exchange , client generate a secret key and encrypt with server public key
4. Communicate with each other by secrete key

ps: 
- secret key is called session key , too.
- cipher suites = encryption algo. list


Trouble
- server private key is not private anymore , it will be eavesdropped by hacker , pretenting real server.
