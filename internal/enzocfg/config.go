package enzocfg

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type EnzoCmdCfg struct {
	Server string `yaml:"server"`
	Token  string `yaml:"token"`
}

func (ecfg *EnzoCmdCfg) Load() error {
	b, err := os.ReadFile(os.ExpandEnv("$HOME/.enzo/config"))
	if err != nil {
		return fmt.Errorf("error while loading config: %s", err.Error())
	}
	return yaml.Unmarshal(b, ecfg)
}

func (ecfg *EnzoCmdCfg) Defaults() error {
	ecfg.Server = "localhost:29451"
	ecfg.Token = ""
	return nil
}
