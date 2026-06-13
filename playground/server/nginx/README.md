# NGINX


# NGINX

nginx_note

Nginx conf
- [special variable](http://nginx.org/en/docs/varindex.html)
- [NGINX-LOCATION設定範例](http://homeserver.com.tw/2010/05/21/nginx-location%E8%A8%AD%E5%AE%9A%E7%AF%84%E4%BE%8B/)
- [nginx配置location总结及rewrite规则写法](http://seanlook.com/2015/05/17/nginx-location-rewrite/)
- [Nginx source learning resource](http://www.cnblogs.com/yjf512/archive/2012/06/13/2548515.html)

### nginx default from apt-get
/etc/nginx/
/usr/share/nginx/html
/var/log/nginx




### build from soucre
--with_http_stub_status_module , for 線上人數

    location /nginx_status {
      stub_status on;
      access_log off;
      allow all; # 允許看到的 IP
    }







### # problem
1.  ipv6 nameserver in /etc/resolv.conf 造成 nginx 無法解析 => 移除ipv6的 nameserver
    在/etc/network/interface 需要的網卡加上 dns-nameservers 8.8.8.8 8.8.4.4




# cmd
- nginx -V , show version
- nginx -T , test configuration
- nginx -s [start,quit,reopen,reload] , signal
- nginx -c conf_file , set conf file

