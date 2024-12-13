package del

import (
	"fmt"
	"net/http"
	"os"
)

func Run(name string, host string, port int) {
	addr := fmt.Sprintf("http://%s:%d/node/%s", host, port, name)

	req, err := http.NewRequest("DELETE", addr, nil)
	if err != nil {
		fmt.Printf("error occured creating request: %s\n", err)
		os.Exit(1)
	}
	c := http.Client{}
	res, cErr := c.Do(req)
	if cErr != nil {
		fmt.Printf("error occured sending request: %s\n", err)
		os.Exit(1)
	}
	if res.StatusCode != 200 {
		fmt.Printf("unknown error deleting node %s..\n", name)
		os.Exit(1)
	}
	fmt.Printf("Successfully deleted node %s\n", name)
}
