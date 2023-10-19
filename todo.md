# TODO List

- Go 内置数据结构(slice map channel goroutine) [文章](https://zhuanlan.zhihu.com/p/341945051)
- 数据结构和算法
- 数据库
- Redis
- WebSocket 如何高阶使用(心跳长连接、、)
- Kafka

---

## Channel

[Channel 的读取与写入](https://developer.aliyun.com/article/790587)
[Channel 面试题和解决过程](https://juejin.cn/post/7221495028031602748)
[Go Channel 用法 特性 最佳实践](https://zhuanlan.zhihu.com/p/613771870)
[Channel 全面解析](https://zhuanlan.zhihu.com/p/395278270)

## Kafka

[Kafka Broker 的那些事](https://juejin.cn/post/7136170688603226126)
[Kratos 使用 Kafka 视频](https://www.bilibili.com/video/BV1X24y1W7qY/)

## Redis

Redis 大 Key
[Redis 大 Key 多 Key 拆分](https://cloud.tencent.com/developer/article/1454332)
[redis Key 的管理](https://blog.csdn.net/qq_41834798/article/details/106769383)
[大厂如何使用 Redis docs](https://bytedance.feishu.cn/file/TcbGb6isWoTKbLxCaqfc77Qqnee)

## 数据结构和算法

### 数据结构

- Hash table

### 算法

[LeetCode Interview 75](https://leetcode.cn/studyplan/coding-interviews/)
[Golang 链表逆序](https://zhuanlan.zhihu.com/p/452832827)

## 数据库(MySQL)

[全解 MySQL](https://juejin.cn/post/7173581171547176997#heading-1)

## 资源

[build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md)

[Go 语言设计与实现](https://draveness.me/golang/)
[Go 语言高性能编程](https://geektutu.com/post/high-performance-go.html)

[Go 标准库 by example](https://books.studygolang.com/The-Golang-Standard-Library-by-Example)
[Go 语言必知必会](https://golang.dbwu.tech/)
[Go 语言专栏 Go 语言的点点滴滴](https://juejin.cn/column/7042670439683850270)

---

# Breaking!

---

# [Maybe Deprecated!!]

# 需要加深理解和解决的点

- select
- interface
- struct 的定义
- slice 二维 slice
- map
- 匿名函数
- select 的使用
- []byte
- 网络 I/O 调用 net 包下的网络 I/O 操作是不会阻塞线程的。当发起网络 I/O 请求时，go 运行时会通过操作系统提供的 epoll 机制注册 I/O 事件，不会挂起实际干活的线程，只会切换 goroutine 而已。
- WithGroup

# 知识点

- slice 是不限制数组长度的 array s:= make([]string,3)
- slice[low:high] low 为起始位置 high 为 slice 容量
- Channel 是在 goroutine 之间传输数据 ch:=make(chan DataType) 例如 int interface{}

# 要求点

- [面试题](https://www.topgoer.cn/docs/gomianshiti/mianshiti)
- goroutine interface 等 Golang 特性
- 掌握基本的数据结构和算法
- 熟悉缓冲技术，Redis 消息队列
- 熟悉主流数据库 MySQL Postgresql
- 熟悉 TCP/IP 协议、 http 协议 Websocket 编程！！！
- Linux 环境操作 监控等命令

-[常见面试题](https://zhuanlan.zhihu.com/p/471490292)

# Source

[Go Web Example](https://gowebexamples.com/)

# 代码复习 面试题

https://www.topgoer.cn/docs/gomianshiti/mian2

https://learnku.com/articles/56078
[Go 基础语法面试题 ](https://geektutu.com/post/qa-golang-1.html#)

[Go 语言面试题 ！！！](https://geektutu.com/series/#Go%20%E8%AF%AD%E8%A8%80%E9%9D%A2%E8%AF%95%E9%A2%98)

[Context](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/)

---

# 计划

1. 先 TCP/IP HTTP Socket
2. Go Redis 了解怎么写 理解缓存原理
3. 复习 Websockets 原理
4. 看看基本数据结构和算法
5. 复习几个盲点 select、byte slice、interface、struct、Methods！！
6. 继续研究如何使用 Docker Gin
7. goroutine、channel mod
8. RESTful API （可与 websocket 等网络协议一起看）
