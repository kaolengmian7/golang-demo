package main

import (
	"fmt"
	"math/rand"
	"time"

	"strconv"
)

func main() {
	result := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		str := strconv.Itoa(rand.Intn(1000000))
		result += str + ","
	}
	for i := 0; i < 5; i++ {
		str := strconv.Itoa(rand.Intn(100))
		result += str + ","
	}
	for i := 0; i < 5; i++ {
		str := strconv.Itoa(rand.Intn(1000))
		result += str + ","
	}
	for i := 0; i < 5; i++ {
		str := strconv.Itoa(rand.Intn(10000))
		result += str + ","
	}
	fmt.Println(result)
}
