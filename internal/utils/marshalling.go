package utils

import "encoding/json"

func Unmarshal[T any](data []byte) (*T, error) {
	val := new(T)
	err := json.Unmarshal(data, val)
	return val, err
}
