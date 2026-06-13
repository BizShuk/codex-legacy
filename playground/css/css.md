# Cascading Style Sheet #


### selector ###
- `<tag>`
- `.<className>`
- `#<id>`
- `[attibute=""]` , `[attibute~=""]` , `[attibute|=""]` , `[attibute^=""]` , `[attibute$=""]` , `[attibute*=""]` 
    + `~` , The following example selects all elements with a title attribute that contains a space-separated list of words, one of which is "flower"
    + `|` , start with or attach subsequently with `-`
    + `^` , start with 
    + `$` , end with
    + `*` , contain 
- selectorA , selectorB , both A and B have attributes 
- selectorA   selectorB , any B in A(include children of children)
- selectorA > selectorB , any B after A(one level parent)
- selectorA + selectorB , any B after A(sibling) immediately
- selectorA ~ selectorB , any B after A(sibling)

### pseudo class ###
- empty
- a link
    + hover
    + link
    + visited
    + active
    + focus
    + target
- input
    + checked
    + disabled
    + invalid
    + read-only
    + read-write
    + required
    + optional
- lang(language)
- after
- before
- first-letter
- first-line
- first-child
- first-of-type
- last-child
- last-of-type
- nth-child(n)
- nth-last-child(n)
- nth-last-of-type(n)
- nth-of-type(n)
- only-of-type
- only-child

- in-range
- out-of-range


### diplay ###
- block
- inline-block , space between divs , because space between element
- flex
- static
- fixed
- absolute

### flex ###
- display: flex
- align-item: center , vertical align
- justify-content: , horizontal align


### float ###
- float:left right
- margin:0 auto;
http://blog.yam.com/hanasan/article/35806444


### overflow ###
- hidden
- scroll
- auto
外面hidden 裡面scroll 但把scrollbar移到外面container看不見的地方






### treaks ###

1. 圖片邊界 顏色變換
box-border 讓 border 在 container 內 , 並用 rgba(0,0,0,0.1) 讓0.1透明度的黑色邊界 蓋在原圖上面 => 0.1透明之黑色邊界+原圖邊界顏色

2. arrow by border
http://fundesigner.net/css-triangle/

3. div vertical and horizontal align
    1. 
        * margin: auto
        * top: 0; left: 0; bottom: 0; right: 0;
        * position: abosulte;
        * This requires the parent to be positioned (i.e. position: relative)
    2. 
        * vertical-align : center   , **Need a alignment by other element , like ::before**
        * line-height:50% for one line
        * add a div height:50% and margin-bottom: -100px;
    3. add floater on the top of element
