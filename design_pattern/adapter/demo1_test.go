package adapter

import (
	"testing"
)

// demo1：传统适配器模式，所有流行的编程语言都可以实现的适配器。
type xmlAdapter struct {
	JsonHandler
}

func (x *xmlAdapter) ProcessXml(xml string) {
	// 模拟对象转换：json -> xml
	json := []byte(xml)
	// 处理 xml 文件
	x.ProcessJson(json)
}

func TestAdapter(t *testing.T) {
	inputJson := []byte("json_file")
	// 原客户端
	handler := &jsonHandler{}
	handler.ProcessJson(inputJson)

	// 客户端接入适配器，利用 原有的 JsonHandler 处理 xml 文件。
	inputXml := "xml_file"
	adapter := &xmlAdapter{&jsonHandler{}}
	adapter.ProcessXml(inputXml)
}
