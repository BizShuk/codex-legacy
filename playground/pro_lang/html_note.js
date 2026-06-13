HTML

Iframe
	$('#i1').load(function(){
	  	alert(3);
	});
	 $('#i1').ready(function(){
   		alert(4);
	});
	// output : 4 3 , iframe 先readey 才load

	alignment center: need display block or text-align:center;

z-index
	排列越後面的z-index越高
    position static會被其他的position蓋掉
    relative會跟static一起排


Difference GET, POST
	GET
		附加在URL後, 長度連同 URI 共 255 字元(1024 bytes)
		可以不必利用表單即可傳送參數
		Get 的 response 會被 cache 下來
		Get 的執行速度比較快，可加入書簽 (Bookmark) 中

	POST
		放在 HTTP request 的 message body 內傳送遞
		使用 Post 時，傳入的字串長度很大。
		Post 傳輸的數據量可以達到 2M。
		Post 的 response 不會被 cache 下來。
		Post 的執行速度比較慢，不可加入書簽 (Bookmark) 中