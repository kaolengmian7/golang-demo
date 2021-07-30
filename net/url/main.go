package main

import (
	"fmt"
	"net/url"
)

func main() {
	url1 := "https://haolaoda.com/test/abc.html?id=123456"
	//url 转码函数
	//为什么转码函数起这个名字？
	//个人一点点见解（不一定对）：query查询，escape逃逸。逃逸查询也就意味着查不到，查不到=被转码
	escape := url.QueryEscape(url1)
	fmt.Printf("%s \n", escape)
	//QueryUnescape是有err返回值的。如果某个%后面没有紧跟着两个十六进制数字，就会报错
	unescape, err := url.QueryUnescape(escape)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	fmt.Printf("%s \n", unescape)
	//解析url
	parse, err := url.Parse(url1)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	fmt.Printf("%s \n", parse.Scheme)
	fmt.Printf("%s \n", parse.Opaque)
	fmt.Printf("%s \n", parse.Host)
	fmt.Printf("%s \n", parse.Path)
	fmt.Printf("%s \n", parse.RawQuery)
	fmt.Printf("%s \n", parse.RawPath)

	//url.parseQuery 函数会将一个url字符串的[?query]解析成map
	//如果业务代码在之前已经用过url.parse方法，那么建议直接使用 parse.Query()方法来获得[?query]map
	//parseQuery := parse.Query()
	parseQuery, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	//源码中，parse.RawQuery转换成的map是一个名叫Values 的结构体
	//下面是其增删改查操作
	fmt.Printf("%s \n", parseQuery.Get("id"))
	parseQuery.Set("id", "789789")
	fmt.Printf("%s \n", parseQuery.Get("id"))
	parseQuery.Add("passwd", "unnoo")
	fmt.Printf("%s \n", parseQuery.Get("passwd"))
	parseQuery.Del("id")
	fmt.Printf("%s \n", parseQuery.Encode())
	parse.RawQuery = parseQuery.Encode()
	fmt.Printf("%s \n", parse.String())

	//url.prase 函数能解析url中的用户名和密码，将其解析为Userinfo结构体。当然，已经很少有人在url中携带这种机密信息了
	urlWithUserInfo := "https://username:unnoo@haolaoda.com/test/abc.html?id=123456"
	parse, err = url.Parse(urlWithUserInfo)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	fmt.Printf("%v \n", parse.User)
}
