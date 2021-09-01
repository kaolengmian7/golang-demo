package map_lab

import (
	"fmt"
	"testing"
)

// map 遍历测试，map输出是乱序的
func TestMapForRange(t *testing.T) {
	m := make(map[int64]string)
	m[1] = "a"
	m[2] = "b"
	m[3] = "c"
	m[4] = "d"
	m[5] = "e"
	for _, v := range m {
		fmt.Println(v)
	}
}

type User struct {
	Id int64
}

func TestMapToSlice(t *testing.T) {
	//创建了一个map value类型是字符串。string是基本类型，默认值是""。在访问不存在的key时，map[k]的输出结果是""
	m := make(map[int64]string)
	m[1] = "a"
	m[2] = "b"
	m[3] = "c"
	m[4] = "d"
	m[5] = "e"
	strSlice := make([]string, 0, 5)
	for i := int64(6); i < 10; i++ {
		strSlice = append(strSlice, m[i])
	}
	fmt.Println(strSlice)
	fmt.Println(len(strSlice))
	for _, str := range strSlice {
		fmt.Print(str)
	}
	//创建了一个map value类型是结构体。结构体默认值是nil。在访问不存在的key时，map[k]的输出结果是nil，如果此时再调用结构体的属性，就会报错nil pointer
	userMap := make(map[int64]*User)
	userSlice := make([]*User, 0, 5)
	for i := int64(0); i < 3; i++ {
		userSlice = append(userSlice, userMap[i])
	}
	fmt.Println(userSlice)
	fmt.Println(len(userSlice))
	for _, user := range userSlice {
		fmt.Print(user.Id)
	}
}
