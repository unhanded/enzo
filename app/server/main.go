package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/unhanded/enzo-vsm/app/server/internal/runtime"
	"github.com/unhanded/enzo-vsm/internal/enzo"
)

func main() {
	app := fiber.New(fiber.Config{ReadTimeout: time.Second * 3, WriteTimeout: time.Second * 5})
	m := enzo.NewMesh()
	v := runtime.CreateVSM(m)
	initErr := v.Init()

	if initErr != nil {
		panic(initErr)
	}

	app.Post("/apply", runtime.HandleApply(v))
	app.Post("/submit", runtime.HandleSubmit(v))
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.HandlerFor(v.Prm, promhttp.HandlerOpts{})))
	app.Listen(":29451")
}
