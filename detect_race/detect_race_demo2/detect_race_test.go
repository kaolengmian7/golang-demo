package detect_race_demo2

import (
	"log"
	"testing"
)

/**
对User接口的并发赋值
可能出现类似：Type指向了Jerry类型，而Data指向了Tom这种错误

检测data race: GORACE="halt_on_error=1" go test -race detect_race_test.go
*/

type User interface {
	Hello()
}

type Tom struct {
	name string
}

func (s *Tom) Hello() {
	//if s.name == "Jerry" {
	//	panic("Tom is not Jerry")
	//}
	log.Printf("Tom say: I'm %s", s.name)
}

type Jerry struct {
	name string
}

func (s *Jerry) Hello() {
	//if s.name == "Tom" {
	//	panic("Jerry is not Tom")
	//}
	log.Printf("Jerry say: I'm %s", s.name)
}

func Test(t *testing.T) {
	tom := &Tom{"Tom"}
	jerry := &Jerry{name: "Jerry"}

	var u User = tom

	var loopA, loopB func()
	loopA = func() {
		u = tom
		go loopB()
	}

	loopB = func() {
		u = jerry
		go loopA()
	}

	go loopA()

	for {
		u.Hello()
	}
}
