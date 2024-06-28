package probe

import (
	"fmt"
	"os"
	"strings"

	"github.com/unhanded/enzo/internal/enzo"
	"github.com/unhanded/enzo/internal/utils"
)

func VerifiedOpenFile(fp string) (*enzo.EnzoItem, error) {
	if err := guardFileEnding(fp); err != nil {
		return nil, err
	}

	b, fileErr := os.ReadFile(fp)
	if fileErr != nil {
		return nil, fileErr
	}

	item, uErr := utils.Unmarshal[enzo.EnzoItem](b)
	if uErr != nil {
		return nil, uErr
	}

	if !item.Validate() {
		return nil, fmt.Errorf("invalid item definition")
	}

	return item, nil
}

func guardFileEnding(fp string) error {
	if !strings.HasSuffix(fp, ".item.enzo") {
		return fmt.Errorf("Error: expected .item.enzo file extension, got file with name %s", fp)
	}
	return nil
}
