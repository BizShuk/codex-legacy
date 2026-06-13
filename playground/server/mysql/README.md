# MYSQL

### TODO
1. how to set in-memory mode
2. keep connection
3. send log by socket to log service
4. 


### db engine difference
Between 
- MyISAM
- HEAP
- Archive
- InnoDB
- NDB

check [here](http://ssorc.tw/663)

下述儲存引擎是最常用的：
- MyISAM：預設的MySQL插件式儲存引擎，它是在Web、數據倉儲和其他應用環境下最常使用的儲存引擎之一。注意，通過更改STORAGE_ENGINE配置變數，能夠方便地更改MySQL伺服器的預設儲存引擎。
- InnoDB：用於事務處理應用程式，具有眾多特性，包括ACID事務支援。
- BDB：可替代InnoDB的事務引擎，支援COMMIT、ROLLBACK和其他事務特性。
- Memory：將所有數據保存在RAM中，在需要快速搜尋引用和其他類似數據的環境下，可提供極快的訪問。
- Merge：允許MySQL DBA或開發人員將一系列等同的MyISAM資料表以邏輯方式組合在一起，並作為1個對象引用它們。對於諸如數據倉儲等VLDB環境十分適合。
- Archive：為大量很少引用的歷史、歸檔、或安全審計訊息的儲存和檢索提供了完美的解決方案。
- Federated：能夠將多個分離的MySQL伺服器連結起來，從多個物理伺服器建立一個邏輯資料庫。十分適合於分佈式環境或數據集市環境。
- Cluster/NDB：MySQL的叢集式資料庫引擎，尤其適合於具有高性能搜尋要求的應用程式，這類搜尋需求還要求具有最高的正常工作時間和可用性。
- Other：其他儲存引擎包括CSV（引用由逗號隔開的用作資料庫資料表的檔案），Blackhole（用於臨時禁止對資料庫的應用程式輸入），以及Example引擎（可為快速建立定制的插件式儲存引擎提供幫助）。
請記住，對於整個伺服器或方案，您並不一定要使用相同的儲存引擎，您可以為方案中的每個資料表使用不同的儲存引擎，這點很重要。


### lock
- table lock , MyISAM default
- row lock , InnoDB default , required clear pkey

details config and how they work?

### configure file
mysqlId.cnf , default in /etc/mysql/mysql.conf.d/



### note
- [4 performance of mysql](http://openlife.cc/blogs/2011/may/4-performance-fixes-mysql-large-servers) 
    + linux bug 
    + innodb contention on 100+GB buffer
    + vm.swappiness
    + 1 CPU socket