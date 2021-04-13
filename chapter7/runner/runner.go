package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// 程序生命周期控制

// 指定事件内完成，正常结束
// 程序在指定时间内完成一组超时任务，超时退出,"自杀"
// 收到操作系统的的退出信号也会退出

type Runner struct {
	// 操作系统信号
	interrupt chan os.Signal

	// 正常运行信号量，可能会返回错误
	complete chan error

	//超时信号量
	timeout <-chan time.Time

	// 一组按序执行的任务函数
	tasks []func(int)
}

// 任务超时时会返回
var ErrTimeout = errors.New("received timeout")

// 收到操作系统停止任务时返回
var ErrInterrupt = errors.New("received interrupt")

// 工厂函数
func NewRunner(duration time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(duration), //到时间之后会自动收到一个时间的返回值
	}
}

// ...可变参数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start:执行任务并监视通道事件
func (r *Runner) Start() error {
	//接收系统的所有中断消息
	signal.Notify(r.interrupt, os.Interrupt)

	//用不同goroutine执行不同任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

// run：执行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.goInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

// goInterrupt:验证是否接收到了中断信号
func (r *Runner) goInterrupt() bool {
	select {
	//第一次中断事件触发
	case <-r.interrupt:
		//停止接收后续中断事件
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
