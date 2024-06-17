package enzoctl

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type EnzoCtlCfg struct {
	Server string `yaml:"server"`
	Token  string `yaml:"token"`
}

func (ecfg *EnzoCtlCfg) Load() error {
	b, err := os.ReadFile(os.ExpandEnv("$HOME/.enzo/config"))
	if err != nil {
		return fmt.Errorf("error while loading config: %s", err.Error())
	}
	return yaml.Unmarshal(b, ecfg)
}
