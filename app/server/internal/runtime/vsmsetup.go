package runtime

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/unhanded/enzo-vsm/internal/enzo"
	"github.com/unhanded/enzo-vsm/pkg/vsm"
)

func CreateVSM(m vsm.MeshNetwork) *enzo.Vsm {
	v := &enzo.Vsm{Network: m, Prm: prometheus.NewRegistry()}
	return v
}

func HandleApply(v *enzo.Vsm) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		data := c.Body()
		msg, err := v.Apply(data)
		if err != nil {
			return c.SendStatus(500)
		}
		_, writeErr := c.Status(200).WriteString(msg)
		return writeErr
	}

}

func HandleSubmit(v *enzo.Vsm) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		data := c.Body()
		err := v.Submit(data)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return c.SendStatus(500)
		}
		return c.SendStatus(200)
	}

}
