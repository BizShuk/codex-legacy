virtual host

<VirtualHost 80:*>
	ServerName shuk.info			# 如果有兩個virtualhost 會依照進來的domain 執行下面的功能,都不符合 則第一個優先
	DocumentRoot /home/shuk/web		# 網頁根目錄
</VirtualHost>
