# Sorting

### ref.


### distinguish sort
- internal sort
- external sort

### stable
- stable , same key with same position
- unstable , same key with different position

### level
- 簡單排序
- 高等排序
    
### sorting list
sorting name , worst case avg case , stable , additional space , note
- bubble sort , O(n2) , O(n2) , stable , O(1) , better smaller n
- selection sort ,  O(n2) , O(n2) , unstable , O(1) , better smaller n , part-sorted better
- insertion sort , O(n2) , O(n2) , stable , O(1) , better smaller n , most-sorted better
- quick sort , O(n2) , O(nlog2 n) , unstable , O(n)~O(log n) , sorted is worst case
- heap sort , O(nlog2 n) , O(nlog2 n) , unstable , O(1) ,
- shell sort , O(n)~O(n2) , O(n(log2 n)2) , stable , O(1) , better smaller n
- merge sort , O(nlog2 n) , O(nlog2 n) , stable , O(n) , external sorting
- radix sort , O(nlogb B) , O(nlog2 n) , stable , O(n) , k:box count b:base
- counting sort , O(n2) , O(n2) , stable , O(1) , better smaller n
- bucket sort , O(n2) , O(n2) , stable , O(1) , better smaller n
