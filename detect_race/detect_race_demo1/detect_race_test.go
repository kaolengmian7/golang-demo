package detect_race_demo1

import (
	"log"
	"testing"
)

/**
对User接口的并发赋值
nil pointer panic
*/

type User interface {
	Hello()
}

type Tom struct {
	name string
}

func (s *Tom) Hello() {
	log.Printf("Tom say: I'm %s", s.name)
}

type Jerry struct {
	id   int
	name string
}

func (s *Jerry) Hello() {
	log.Printf("Jerry say: I'm %s", s.name)
}

func Test(t *testing.T) {
	tom := &Tom{"Tom"}
	jerry := &Jerry{id: 1, name: "Jerry"}

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
