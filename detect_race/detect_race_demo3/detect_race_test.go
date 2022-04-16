package detect_race_demo3

import (
	"fmt"
	"os"
	"sync"
	"testing"
)

// 循环创建 goroutine 打印 i 变量，goroutine使用临时变量会产生 data race
func TestDataRaceForRangeAdd(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i) // Not the 'i' you are looking for.
			wg.Done()
		}()
		// 使用time.sleep 后，因为for循环变慢了，程序正常打印出0 1 2 3 4
		//time.Sleep(time.Second)
	}
	wg.Wait()
}

// 不小心在多个 goroutine 共享了 error 变量
func TestDataRaceError(t *testing.T) {
	data := []byte("data")
	res := make(chan error, 2)

	f1, err := os.Create("file1")
	go func() {
		// This err is shared with the main goroutine,
		// so the write races with the write below.
		_, err = f1.Write(data)
		res <- err
		f1.Close()
	}()

	f2, err := os.Create("file2") // The second conflicting write to err.
	go func() {
		_, err = f2.Write(data)
		res <- err
		f2.Close()
	}()

}
