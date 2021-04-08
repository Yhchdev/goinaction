package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)

	go Runner(baton)
	// 开始比赛
	baton <- 1

	wg.Add(1)
	wg.Wait()

}

func Runner(baton chan int) {
	//新的接力者
	var newRunner int

	// 接到接力棒
	runner := <-baton

	fmt.Printf("第%d位接力者开始起跑了\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		go Runner(baton)
	}

	//跑步中
	time.Sleep(2 * time.Second)
	fmt.Printf("%d在跑步中。。。", runner)
	time.Sleep(2 * time.Second)

	//第四位选手已经跑完了比赛
	if runner == 4 {
		fmt.Printf("%d跑完了，接力赛结束\n", runner)
		wg.Done()
		return
	}

	fmt.Printf("%d将接力棒交给了%d\n\n", runner, newRunner)

	baton <- newRunner
}
