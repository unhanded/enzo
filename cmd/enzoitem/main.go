package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/unhanded/enzo/internal/enzocfg"
	"github.com/unhanded/enzo/pkg/enzo"
)

func main() {
	var filepath string
	flag.StringVar(&filepath, "f", "", "Enzo file")

	var interval int
	flag.IntVar(&interval, "i", 0, "Interval in milliseconds")

	flag.Parse()

	cfg := enzocfg.EnzoCmdCfg{}
	err := cfg.Load()
	if err != nil {
		fmt.Printf("Error loading config, using default. Error: %s\n", err.Error())
		cfg.Defaults()
	}
	if interval > 0 {
		requestSubmitInterval(cfg.Server, filepath, interval)
		return
	} else {
		requestSubmit(cfg.Server, filepath)
	}
}

func requestSubmitInterval(addr string, fp string, interval int) {
	for {
		requestSubmit(addr, fp)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func requestSubmit(addr string, fp string) {
	// Read file
	data, err := os.ReadFile(fp)

	if err != nil {
		fmt.Printf("Error reading file: %s\n", err.Error())
		return
	}

	cfg := enzo.WorkItemConfig{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Printf("Error unmarshalling json: %s\n", err.Error())
		return
	}

	rdr := bytes.NewBuffer(data)

	res, reqErr := http.Post("http://"+addr+"/submit", "text/plain", rdr)
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
