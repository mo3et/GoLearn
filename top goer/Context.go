package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// e.g.
////////////////////////////////////////////
// 初始的例子

func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	// 如何接收外部命令实现退出
	wg.Done()
}

func main() {
	wg.Add(1)
	go worker()
	// 如何实现结束子Goroutine
	wg.wait()
	fmt.Println("over")
}

// 全局变量方式

package main

import (
    "fmt"
    "sync"

    "time"
)

var wg sync.WaitGroup
var exit bool

// 全局变量方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易统一
// 2. 如果worker中再启动goroutine，就不太好控制了。

func worker() {
    for {
        fmt.Println("worker")
        time.Sleep(time.Second)
        if exit {
            break
        }
    }
    wg.Done()
}

func main() {
    wg.Add(1)
    go worker()
    time.Sleep(time.Second * 3) // sleep3秒以免程序过快退出
    exit = true                 // 修改全局变量实现子goroutine的退出
    wg.Wait()
    fmt.Println("over")
}

// Channel方式
package main

import (
    "fmt"
    "sync"

    "time"
)

var wg sync.WaitGroup

// 管道方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的channel

func worker(exitChan chan struct{}) {
LOOP:
    for {
        fmt.Println("worker")
        time.Sleep(time.Second)
        select {
        case <-exitChan: // 等待接收上级通知
            break LOOP
        default:
        }
    }
    wg.Done()
}

func main() {
    var exitChan = make(chan struct{})
    wg.Add(1)
    go worker(exitChan)
    time.Sleep(time.Second * 3) // sleep3秒以免程序过快退出
    exitChan <- struct{}{}      // 给子goroutine发送退出信号
    close(exitChan)
    wg.Wait()
    fmt.Println("over")
}

// 官方方案

package main

import (
    "context"//
    "fmt"
    "sync"

    "time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
    for {
        fmt.Println("worker")
        time.Sleep(time.Second)
        select {
        case <-ctx.Done(): // 等待上级通知
            break LOOP
        default:
        }
    }
    wg.Done()
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    wg.Add(1)
    go worker(ctx)
    time.Sleep(time.Second * 3)
    cancel() // 通知子goroutine结束
    wg.Wait()
    fmt.Println("over")
}

// 子goroutine 又开启另外一个goroutine时

package main

import (
    "context"
    "fmt"
    "sync"

    "time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
    go worker2(ctx)
LOOP:
    for {
        fmt.Println("worker")
        time.Sleep(time.Second)
        select {
        case <-ctx.Done(): // 等待上级通知
            break LOOP
        default:
        }
    }
    wg.Done()
}

// 与上面的新增部分
func worker2(ctx context.Context) {
LOOP:
    for {
        fmt.Println("worker2")
        time.Sleep(time.Second)
        select {
        case <-ctx.Done(): // 等待上级通知
            break LOOP
        default:
        }
    }
}
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    wg.Add(1)
    go worker(ctx)
    time.Sleep(time.Second * 3)
    cancel() // 通知子goroutine结束
    wg.Wait()
    fmt.Println("over")
}


///////////////////////////////////////////////

// Context 知识
// Context用来简化 对于处理单个请求的多个goroutine直接与请求域的数据、取消信号、截止时间等相关操作
// 对服务器传入的请求应该创建上下文Context  对服务器的传出调用应该接受上下文Context
// 他们直接的函数调用链比赛传递Context  or 使用WithCancel 、WithDeadline、WithTimeout 或WithValue创建的派生Context
// 当一个Context被取消时，他派生的所有Context也被取消

	// Context接口
// context.Context是一个接口 该接口定义了四个需要实现的方法。具体签名如下
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
// 其中：
// Deadline方法需要返回当前Context被取消的时间，也就是完成工作的截止时间（deadline）；
// Done方法需要返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done方法会返回同一个Channel；
// Err方法会返回当前Context结束的原因，它只会在Done返回的Channel被关闭时才会返回非空的值；
// 如果当前Context被取消就会返回Canceled错误；
// 如果当前Context超时就会返回DeadlineExceeded错误；
// Value方法会从Context中返回键对应的值，对于同一个上下文来说，多次调用Value 并传入相同的Key会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据；

// Background()和TODO()
// Go内置Background()和TODO() 分别返回一个实现了Context接口的background和todo 
// Background()主要用关于main函数、初始化及测试代码中，作为Context这个树状结构的最顶层Context，也就是Root Context
// Func TODO() 如果不知道什么时候使用Context，可以使用这个
// background和todo 本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context

// 四个With系列函数 WithCancel、WithDeadline、WithTimeOut、WithValue
	// WithCancel
