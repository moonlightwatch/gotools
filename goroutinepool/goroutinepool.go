package goroutinepool

import (
	"errors"
	"fmt"
	"sync"
)

type GoroutinePool[T, R interface{}] struct {
	workerFunc   func(T) R
	t            int
	waitGroup    sync.WaitGroup
	waittingChan chan T
	resultChan   []chan Result[T, R]
}

type Result[T, R interface{}] struct {
	Result R
	Arg    T
	Err    error
}

// New 创建一个协程池，T 为参数类型，R 为返回值类型。参数 f 为执行函数。参数 t 为并发协程数。
// 协程池执行函数 f 的返回值类型必须为 R。
// 协程池执行函数 f 的参数类型必须为 T。
// 协程池等待队列的长度为 queueLen。
func New[T, R interface{}](f func(arg T) R, t int, queueLen int) *GoroutinePool[T, R] {
	p := &GoroutinePool[T, R]{
		workerFunc:   f,
		t:            t,
		waitGroup:    sync.WaitGroup{},
		waittingChan: make(chan T, queueLen),
		resultChan:   []chan Result[T, R]{},
	}
	if p.t < 1 {
		p.t = 1
	}
	p.start()
	return p
}

// Results 获取结果的channel
func (p *GoroutinePool[T, R]) Results() <-chan Result[T, R] {
	c := make(chan Result[T, R])
	p.resultChan = append(p.resultChan, c)
	return c
}

// Shutdown 关闭等待队列，并等待所有协程执行完毕
func (p *GoroutinePool[T, R]) Shutdown() {
	close(p.waittingChan)
	p.waitGroup.Wait() // 等待所有协程结束
	for _, c := range p.resultChan {
		close(c)
	}
}

func (p *GoroutinePool[T, R]) start() {
	for i := 0; i < p.t; i++ {
		p.waitGroup.Add(1)
		go p.runner() // 执行协程
	}
}

func (p *GoroutinePool[T, R]) runner() {
	defer p.waitGroup.Done() // 结束时计数
	// 从 waittingChan 获取参数，执行
	for args := range p.waittingChan {
		p.safeRun(args)
	}
}

func (p *GoroutinePool[T, R]) Submit(args T) {
	p.waittingChan <- args
}

// safeRun 拦截函数执行的错误
func (p *GoroutinePool[T, R]) safeRun(args T) {
	defer func() {
		if err := recover(); err != nil {
			// fmt.Printf("Run error: %+v\n", err)
			p.pushResult(Result[T, R]{
				Arg: args,
				Err: errors.New(fmt.Sprint(err)),
			})
		}
	}()
	result := p.workerFunc(args)

	go p.pushResult(Result[T, R]{
		Arg:    args,
		Result: result,
	})
}

func (p *GoroutinePool[T, R]) pushResult(r Result[T, R]) {
	for _, ch := range p.resultChan {
		ch <- r
	}
}
