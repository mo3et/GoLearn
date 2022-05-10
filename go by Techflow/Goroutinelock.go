package main

import "sync"

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

func main() {

}
