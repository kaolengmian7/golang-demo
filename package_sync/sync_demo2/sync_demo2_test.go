package sync_demo2

import (
	"log"
	"testing"
)

func TestStateIsStarving(t *testing.T) {
	mutexLocked := int32(1)   // 0001
	mutexStarving := int32(4) // 0100
	log.Printf("=== mutexLocked 按位与 mutexStarving 结果：%d ===\n", mutexLocked|mutexStarving)

	var old int32
	old = 1 // 0001 正常状态
	log.Println(old&(mutexLocked|mutexStarving) == mutexLocked)

	old = 5 // 0101 饥饿状态
	log.Println(old&(mutexLocked|mutexStarving) == mutexLocked)
}
