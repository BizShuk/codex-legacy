# Sudofile

## group

with `%` ahead, it mean group

## example

> Cmnd_Alias     CAPTURE = /usr/sbin/tcdump
> Cmnd_Alias     SERVERS = /usr/sbin apache2ctl, /usr/bin/htpasswd
> Cmnd_Alias     NETALL = CAPTURE, SERVERS
> %netadmin ALL=NETALL

group `netadmin` haver permission on `CAPTURE` and `SERVERS`
