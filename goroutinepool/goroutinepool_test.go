package goroutinepool

import (
	"testing"
)

// TestGoroutinePool 测试GoroutinePool的正常行为和边界条件
func TestGoroutinePool(t *testing.T) {
	// 测试参数类型为int，返回值类型为int的函数
	testFunc := func(i int) int {
		return i * 2
	}

	// 创建一个协程池，设置并发协程数为2，等待队列长度为1
	pool := New(testFunc, 2, 1)
	defer pool.Shutdown()

	// 获取结果的channel
	resultChan := pool.Results()

	// 提交任务
	for i := 0; i < 5; i++ {
		pool.Submit(i)
	}

	// 检查结果是否正确
	for i := 0; i < 5; i++ {
		res := <-resultChan
		if res.Err != nil {
			t.Errorf("Unexpected error for input %d: %v", res.Arg, res.Err)
		} else if res.Result != testFunc(res.Arg) {
			t.Errorf("Expected %d, got %v for input %d", testFunc(res.Arg), res.Result, res.Arg)
		}
	}
}
