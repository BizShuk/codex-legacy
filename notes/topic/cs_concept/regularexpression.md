# Regular Expression

### Common Used

##### find same list number and put it in next line

find: ^(\d)(.*)\n([.\n\w\W]*)(\1.*)
replace: $1$2\n> $4\n$3

##### unique with sorted list

find: ^(.*)(\n\1)+
replace: $1

### complex regex

`[^/]`, anything not `/`
`.*?` or `.+?`, shortest match
`(?:abc)`, abc is optional
`.[^.]*`, `.*` 但不含 `..*`

### symbol

- ^, starting point
- $, ending point
- ., any char
- +, match previous char 1 or more time
- ?, match previous char 0 or 1 time
- *, match previous char 0 or more time
- ?:, optional
- \b, 字元邊界, 空白,標點,string開始/結尾

##### Reference

- [regex.com](http://www.regexr.com/)
- [regex info](http://www.regular-expressions.info/lookaround.html)
- [regex](http://www.dotblogs.com.tw/johnny/archive/2010/01/25/13301.aspx)
- [regex tester](https://regex101.com/#pcre)
