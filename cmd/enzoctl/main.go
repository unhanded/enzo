package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/unhanded/enzo/internal/enzocfg"
	"github.com/unhanded/enzo/pkg/enzo"
	"gopkg.in/yaml.v3"
)

func requestApply(addr string, fp string) {
	// Read file
	data, err := os.ReadFile(fp)

	if err != nil {
		fmt.Printf("Error reading file: %s\n", err.Error())
		return
	}

	cfg := enzo.WorkcenterConfig{}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Printf("Error unmarshalling yaml: %s\n", err.Error())
		return
	}

	rdr := bytes.NewBuffer(data)

	res, reqErr := http.Post("http://"+addr+"/apply", "text/plain", rdr)
	if reqErr != nil {
		fmt.Printf("error making request %s\n", reqErr.Error())
		return
	}
	fmt.Printf("Code: %s\n", res.Status)
	content := make([]byte, res.ContentLength)
	res.Body.Read(content)
	fmt.Printf("Response: %s\n", content)
	return
}

func main() {
	cfg := enzocfg.EnzoCmdCfg{}
	err := cfg.Load()
	if err != nil {
		fmt.Printf("Error loading config, using default. Error: %s\n", err.Error())
		cfg.Defaults()
	}

	var filepath string
	flag.StringVar(&filepath, "f", "", "Enzo file")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("No verb provided")
		return
	}
	verb := args[0]
	switch verb {
	case "apply":
		requestApply(cfg.Server, filepath)
	default:
		fmt.Println("Unknown verb")
	}
}
