//@program: superlion
//@author: yanjl
//@create: 2023-09-12 10:09
package async

import "sync"

type Pool struct {
	capacity       uint64            // 最大协程数
	runningWorkers uint64            // 当前正在运行的协程数
	status         int64             // 协程池的状态
	chTask         chan *Task        // 执行任务的 channel
	PanicHandler   func(interface{}) // 处理协程中的 panic 异常
	sync.Once                        // 防止多次调用stop
	sync.Mutex                       // 用于锁定协程池的状态和 channel
}

type Task struct {
}

func NewPool(n uint64, panicHandler func(interface{})) *Pool {
	return &Pool{
		capacity:     n,
		status:       Running,
		chTask:       make(chan *Task, n),
		PanicHandler: panicHandler,
	}
}
