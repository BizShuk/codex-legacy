# NGINX configue


### http

### server 
- `listen <port>;`
- `server_name <domain> [<domain> ...];` , seperate by space
- `location` , check [location](#location)

### upstream
`upstream <upstream name>`

-`server <host>:<port> [weight=<num>]`


### location
`location [ |~|=|~=|...] <uri>`

- `proxy_pass http[s]://<upstream>/[uri];` , check [upstream](#upstream)
- 