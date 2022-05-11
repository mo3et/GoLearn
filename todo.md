# 需要加深理解和解决的点
- interface
- struct的定义
- slice 二维slice
- map 
- 匿名函数
- select的使用
- []byte
- 网络I/O 调用net包下的网络I/O操作是不会阻塞线程的。当发起网络I/O请求时，go运行时会通过操作系统提供的epoll机制注册I/O事件，不会挂起实际干活的线程，只会切换goroutine而已。
- WithGroup


# 知识点
- slice 是不限制数组长度的array  s:=make([]string,3)
- slice[low:high]low为起始位置 high为slice容量
- Channel是在goroutine之间传输数据 ch:=make(chan DataType) 例如int interface{}


# 要求点
- [面试题](https://www.topgoer.cn/docs/gomianshiti/mianshiti)
- goroutine interface 等 Golang特性
- 掌握基本的数据结构和算法
- 熟悉缓冲技术，Redis 消息队列
- 熟悉主流数据库 MySQL Postgresql
- 熟悉TCP/IP协议、 http协议 Websocket编程！！！
- Linux环境操作 监控等命令

-[常见面试题](https://zhuanlan.zhihu.com/p/471490292)

# Source

[Go Web Example](https://gowebexamples.com/)


# 代码复习 面试题
https://www.topgoer.cn/docs/gomianshiti/mian2


https://learnku.com/articles/56078
[Go 基础语法面试题 ](https://geektutu.com/post/qa-golang-1.html#)

[Go语言面试题 ！！！](https://geektutu.com/series/#Go%20%E8%AF%AD%E8%A8%80%E9%9D%A2%E8%AF%95%E9%A2%98)

[Context](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/)