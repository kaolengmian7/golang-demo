package main

import (
	"fmt"
	"sort"
)

func main() {
	//字符串排序
	//根据文档的描述
	//sort函数的参数需要 sort包下的Interface类型，sort.StringSlice实现了该接口，可以作为参数传入
	strList := make([]string, 5)
	strList[0] = "bbb"
	strList[1] = "aaa"
	strList[2] = "zzz"
	strList[3] = "ccc"
	strList[4] = "abv"
	strSlice := sort.StringSlice(strList)
	//根据输出可见，字符串排序是根据首字母（a-z）的顺序排序的,如果首字母相同，则对比第二个字母
	fmt.Printf("%v", strSlice.Search("bbb"))
	//strSlice.Sort()
	//sort.Sort(strSlice) 写法二
	fmt.Printf("%v", strList)
	//Search函数使用二分搜索法查找，返回下标
	//因为使用二分搜索法，Search方法一定是先排序再查找
	//所以一般使用Search找到值x在插入一个有序的、可索引的数据结构时，应插入的位置。
	fmt.Printf("%v", strSlice.Search("aaa"))

	//中文字符串排序
	//排序默认是按照Unicode码点的顺序的, 如果需要按照拼音排序, 可以通过GBK转换实现, 自定义一个排序接口
	chStrList := make([]string, 5)
	chStrList[0] = "啊啊啊"
	chStrList[1] = "不是"
	sort.Sort(sort.StringSlice(chStrList))
	fmt.Printf("%v", chStrList)

}
