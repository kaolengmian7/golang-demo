package adapter

import "testing"

// demo2：多重继承的编程语言可以实现更优雅的适配器
type adapter struct {
	jsonHandler JsonHandler
	xmlHandler  XmlHandler
}

func TestAdapterInGolang(t *testing.T) {
	inputJson := []byte("json_file")
	inputXml := "xml_file"
	// 原客户端
	handler := &jsonHandler{}
	handler.ProcessJson(inputJson)

	// 客户端接入适配器
	adapter := &adapter{&jsonHandler{}, &xmlHandler{}}
	adapter.jsonHandler.ProcessJson(inputJson) // 想用 json 用 json
	adapter.xmlHandler.ProcessXml(inputXml)    // 想用 xml 用 xml
}

//// 自定义
//type adapter struct {
//	jsonHandler JsonHandler
//}
//
//func (x *adapter) ProcessXml(xml string) {
//	// 自定义处理逻辑
//}
