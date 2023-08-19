package main

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"log"
	"net/http"
	"recipio-admin/handlers"
)

func main() {
	app := pocketbase.New()
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method:  http.MethodPost,
			Path:    "/api/send-invite",
			Handler: handlers.SendInvite(app),
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})
		return nil
	})
	app.OnRecordAfterCreateRequest("users").Add(func(e *core.RecordCreateEvent) error {
		e.Record.Set("emailVisibility", true)
		return nil
	})
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