// WitchCancel返回带有行Done通道的父节点的副本。当调用返回的cancel函数 or close context的Done 通道时
// 将关闭返回Context的Done通道，无论先发生什么情况
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

// 取消此Context将释放与其相关联的资源，因此代码应该再次Context中运行的操作完成后立刻调用cancel
func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return结束该goroutine，防止泄露
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
func main() {
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // 当我们取完需要的整数后调用cancel

for n := range gen(ctx) {
	fmt.Println(n)
	if n == 5 {
		break
	}
}
}
// gen函数在单独的goroutine中生成证书并将它们发送到返回的通道。
// gen的调用者在使用生成的证书之后需要Cancel Context 以免gen启动的内部goroutine发生泄漏

// WithDeadline
func WithDeadline(parent Context,deadline time.Time)(Context,CancelFunc)

// 返回parentContext的副本,并将deadline调整为不迟于d。
// 如果parent Context的deadline 已经早于d,这WithDeadline(parent,d)在语义上等同于Parent Context
// 当截止日过期时，当调用返回的cancel函数时，或者当父上下文的Done通道关闭时，返回上下文的Done通道将被关闭.

// 取消此Context将释放与其关联的资源,因此代码应该在此Context运行的操作完成后立刻调用Cancel

func main(){
	d:=time.Now().Add(50 * time.Millisecond)
	ctx,cancel:=context.WithDeadline(context.Background(),d)

	// 尽管ctx会过期，当在任何情况下调用他的cancel函数都是很好的实践
	// 如果不这样做，可能会使Context及其parent class存货的世界超过必要的时间

	// 定义了一个50ms过期的deadline 然后调用context.WithDeadline(context.Background(),d) 
	// 得到一个ctx和一个cencel，然后用一个select 让 func main陷入等待
	// 等待1秒后打印overslept退出 或者等待ctx过期后退出。 
	// 因为ctx 50s后就过期，所以ctx.Done()会想收到值，上面代码会Print ctx.Err()取消原因
	defer cancel()

	select{
	case <- time.After(1 * time.Second):
		fmt.Println("overslept")
	cast <- ctx.Done():
		fmt.Println(ctx.Err())
	}
}

// WithTimeout
func WithTImeout(parent Context,Timeout time.Duration)(context,CancelFunc)

// 取消此Context将释放与其相关的资源，因此代码应该在此Context中运行的操作完成后立刻调用cencel
// 通常用于数据库或者网络连接的超时控制

package main

import(
	"context"
	"fmt"
	"sync"

	"time"
)

// context.WithTimeout

var wg sync.WaitGroup

func worker(ctx context.Context){
	LOOP:
	for{
		fmt.Println("db connecting ...")
		time.Sleep(time.Millisecond * 10)//架设正常连接数据库耗时10ms
		select{
		case <-ctx.Done(): //50ms 后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main(){
	// 设置一个50ms的timeout
	ctx,cancel:=context.WithTimeout(context.Background(),time.Millisecond*50)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second *5)
	cancel() //通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}

// WithValue
func WithValue(parent Context,key val interface{}) Context
// WithValue返回parent Context的副本，其中与key关联的值为val

// 仅对API和进程间传递请求域的数据使用Context 值,而不是使用它来传递可选参数给函数 
// 所提供的键必须是可比较的，并且不应该是string类型或任何内置类型，避免使用Context在包之间发生冲突。
// WithValue的用户应该为 key 定义之间的类型。为了避免在分配给interface{}时进行分配，Context key通常有具体类型struct{}
// 或者 导出的Context 关键变量的静态类型应该是指针或接口

package main

import (
    "context"
    "fmt"
    "sync"

    "time"
)

// context.WithValue

type TraceCode string

var wg sync.WaitGroup

func worker(ctx context.Context) {
    key := TraceCode("TRACE_CODE")
    traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
    if !ok {
        fmt.Println("invalid trace code")
    }
LOOP:
    for {
        fmt.Printf("worker, trace code:%s\n", traceCode)
        time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
        select {
        case <-ctx.Done(): // 50毫秒后自动调用
            break LOOP
        default:
        }
    }
    fmt.Println("worker done!")
    wg.Done()
}

func main() {
    // 设置一个50毫秒的超时
    ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
    // 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
    ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
    wg.Add(1)
    go worker(ctx)
    time.Sleep(time.Second * 5)
    cancel() // 通知子goroutine结束
    wg.Wait()
    fmt.Println("over")
}

// 使用Context 注意事项
// 1.推荐以参数的方式显示传递Context
// 2.以Context作为参数的函数方法，应该把Context作为第一个参数。
// 3.给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO()
// 4.Context的Value相关方法应该传递请求域的必要数据，不应该用于传递可选参数
// 5.Context是线程安全的，可以放心的在多个goroutine中传递