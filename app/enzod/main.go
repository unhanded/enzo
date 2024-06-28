package main

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/unhanded/enzo/internal/enzo"
	"github.com/unhanded/enzo/internal/utils"
	"github.com/unhanded/flownet/pkg/flownet"
)

type Net = flownet.FNet[enzo.AuxData]

func main() {
	network := enzo.NewNetwork()

	app := fiber.New(
		fiber.Config{
			ReadTimeout:           time.Second * 3,
			WriteTimeout:          time.Second * 5,
			ServerHeader:          "enzo",
			DisableStartupMessage: true,
		},
	)

	useNodeGetHandler(app, network)
	useNodePostHandler(app, network)
	useNodeDeleteHandler(app, network)
	useProbeHandler(app, network)
	useStatsHandler(app, network)

	usePrometheus(app)

	app.Listen(":8080")
}

func useNodeGetHandler(app *fiber.App, network Net) {
	app.Get("/nodes", func(c *fiber.Ctx) error {
		nodes := enzo.NodeCollection{Nodes: network.Nodes()}

		b, err := json.Marshal(nodes)
		if err != nil {
			c.Status(500)
			return err
		}
		c.Write(b)
		return c.SendStatus(200)
	})
}

func useStatsHandler(app *fiber.App, network Net) {
	app.Get("/stats", func(c *fiber.Ctx) error {
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
}

func useNodePostHandler(app *fiber.App, network Net) {
	app.Post("/node", func(c *fiber.Ctx) error {
		nn := &enzo.NetNode{}
		jErr := json.Unmarshal(c.Body(), &nn)
		if jErr != nil {
			return c.Status(400).JSON(map[string]string{"error": jErr.Error()})
		}
		if nn.NodeId == "" {
			return c.SendStatus(400)
		}
		for _, n := range network.Nodes() {
			if n.Id() == nn.Id() {
				network.RemoveNode(n.Id())
			}
		}
		network.AddNodes(nn)
		return c.Status(200).JSON(map[string]string{"status": "ok"})
	})
}

func useNodeDeleteHandler(app *fiber.App, network Net) {
	app.Delete("/node/:nodeId", func(c *fiber.Ctx) error {
		nId := c.Params("nodeId")
		if nId == "" {
			return c.Status(400).JSON(map[string]interface{}{"error": "nodeId not found"})
		}
		network.RemoveNode(nId)
		return c.SendStatus(200)
	})
}

func useProbeHandler(app *fiber.App, network Net) {
	app.Post("/probe", func(c *fiber.Ctx) error {
		item, err := utils.Unmarshal[enzo.EnzoItem](c.Body())

		if err != nil {
			return c.Status(400).JSON(map[string]string{"error": err.Error()})
		}

		result, evalErr := network.Eval(item)
		if evalErr != nil {
			return c.Status(400).JSON(map[string]interface{}{"error": evalErr.Error()})
		}

		jsonErr := c.JSON(result.RelativeFlow())
		if jsonErr != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(200)
	})
}

func usePrometheus(app *fiber.App) {
	dg := prometheus.DefaultGatherer
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.HandlerFor(dg, promhttp.HandlerOpts{})))
}
