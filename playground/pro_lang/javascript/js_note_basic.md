reference:
- [js tips](http://ejohn.org/apps/learn/)
- [js using patterns](http://shichuan.github.io/javascript-patterns/#function-patterns)
- [极客标签](http://www.gbtags.com/)
- [码农周刊](http://weekly.manong.io/issues/)
- [前端周刊](http://www.feweekly.com/issues)
- [极客头条](http://geek.csdn.net/)
- [Startup News](http://news.dbanotes.net/)
- [Hacker News](https://news.ycombinator.com/news)
- [InfoQ](http://www.infoq.com/)
- [w3cplus](http://www.w3cplus.com/)
- [Stack Overflow](http://stackoverflow.com/)
- [Atp](http://atp-posts.b0.upaiyun.com/posts/)
- [web開發筆試面試](http://mianshiti.diandian.com/)



### # special note
- delegating(ie8以下不支援) and bubbling
    當window接觸到event 會從window往下傳(delegating) 到了目標後會先觸發bubble event 然後才是delegated event 在往上傳(bubbling)
    e.preventDefault() => stop default event
        e.retrunValue = false , for ie
    e.stopPropagation() => stop bubbling
        e.cancelBubble = true , for ie

    jQuery
        delegate
            $("ul").delegate("li", "click", function(){
                $this).hide();
            });
            // 將ul下面的li都加上click event 即使是新增的



- window.location相關的redirect
    如果用 /xxxx/xxx 跳轉domain下的連結 在port非80的狀況下會有問題


js延迟加载
    defer和async、动态创建DOM方式（用得最多）、按需异步载入js

cross domain
    可解決方案 jsonp、 iframe、window.name、window.postMessage、服务器上设置代理页面


- 硬記
    console.log(typeof null);           // object
    console.log(typeof {});             // object
    console.log(typeof []);             // object ,but [] instanceof Array
    console.log(typeof undefined);      // undefined
    console.log(typeof new Array());    // object
    console.log(typeof function(){});   // function
    console.log(typeof new Date())      // object



優化手段
    for (var i = size; i < arr.length; i++) {}
    for 循环每一次循环都查找了数组 (arr) 的.length 属性，在开始循环的时候设置一个变量来存储这个数字，可以让循环跑得更快：
    for (var i = size, length = arr.length; i < length; i++) {}

    yahoo 14條優化手段
        http://inspiregate.com/internet/seo/62-yahoo-14-rules-to-accelerate-access-to-the-site.html
    高性能網站建構指南
        http://div.io/topic/371

diff == vs ===
    "=="    =>  會自動作型別轉換, ex: false == "" == 0
    "==="   =>  絕對相等 type與value 皆須相同

    about object
        var a = {name: "hello"};
        var b = {name: "hello"};
        a === b; //false

    特殊狀況
        NaN   ==  NaN   =>  false   ,特殊狀況NaN不屬於任何值 沒有值=>永遠不會相等
        NaN   === NaN   =>  false   ,特殊狀況NaN不屬於任何值 沒有值=>永遠不會相等
        true  ==  '1'   =>  true
        true  ==  '5'   =>  false
        false ==  '0'   =>  true
        true  === '1'   =>  false
        true  === '5'   =>  false
        false === '0'   =>  false

difference cookie , sessionStorage , localStorage

    Cookie在每個HTTP request送出時都會被送到Server端，不管你沒有要用到Cookie中的資訊，在某種程度上會拖慢執行的效能與浪費不必要的網路頻寬
    Cookie送出的資料本身並沒有加密，因此除非我們用SSL一類的技術做加密，否則Cookie中不宜放任何重要的資訊
    Cookies最大才4KB，不可能存太多資料

    HTML5的Storage主要分為兩種：localStorage與sessionStorage，
    這兩者主要在生命週期上有較明顯的差別，
    localStorage的生命週期較長，原則上要等到透過Javascript將內容清掉或者使用者清空Cache時才會消失；
    sessionStorage則是在Browser/Tab關閉時就會清空

