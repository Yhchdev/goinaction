package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	// 伪随机数设置种子，保证每次生成的随机数都是不一样的
	rand.Seed(time.Now().UnixNano())
}

// learn无缓冲区buffer，类比打网球，球员一接一收
func main() {
	count := make(chan int)

	wg.Add(2)

	go play("bob", count)
	go play("lisa", count)

	// 开始发球
	count <- 1

	wg.Wait()
}

func play(name string, count chan int) {
	defer wg.Done()

	for {
		// 接到球
		ball, ok := <-count
		if !ok {
			// 通道被关闭
			fmt.Printf("恭喜%s,你赢了\n", name)
			return
		}

		// 模拟丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			close(count)
			fmt.Printf("%s丢掉了第%d个球\n", name, ball)
			return
		}

		fmt.Printf("%s 打出第 %d 个球\n", name, ball)
		ball++

		// 将球打出去
		count <- ball
	}
}
