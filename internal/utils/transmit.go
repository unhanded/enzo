package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func PostTransmit(data any, addr string) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(b)

	res, httpErr := http.Post(addr, "application/json", reader)
	if httpErr != nil {
		return nil, httpErr
	}

	var resData = make([]byte, res.ContentLength)

	_, readErr := res.Body.Read(resData)
	if readErr != nil && len(resData) == 0 {
		return nil, fmt.Errorf("error in reading response: %s", readErr.Error())
	}
	retMsg := fmt.Sprintf("%s\nResponse:\n%s\n", res.Status, resData)
	return []byte(retMsg), nil
}
