package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"time"
)

func helloworld(c *fiber.Ctx) error {
	return c.SendString("Hello world")

}
func main() {
	app := fiber.New()

	app.Get("/", helloworld)
	app.Listen(":3000")

}

func Run(ctx context.Context.Context) error {
	log := logger.GetFromContextOrDefault(ctx)
	log.Debug().Msg("starting teams sync worker")

	ticker := time.NewTicker(w.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Debug().Msgf("running teams sync in %s", w.Interval.String())
			err := w.Sync(ctx)
			if err != nil {
				log.Err(err).Msg("teams sync")
			}
		case <-ctx.Done():
			log.Debug().Msg("teams sync stopped because context has been canceled")

			return ctx.Err()
		}
	}

	return nil
}