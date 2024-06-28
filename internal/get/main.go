package get

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Run(host string, port int) {
	hostport := fmt.Sprintf("%s:%d", host, port)
	addr := fmt.Sprintf("http://%s/nodes", hostport)
	res, err := http.Get(addr)
	if err != nil {
		fmt.Printf("error occured getting nodes: %s\n", err.Error())
		os.Exit(1)
	}
	data := map[string]interface{}{}
	dec := json.NewDecoder(res.Body)
	decErr := dec.Decode(&data)
	if decErr != nil {
		fmt.Printf("error occured parsing response: %s\n", decErr.Error())
		os.Exit(1)
	}
	printable, mErr := json.MarshalIndent(data, "", "  ")
	if mErr != nil {
		fmt.Printf("error occured generating printout: %s\n", decErr)
		os.Exit(1)
	}
	fmt.Printf("Response from %s:\n%s\n", hostport, printable)
}
