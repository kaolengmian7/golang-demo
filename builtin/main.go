package main

import "fmt"

func main() {
	testNilSlice()
}

type user struct {
	id int64
}

func testNilSlice() {
	int64Slice := make([]int64, 10)
	userSlice := make([]*user, 10)
	for _, v := range int64Slice {
		fmt.Printf("%d", v)
	}
	for _, v := range userSlice {
		fmt.Printf("%+v", *v)
	}
}
