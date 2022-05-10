package main

import (
	"sync"
)

// 多个goroutine同时访问一个变量时,需要对资源加锁或者其他操作 保证不会相互冲突或影响

// 同步锁 Mutex RWMutex
// sync.Mutex 最简单基础的同步锁 当goroutine持有锁，其他goroutine得等锁释放之后才可以尝试持有
// RWMutex 是读写锁 支持一写多读(允许支持多个goroutine同时持有读锁,只允许一个goroutine持有写锁)
// 当有goroutine持有读锁的时候，会阻止写操作。有goroutine持有写锁的时候，无论读写都会被堵塞

// 如果是Read多过于Write的场景,可以使用RWMutex
// 如果Write为主，那是用哪个都差不多

// 当前有多个goroutine, 但是只希望持有锁的goroutine执行
func MutexLock() {
	var lock sync.Mutex
	// 虽然for 启动了 10个goroutine 但是由于互斥锁存在 同一时刻只能有一个goroutine在执行
	for i := 0; i < 10; i++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			// do something
		}()
	}
}

// RWMutex区分了读写锁 有4个api 写锁：Lock、Unlock、 读锁：RLock、RLock 用法和Mutex一样

// 全局操作一次
// 有些场景 要求我们某段代码只能执行一次 例如单例模式
// 无论运行的过程中初始化多少次，得到的都是同一个实例 目的是减去创建实例的世界 尤其是数据库连接 hbase连接
// sync库提供了once 他可以传入一个函数,只允许全局执行这个函数一次
// 执行结束之前,其他goroutine执行到once语句的指挥会被阻塞
// 保证只有一个goroutine在执行once.当once执行结束之后,再次执行到这里，once语句内容会被跳过

// once e.g.
type Test struct{}

var test Test

var flag = false

func create() {
	test = Test{}
}

func inits() Test {
	// if !flag {
	// 	test = Test{}
	// 	flag = true
	// 	fmt.Println("done.")
	// }
	var once sync.Once
	once.Do(create)
	return test
}

//waitGroup
// 在主程序不知道goroutine执行结束的时间,如只是要依赖goroutine执行的结果 当然可以通过channel实现
// 希望等到goroutine执行结束后再执行后面的逻辑 用waitGroup

// waitgroup用法 有三个： Add、Done、Wait
// waitgroup内部 存储了当前有多少个goroutine在执行 当调用Add x时,表示产生x个新的goroutine
// 当这些goroutine执行完了, 让它调用 Done,表示执行结束了一个goroutine
// 当所有goroutine都执行完Done之后, wait的阻塞就会结束

// e.g WaitGroup
func waitGroupexam() {

	sample := Sample{}
	wg := sync.WaitGroup{}

	go func() {
		// 添加一个正在执行的goroutine
		wg.Add(1)
		// 执行完成后Done一下
		defer wg.Done()
		sample.JoinUserFeature()
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		sample.JoinItemFeature()
	}()

	wg.Wait()
	// do something
}

func main() {
	inits()
}
