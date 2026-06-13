# Browser Hack #


| Inline CSS hack \ Browser		      | IE6	| IE7 | IE8 | Firefox | Chrome | Safari |
|-------------------------------------|-----|-----|-----|---------|--------|--------|
| selector { *Attr }				  | YES	| YES |     |         |        |        |
| selector { _Attr }				  | YES |     |     |         |        |        | 
| | selector { Attr\9 }				  | YES	| YES |	YES |         |        |        |
| nth-of-type(1) selector { Attr }	  | YES	| YES |		|    YES  |        |        |



-webkit for 
-ms
-moz

There are a few issues with older versions of Internet Explorer (8 and older). IE 8 doesn't recognize border-box on elements with min/max-width or min/max-height (this used to affect Firefox too, but it was fixed in 2012). IE 7 and below do not recognize box-sizing at all, but there's a polyfill that can help.