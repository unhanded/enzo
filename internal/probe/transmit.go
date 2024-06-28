package probe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/unhanded/enzo/internal/enzo"
)

func Transmit(item *enzo.EnzoItem, host string, port int) error {
	b, err := json.Marshal(item)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(b)

	addr := fmt.Sprintf("http://%s:%d/probe", host, port)
	res, httpErr := http.Post(addr, "application/json", reader)
	if httpErr != nil {
		return httpErr
	}

	var resData = []byte{}

	_, readErr := res.Body.Read(resData)
	if readErr != nil && len(resData) == 0 {
		return fmt.Errorf("error in reading response: %s", readErr.Error())
	}
	fmt.Printf("%s\nResponse:\n%s\n", res.Status, resData)
	return nil
}
