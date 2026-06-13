# HTTP Protocal

POST /luajit.lua HTTP/1.1
Host: 192.168.2.52
Content-type: text/plain
Content-length: 10
Connection: Close
status=mymessage;

head="POST <http://twitter.com/statuses/update.json> HTTP/1.1\n
Host: twitter.com\n
Authorization: Basic myname:passwordinbs64\n
Content-type: application/x-www-form-urlencoded\n
Content-length: 10\n
Connection: Close\n
status=mymessage";
