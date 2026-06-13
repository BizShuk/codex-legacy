package hash

// [Pattern]: [Hash Collision] How to deal with hash collistion
// [Solution]: Chaining
// [Solution]: Open Address (Probing)
//      probing algo affects insert,delete, seach time complexity
//      [algo]: linear probing
//      [algo]: quadratic probing
//      [algo]: double hashing probing => reduce clustering
// 不過Open Addressing使用Array存放資料，不需要頻繁使用動態記憶體配置(new/delete/malloc/free)，所以如果load factor沒有超過0.5(有些使用0.7)，那麼Open Addressing會是不錯的選擇。
