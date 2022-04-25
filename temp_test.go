package GOLang_Lab

import (
	"fmt"
	"testing"
)

const (
	mutexLocked           = 1 << iota // 是否上锁
	mutexWoken                        // 是否唤醒
	mutexStarving                     // 是否饥饿
	mutexWaiterShift      = iota      // mutexWaiterShift值为3，通过右移3位的位运算，可计算waiter个数
	starvationThresholdNs = 1e6       // 1ms，进入饥饿状态的等待时间
)

func TestUpdateMap(t *testing.T) {
	m := map[int]*Person{
		1: &Person{"Andy", 10},
		2: &Person{"Tiny", 20},
		3: &Person{"Jack", 30},
	}
	fmt.Println(m[1])

	p := m[1]
	p.name = "KL"
	fmt.Println(m[1])

}

type Person struct {
	name string
	age  int
}
