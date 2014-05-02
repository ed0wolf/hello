package web

import (
	"log"
	"net/http"

	"github.com/zenazn/goji/bind"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"

	"ed0wolf/hello/web/controllers"
	"ed0wolf/hello/web/middleware/assets"
)

var app *web.Mux

func init() {
	app = web.New()

	//Add middleware
	handler := &assets.AssetsHandler{}
	app.Use(handler.HandleAssets)

	//Add routes
	controllers.AddRoutes(app)
}

func Start() {
	http.Handle("/", app)

	listener := bind.Default()
	log.Println("Starting Goji on", listener.Addr())

	bind.Ready()

	err := graceful.Serve(listener, http.DefaultServeMux)

	if err != nil {
		log.Fatal(err)
	}
}
