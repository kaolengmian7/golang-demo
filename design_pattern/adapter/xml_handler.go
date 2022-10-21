package adapter

import "log"

type XmlHandler interface {
	ProcessXml(xml string)
}

type xmlHandler struct{}

func (x *xmlHandler) ProcessXml(xml string) {
	log.Println("xml processing")
}
