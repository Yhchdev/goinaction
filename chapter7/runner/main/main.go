package main

import (
	"fmt"
	"log"
	"os"
	"task"
	"time"
)

var timeOut = time.Duration(4) * time.Second

func main() {
	r := runner.NewRunner(timeOut)

	r.Add(createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrInterrupt:
			log.Println(err)
			os.Exit(1)
		case runner.ErrTimeout:
			log.Println(err)
			os.Exit(2)
		}
		log.Println("Proccess done")
	}

	fmt.Println()
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Process - task: #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
