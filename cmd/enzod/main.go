package main

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/unhanded/enzo/internal/enzo"
)

func main() {
	network := enzo.NewNetwork()
	dg := prometheus.DefaultGatherer

	app := fiber.New(
		fiber.Config{
			ReadTimeout:           time.Second * 3,
			WriteTimeout:          time.Second * 5,
			ServerHeader:          "enzo",
			DisableStartupMessage: true,
		},
	)

	app.Get("/api", func(c *fiber.Ctx) error {
		nodes := network.Nodes()
		nodeIds := []string{}

		for _, n := range nodes {
			nodeIds = append(nodeIds, n.Id())
		}

		d := map[string]interface{}{
			"nodeCount": len(nodes),
			"nodes":     nodeIds,
		}
		return c.Status(200).JSON(d)
	})

	app.Post("/api", func(c *fiber.Ctx) error {
		nn := &enzo.NetNode{}
		jErr := json.Unmarshal(c.Body(), &nn)
		if jErr != nil {
			return c.Status(400).JSON(map[string]string{"error": jErr.Error()})
		}
		for _, n := range network.Nodes() {
			if n.Id() == nn.Id() {
				network.RemoveNode(n.Id())
			}
		}
		network.AddNodes(nn)
		return c.Status(200).JSON(map[string]string{"status": "ok"})
	})

	app.Get("/metrics", adaptor.HTTPHandler(promhttp.HandlerFor(dg, promhttp.HandlerOpts{})))

	app.Listen(":29451")
}
