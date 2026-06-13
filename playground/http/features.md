1. LocalStorage and SessionStorage
2. Web workers
3. Canvas and SVG
4. multi media tag 
5. HTML5 semantic elements
6. WebSQL , but this is not part of HTML5 spec


difference cookie , sessionStorage , localStorage

    Cookie在每個HTTP request送出時都會被送到Server端，不管你沒有要用到Cookie中的資訊，在某種程度上會拖慢執行的效能與浪費不必要的網路頻寬
    Cookie送出的資料本身並沒有加密，因此除非我們用SSL一類的技術做加密，否則Cookie中不宜放任何重要的資訊
    Cookies最大才4KB，不可能存太多資料

    HTML5的Storage主要分為兩種：localStorage與sessionStorage，
    這兩者主要在生命週期上有較明顯的差別，
    localStorage的生命週期較長，原則上要等到透過Javascript將內容清掉或者使用者清空Cache時才會消失；
    sessionStorage則是在Browser/Tab關閉時就會清空