package main

import (
	"fmt"
)

// 多进程 Great 独立  Bad 但开销大
/* 多线程 Great 单个进程里面多个线程 创建、销毁成本低 线程之间通信及共享内存方便 开销小
Bad   开销 资源共享，由于共享更频繁，需要加上锁，避免数据紊乱 和 避免死锁 */
// 协程(Goroutine) 轻量级线程 本质还是线程 特点：协程的调度不是基于操作系统 而是基于程序的
// Goroutine 更像是 程序里的函数 但是在执行的过程中可以随时挂起、随时继续。

// def A():
// 	print '1'
// 	print '2'
// 	print '3'1

// def B():
// 	print 'x'
// 	print 'y'
// 	print 'z'

// TODO  协程 Goroutine
// 如果在一个线程内执行 A 和 B 两个函数 要么先A后B 要么先B后A 结果确定
// 但如果用 协程 来执行A和B 可能 A执行一半转而执行B B输出一条有回到A执行
// 不管执行的过程当中发生了几次 中断和继续,在操作系统当中执行的线程都没有发生变化 (是程序级的调度)

// 共享内存和CSP(通信顺序进程) Communicating Sequential Processes
// 共享内存 问题: 资源抢占、一致性  需要:多线程锁、原子操作等限制来 保证程序执行结果的正确性  CSP 模型

// TODO Golang只使用了 CSP中关于Process/Channel的部分。
// Process映射Goroutine Channel映射Channel。 Goroutine即Golang中的协程
// Goroutine之间没有任何耦合,可以完全并发执行。 Channel用于给Goroutine传递信息,保持数据同步。 Goroutine和Channel之间依然存在耦合
// Goroutine 和 Channel 结构类似于 生产消费者模式,多个线程之间通过队列共享数据，从而保持线程之间独立。

// goroutine
func Add(x, y int) int {
	z := x + y
	fmt.Println(z)
	return z
}

// 启动goroutine 执行
// go Add (3,4)

// 使用goroutine 执行一个匿名函数

go func(x,y int) {
	fmt.Println(x+y)
}(3,4)

// 使用 go 时,不能获取返回值!!
// z:= go Add (3,4)  违法

func maingo1() {
	go Add(3, 4)
	go func(x, y int) {
		fmt.Println(x + y)
	}(3, 4)