package main

import (
	"fmt"
	"log"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"lisa",
	"jason",
}

type namePrinter struct {
	name string
}

// 接口的实现，参数必须一致
func (n *namePrinter) Task() {
	log.Println(n.name)
	time.Sleep(1 * time.Second)
}

func main() {
	/*
		p := work.New(2)

		var wg sync.WaitGroup
		wg.Add(100 * len(names))
		for i := 0; i < 100; i++ {
			for _, name := range names {
				np := namePrinter{name: name}
				go func() {
					p.Run(&np)
					wg.Done()
				}()
			}
		}

		wg.Wait()
		p.Shutdown()


	*/

	const checkMark = "\u2713"
	const ballotX = "\u2717"
	fmt.Println(string(checkMark))
	fmt.Println(string(ballotX))

}
