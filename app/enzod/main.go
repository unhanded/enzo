package main

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/unhanded/enzo/internal/core"
	"github.com/unhanded/enzo/internal/utils"
	"github.com/unhanded/flownet/pkg/flownet"
)

type Net = flownet.FNet[core.AuxData]

func main() {
	network := core.NewNetwork()

	app := fiber.New(
		fiber.Config{
			ReadTimeout:           time.Second * 3,
			WriteTimeout:          time.Second * 5,
			ServerHeader:          "enzod",
			DisableStartupMessage: true,
		},
	)
	dg := prometheus.DefaultGatherer
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.HandlerFor(dg, promhttp.HandlerOpts{})))

	useNodeHandler(app, network)
	useProbeHandler(app, network)
	useStatsHandler(app, network)

	app.Listen(":8080")
}

func useNodeHandler(app *fiber.App, network Net) {
	grp := app.Group("/node")

	grp.Get("/", func(c *fiber.Ctx) error { return c.JSON(core.NodeCollection{Nodes: network.Nodes()}) })

	grp.Post("/:nodeId", func(c *fiber.Ctx) error {
		nodeId := c.Params("nodeId")
		if nodeId == "" {
			return c.Status(400).JSON(map[string]string{"error": "nodeId not found"})
		}
		nn := core.NetNode{}
		jErr := json.Unmarshal(c.Body(), &nn)
		if jErr != nil {
			return c.Status(400).JSON(map[string]string{"error": jErr.Error()})
		}

		nn.NodeId = nodeId
		for _, n := range network.Nodes() {
			if n.Id() == nn.Id() {
				network.RemoveNode(n.Id())
			}
		}
		network.AddNodes(&nn)
		return c.Status(200).JSON(map[string]string{"status": "ok"})
	})

	grp.Delete("/:nodeId", func(c *fiber.Ctx) error {
		var nId string = ""
		if nId := c.Params("nodeId"); nId == "" {
			return c.Status(400).JSON(map[string]interface{}{"error": "nodeId not found"})
		}
		if err := network.RemoveNode(nId); err != nil {
			return c.SendStatus(500)
		}
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

func useApiHandler(app *fiber.App) {
	routes := app.GetRoutes()
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{"routes": routes})
	})
}

func useProbeHandler(app *fiber.App, network Net) {
	app.Post("/probe", func(c *fiber.Ctx) error {
		item, err := utils.Unmarshal[core.EnzoItem](c.Body())

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
