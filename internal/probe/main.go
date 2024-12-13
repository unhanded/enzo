package probe

import (
	"fmt"
	"os"

	"github.com/unhanded/enzo/internal/utils"
)

func Run(fp string, host string, port int) {
	item, verifyErr := VerifiedOpenFile(fp)
	if verifyErr != nil {
		fmt.Printf("encountered error: %s", verifyErr.Error())
		os.Exit(1)
	}
	addr := fmt.Sprintf("http://%s:%d/probe", host, port)
	res, transmitErr := utils.PostTransmit(item, addr)
	if transmitErr != nil {
		fmt.Printf("error on transmit: %s\n", transmitErr.Error())
		os.Exit(1)
	}
	fmt.Printf("%s", res)
}
