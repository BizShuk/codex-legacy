# Algorithm

### BUD Optimizatoin

- B = Bottleneck
- U = Unnecessary work
- D = duplicated work

### list

- BFS breadth-first search, push to queue and operate when dequeu
- DFS Depth-first search, push to stack and operate when pop
- Greedy
- backtracking
- sorting ,
  - [Bubble sort](algo/sort/bubble_sort.go) ,
  - [Insertion sort](algo/sort/insertion_sort.go) ,
  - [Quick sort](algo/sort/quick_sort.go) ,
  - [Selection sort](algo/sort/selection_sort.go) ,
  - [Sort2Dcolumn](algo/sort/sort2Dcolumn.go)
- Crypto
  - [SummaryRange](compress/summaryRange.go)
- Dynamic Programming
    Kadane's Algorithm
  - [MaxProfit multi transactions](dy/maxProfit_multi_transactions.go)
  - [MaxProfit one transaction](dy/maxProfit_one_transaction.go)
- Backtracking
  - Combine elements , [bf](backtracking/combination_bf.go) , [recur](backtracking/combination_recur.go) ,  goroutine of recursive
  - [Is valid parenthese](backtracking/valid_parenthese.go)
  - [parenthese combine](backtracking/parenthese_combine.go)

### DS

data structure

- [x] [Linklist](linklist/) , single , double , circle ,
- [ ] [Tree]
  - btree
  - multiple tree
  - red-black tree
  - full binary tree
  - complete binary tree
- [ ] [Stack]
- [ ] [Queue]
  - double-ended queue
  - priority queue
- [ ] [heap]  
- [ ] [map]
- [ ] [set]
- [ ] [hash]

skills:

- resource management
- operation function

### Reference

- [Top 7 Algo.](https://codingsec.net/2016/03/7-algorithms-data-structures-every-programmer/)
- [Top 10 Algo](http://www.techbang.com/posts/18438-ruled-the-worlds-top-ten-algorithms)
- [演算法筆記](http://www.csie.ntnu.edu.tw/~u91029/index.html)

### Algorithm list

- A\* search
- short url
- right-sided view b-tree
- Iru algo
- promise
- parser
- diff , lonest common pattern
- <http://www.csie.ntnu.edu.tw/~u91029/index.html>
- [LZ77 and LZ78](https://en.wikipedia.org/wiki/LZ77_and_LZ78#LZ77)
- [DEFLATE](https://en.wikipedia.org/wiki/DEFLATE)
- [OS scheduling algo.](http://www.tutorialspoint.com/operating_system/os_process_scheduling_algorithms.htm)
- [Scheduling algorithms](https://en.wikipedia.org/wiki/Category:Scheduling_algorithms) , include disk , network , and processor
- [Scheduling disciplines](https://en.wikipedia.org/wiki/Scheduling_(computing))
- [Huffman Code](huffman_code)
- base64
- [八皇后問題](https://zh.wikipedia.org/wiki/%E5%85%AB%E7%9A%87%E5%90%8E%E9%97%AE%E9%A2%98)
- 貪婪算法
- 啟發式算法
- 动态编程
- 原地算法』(in-place algorithm)
- 『外部排序』(external sorting)。
- 最近邻居算法』
- 数据库中还使用了其它启发式算法，像『模拟退火算法（Simulated Annealing）』、『交互式改良算法（Iterative Improvement）』、『双阶段优化算法（Two-Phase Optimization）』
- join algo
  - hash join
  - merge join
- buffer algo
  - 2Q（类LRU-K算法）
  - CLOCK（类LRU-K算法）
  - MRU（最新使用的算法，用LRU同样的逻辑但不同的规则）
  - LRFU（Least Recently and Frequently Used，最近最少使用最近最不常用）
  - LRU代表最近最少使用（Least Recently Used)
    一個buffer , 類似推文至底
  - LRU-K
  - 推测预读法（比如：如果查询执行器想要数据1、3、5，它不久后很可能会要 7、9、11），或者顺序预读法（这时候缓存管理器只是读取一批数据后简单地从磁盘加载下一批连续数据）。
  - 缓冲/缓存命中率
- encoding algo.
  - char range
  - could encode and decode?
  - security
  - sha1
  - base64
  - md5
