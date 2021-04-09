package main

import (
	"errors"
	"os"
	"time"
)

// 程序生命周期控制

// 程序在指定时间内完成一组超时任务，超时退出
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

// 收到操作系统停止任务是返回
var Interrupt = errors.New("received interrupt")

// 工厂函数
func NewRunner(duration time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(duration),
	}
}

// ...可变参数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func main() {

}
