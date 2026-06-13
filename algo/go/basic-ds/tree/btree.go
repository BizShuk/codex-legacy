package tree

// [Pattern]: [Balanced Binary tree] tree height difference between left/right sub-tree is less/equal than 1
// [Pattern]: [B+Tree] M-children balanced tree (MySQL)
//                     1. No data stored in nodes, only index
//                     2. Multiple key in a node
//                     3. Lower tree height => more index range content in disk. I/O optimization.
//                     4. Traverse only on leaf nodes.
// [Pattern]: [B-Tree] Similar to B+Tree, but nodes contain data

// [Note]: What does Mongo use B-Tree and MySQL use B+Tree
// B-Tree contain data in nodes. It means key is what we look for. Hash can be used to find the key as well. It increases the look up speed. In B+Tree, the key might have collison in the same leaf node. MySQL is relational DB. It needs to check all relationship between entity instead of just single data entity. B+Tree uses around logN time to find all related data entities. Hash cannot support range of data or fuzzy search, and lower performance if high hash key collision.
