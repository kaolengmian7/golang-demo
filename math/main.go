package main

import (
	"fmt"
	"math/rand"
	"time"
)

//随机打乱 slice 并输出
func main() {
	dest := MicSlice([]int64{1, 2, 3, 4, 5, 6, 7}, 2)
	fmt.Println("result....", dest)
}

func MicSlice(userIds []int64, count int) []int64 {
	temp := make([]int64, len(userIds))
	copy(temp, userIds)
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(temp), func(i, j int) {
		temp[i], temp[j] = temp[j], temp[i]
	})
	fmt.Println("temp....", temp)
	result := make([]int64, 0)
	for index, value := range temp {
		if index == count {
			break
		}
		result = append(result, value)
	}
	return result
}
