package apply

import (
	"fmt"
	"os"

	"github.com/unhanded/enzo/internal/core"
	"github.com/unhanded/enzo/internal/utils"
)

func Run(fp string, host string, port int) {
	b, err := os.ReadFile(fp)
	if err != nil {
		fmt.Printf("error occured while reading file: %s\n", err)
		os.Exit(1)
	}
	node, uErr := utils.Unmarshal[core.NetNode](b)
	if uErr != nil {
		fmt.Printf("error occured while parsing file: %s\n", uErr)
		os.Exit(1)
	}
	res, err := ApplyNode(node, fmt.Sprintf("%s:%d", host, port))

	if err != nil {
		fmt.Printf("error occured while applying node: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s", res)
}

func ApplyNode(node *core.NetNode, serverAddr string) ([]byte, error) {
	return utils.PostTransmit(node, fmt.Sprintf("%s/node/%s", serverAddr, node.NodeId))
}
