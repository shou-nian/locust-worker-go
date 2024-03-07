package main

import (
	"github.com/myzhan/boomer"
	"time"
)

func main() {
	task := &boomer.Task{
		Name: "task1",
		Fn:   task1,
	}

	boomer.Run(task)
}

func task1() {
	start := time.Now()
	time.Sleep(time.Microsecond)
	elapsed := time.Since(start)

	// report test result as a success
	boomer.RecordSuccess("http", "task1", elapsed.Nanoseconds()/int64(time.Millisecond), int64(10))
}
