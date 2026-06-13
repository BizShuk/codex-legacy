# regex


### smallest possible match , non-greedy match
`.*?` or `.+?`


{} , repeat times



### example
- remove html tags , `<(.*?)>`


Regex.md 

- [regex.com](http://www.regexr.com/)
- [regex info](http://www.regular-expressions.info/lookaround.html)
- [regex](http://www.dotblogs.com.tw/johnny/archive/2010/01/25/13301.aspx)
- [regex tester](https://regex101.com/#pcre)

\b 字元邊界
    空白,標點,string開始/結尾



for vim

:/^create.*[^;]$/
=> create xxxxxxxxx  結尾不為 ;

.[^.]*
=> .* 但不含 ..*
~                     
