package apply

import (
	"fmt"
	"os"

	"github.com/unhanded/enzo/internal/enzo"
	"github.com/unhanded/enzo/internal/utils"
)

func Run(fp string, host string, port int) {
	b, err := os.ReadFile(fp)
	if err != nil {
		fmt.Printf("error occured while reading file: %s\n", err)
		os.Exit(1)
	}
	node, uErr := utils.Unmarshal[enzo.NetNode](b)
	if uErr != nil {
		fmt.Printf("error occured while parsing file: %s\n", uErr)
		os.Exit(1)
	}
	res, ptErr := utils.PostTransmit(
		node,
		fmt.Sprintf("http://%s:%d/node", host, port),
	)
	if ptErr != nil {
		fmt.Printf("error occured: %s\n", ptErr)
		os.Exit(1)
	}
	fmt.Printf("%s", res)
}
