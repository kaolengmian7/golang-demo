package json_number

import (
	"encoding/json"
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestJsonUnmarshal(t *testing.T) {
	paramsJson := "{\"id1\":1,\"id2\":1666104922968}"
	res := make(map[string]interface{})
	if err := json.Unmarshal([]byte(paramsJson), &res); err != nil {
		log.Panic(err.Error())
	}
	log.Printf("%+v", res)
}

func TestJsonDecoder(t *testing.T) {
	paramsJson := "{\"id1\":1,\"id2\":1666104922968}"
	res := make(map[string]interface{})

	decoder := json.NewDecoder(strings.NewReader(paramsJson))
	decoder.UseNumber() // 防止 Golang 将大整型转成Float64类型导致精度丢失
	if err := decoder.Decode(&res); err != nil {
		log.Panic(err.Error())
	}

	log.Printf("%+v", res)
}

func TestJsonDecoderNumberType(t *testing.T) {
	paramsJson := "{\"id1\":1,\"id2\":1666104922968}"
	res := make(map[string]interface{})

	decoder := json.NewDecoder(strings.NewReader(paramsJson))
	decoder.UseNumber() // 防止 Golang 将大整型转成Float64类型导致精度丢失
	if err := decoder.Decode(&res); err != nil {
		log.Panic(err.Error())
	}
	log.Printf("%+v", res)

	// Golang 会将整型转换成 json.Number 类型
	typeOf := reflect.TypeOf(res["id1"])
	log.Println(typeOf.String())
	typeOf = reflect.TypeOf(res["id2"])
	log.Println(typeOf.String())

	// 如果想获得 Int64 类型，需要使用 json.Number 的成员方法转换数据类型。
	log.Println(res["id2"].(json.Number).Int64())
}
