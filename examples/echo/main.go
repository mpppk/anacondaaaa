package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mpppk/anacondaaaa"

	"github.com/labstack/echo"
)

func main() {
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")

	e := echo.New()
	e.GET("/", generateCRCTestHandler(consumerSecret))
	e.POST("/", accountActivityEventHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func generateCRCTestHandler(twitterConsumerSecret string) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(anacondaaaa.CRCRequest)
		if err := c.Bind(req); err != nil {
			return err
		}

		response := &anacondaaaa.CRCResponse{
			ResponseToken: anacondaaaa.CreateCRCToken(req.CRCToken, twitterConsumerSecret),
		}
		return c.JSON(http.StatusOK, response)
	}
}

func accountActivityEventHandler(c echo.Context) error {
	events := new(anacondaaaa.AccountActivityEvent)
	if err := c.Bind(events); err != nil {
		return err
	}

	if events.GetEventName() == anacondaaaa.TweetCreateEventsEventName {
		return c.String(http.StatusOK, fmt.Sprintf(
			"tweet event is arrived. first tweet content: %#v", events.TweetCreateEvents[0]))
	}

	return c.NoContent(http.StatusNoContent)
}
