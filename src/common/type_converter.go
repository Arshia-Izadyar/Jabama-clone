package common

import "encoding/json"



func TypeConvert[T any](model interface{}) (*T, error){
	var result T
	bs, err := json.Marshal(&model)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bs, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}