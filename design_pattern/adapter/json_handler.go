package adapter

import "log"

type JsonHandler interface {
	ProcessJson(json []byte)
}

type jsonHandler struct {
}

func (j *jsonHandler) ProcessJson(json []byte) {
	log.Println("json processing")
}
