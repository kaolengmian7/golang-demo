package sync_demo1

import (
	"log"
	"sync"
	"testing"
	"time"
)

/**

 */

var wait sync.WaitGroup
var counter int

func goroutine(id int) {
	for count := 0; count < 2; count++ {
		value := counter
		value++
		time.Sleep(time.Nanosecond)
		counter = value
	}

	wait.Done()
}

func TestRoutine(t *testing.T) {
	for routine := 1; routine <= 2; routine++ {
		wait.Add(1)
		go goroutine(routine)
	}

	wait.Wait()
	log.Println(counter)
}

//func Test(t *testing.T) {
//	for i := 0; i < 100; i++ {
//		counter = 0
//		TestRoutine(t)
//	}
//}
