# sysctl


### vm

##### vm.swappiness
value:`0 ~ 100` (positive on using swap) , recommend it `0 ~ 10`
`0` will cause OOM kill.
<br>


ref:
- [config Swappiness](https://mariadb.com/kb/en/mariadb/configuring-swappiness/)

##### vm.dirty_background_ratio 
value:`#` , recommend it `5 ~ 10`

##### vm.dirty_ratio 
value:`#` , recommend it twice of `vm.dirty_background_ratio`

# Header Text #

### net

#####  net.ipv4.tcp_tw_recycle
set 1 , for to reduce time_wait

##### net.upv4.tcp_tw.reuse
set 1 , for to reduce time_wait
