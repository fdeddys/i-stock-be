package models

type Response struct {
	Rc        string      `json:"rc"`
	Msg       string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	TotalData interface{} `json:"totalData,omitempty"`
}
