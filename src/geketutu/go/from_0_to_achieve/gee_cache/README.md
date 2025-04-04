### 分布式缓存


#### day1-lru

FIFO 先进先出

LFU 最少使用

LRU 最近最少使用

kv := ele.Value.(*entry) 断言，将 ele.Value 转换成 entry类型

#### day2-单机并发缓存

Sync.Mutex 互斥锁


Group 核心数据结构，负责与用户的交互，并且控制缓存值存储和获取的流程

day3 - HTTP 服务端
如何构建Http server

