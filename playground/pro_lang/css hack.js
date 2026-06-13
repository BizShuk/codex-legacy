css hack

Inline CSS hack \ Browser		IE6	IE7	IE8	Firefox	Chrome	Safari
selector { *Attr }				YES	YES	 	 	 	 
selector { _Attr }				YES	 	 	 	 	 
selector { Attr\9 }				YES	YES	YES	 	 	 
nth-of-type(1) selector { Attr }	 	 	YES		YES		YES
@media screen and (-webkit-min-device-pixel-ratio:0){ selector { Attr } }	 	 	 	 	
													YES		YES




arrow by border
http://fundesigner.net/css-triangle/

scrollbar hidden
外面hidden 裡面scroll 但把scrollbar移到外面container看不見的地方


*****
<link rel="dns-prefech" href="<dns>" />
 使用工具 page insight  來觀測網頁情況






##### #圖片邊界 顏色變換
box-border 讓 border 在 container 內 , 並用 rgba(0,0,0,0.1) 讓0.1透明度的黑色邊界 蓋在原圖上面 => 0.1透明之黑色邊界+原圖邊界顏色