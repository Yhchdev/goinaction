package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	maxgoroutines = 4
	taskLoad      = 10
)

var wg sync.WaitGroup

func init() {
	//rans := rand.Int(time.Now())
}

func main() {

	chs := make(chan string, taskLoad)
	wg.Add(maxgoroutines)

	for i := 0; i < maxgoroutines; i++ {
		go doWork(i, chs)
	}

	for task := 0; task < taskLoad; task++ {
		chs <- fmt.Sprintf("第%d项工作", task)
	}
	// 向有缓冲的chan中写完了数据，可以关闭了
	close(chs)
	wg.Wait()
}

func doWork(i int, chs chan string) {
	defer wg.Done()

	for {
		task, ok := <-chs
		if !ok {
			fmt.Println("处理完了所有的任务")
			return
		}
		fmt.Printf("goroutine:%d 开始做%s \n", i, task)
		time.Sleep(1 * time.Second)
	}
}
