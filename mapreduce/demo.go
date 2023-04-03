package main

import (
	"fmt"
	"sync"
)

/*
这是一个使用 Go 语言编写的 MapReduce 程序简单示例。它演示了如何使用 Go 语言中的并发特性来实现 MapReduce 模型。

这个程序用来统计文本中单词的个数：
	1.定义了一个 MapReduce 函数，它接受输入数据、map 函数和 reduce 函数作为参数。
	2.MapReduce 函数首先执行 map 阶段，将输入数据映射为中间数据。然后，它将中间数据按键分组，并发地执行 reduce 阶段，将每组数据合并为最终结果。

数据处理流程：
String -> Word List -> Word List Groups -> Result
*/

// KeyValue is a type that represents a key-value pair
type KeyValue struct {
	Key   string
	Value int
}

// MapFunc is a type that represents a map function
type MapFunc func(string) []KeyValue

// ReduceFunc is a type that represents a reduce function
type ReduceFunc func(string, []int) int

// MapReduce performs a map-reduce operation on the input data
func MapReduce(input []string, mapFunc MapFunc, reduceFunc ReduceFunc) map[string]int {
	// Map 阶段
	var intermediate []KeyValue
	for _, line := range input {
		intermediate = append(intermediate, mapFunc(line)...)
	}

	// 将数据按照 Key 分组
	groups := make(map[string][]int)
	for _, kv := range intermediate {
		groups[kv.Key] = append(groups[kv.Key], kv.Value)
	}

	// Reduce 阶段
	var wg sync.WaitGroup
	output := make(map[string]int)
	results := make(chan KeyValue)
	go func() {
		for kv := range results {
			output[kv.Key] = kv.Value
		}
	}()
	for key, values := range groups {
		wg.Add(1)
		go func(key string, values []int) {
			defer wg.Done()
			result := reduceFunc(key, values)
			results <- KeyValue{key, result}
		}(key, values)
	}
	wg.Wait()
	close(results)

	return output
}

func main() {
	// Define the input data
	input := []string{
		"The quick brown fox jumps over the lazy dog",
		"The quick brown dog jumps over the lazy fox",
	}

	// Define the map function
	mapFunc := func(line string) []KeyValue {
		var result []KeyValue
		for _, word := range words(line) {
			result = append(result, KeyValue{word, 1})
		}
		return result
	}

	// Define the reduce function
	reduceFunc := func(key string, values []int) int {
		return len(values)
	}

	// Perform the map-reduce operation
	output := MapReduce(input, mapFunc, reduceFunc)

	// Print the result
	for key, value := range output {
		fmt.Printf("%s: %d\n", key, value)
	}
}

// 分词器：将字符串转换为单词列表
func words(line string) []string {
	var result []string
	var word string
	for _, r := range line {
		if r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' {
			word += string(r)
		} else if word != "" {
			result = append(result, word)
			word = ""
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
