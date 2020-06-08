package domain

import (
	"bytes"
	"encoding/json"
	"im-server-go/enums"
	"log"
)

const (
	CODE_TAG string = "code"
	MSG_TAG  string = "msg"
	DATA_TAG string = "data"
)

type ApiResponse struct {
	resp map[string]interface{}
}

func NewApiResponse() *ApiResponse {
	return &ApiResponse{resp: make(map[string]interface{}, 3)}
}

func (ar *ApiResponse) Success(data interface{}) *ApiResponse {
	response := NewApiResponse()
	response.resp[CODE_TAG] = enums.SUCCESS
	response.resp[MSG_TAG] = "操作成功"
	if data != nil {
		response.resp[DATA_TAG] = data
	}
	return response
}

func (ar *ApiResponse) Warn(msg string, data interface{}) *ApiResponse {
	response := NewApiResponse()
	response.resp[CODE_TAG] = enums.WARN
	response.resp[MSG_TAG] = msg
	if data != nil {
		response.resp[DATA_TAG] = data
	}
	return response
}
func (ar *ApiResponse) Error(msg string, data interface{}) *ApiResponse {
	response := NewApiResponse()
	response.resp[CODE_TAG] = enums.ERROR
	response.resp[MSG_TAG] = msg
	if data != nil {
		response.resp[DATA_TAG] = data
	}
	return response
}

func (ar *ApiResponse) ToString() string {
	var s bytes.Buffer
	s.WriteString("[ApiResponse]:")
	for key, value := range ar.resp {
		s.WriteString("key:" + key + " value:" + value.(string))
	}
	return s.String()
}

//编码为json的字节数组类型
func (ar *ApiResponse) JsonBytes() []byte {
	data, err := json.Marshal(ar.resp)
	if err != nil {
		log.Println("[ApiResponse] map to json error :", err)
	}
	return data
}

//编码为json的string类型
func (ar *ApiResponse) JsonString() string {
	data, err := json.Marshal(ar.resp)
	if err != nil {
		log.Println("[ApiResponse] map to json error :", err)
	}
	return string(data)
}
