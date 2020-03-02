package scheduler

import "encoding/json"

type EncodeService interface {
	Encode(interface{}) (string, error)
}

type JsonEncoder struct {
}

func NewJsonEncoder() *EncodeService {
	var encoder EncodeService
	encoder = &JsonEncoder{}
	return &encoder
}

func (j *JsonEncoder) Encode(item interface{}) (string, error) {
	b, err := json.Marshal(item)
	var itemJson = string(b)
	return itemJson, err
}
