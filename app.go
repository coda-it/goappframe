package goappframe

import (
	"errors"
	"github.com/coda-it/goutils/logger"
	"github.com/coda-it/gowebserver"
	"net/http"
)

// App - main application struct
type App struct {
	server    *gowebserver.WebServer
	internals Internals
}

func getServerAddress(port string) (string, error) {
	if port == "" {
		return "", errors.New("server port is not set")
	}
	return ":" + port, nil
}

func checkerHandler(domain string) func(req *http.Request) bool {
	return func(req *http.Request) bool {
		if domain != "" && req.Host == domain {
			return true
		}

		return false
	}
}

// New - creates new App instance
func New(i Internals) *App {
	addr, err := getServerAddress(i.Port)

	if err != nil {
		logger.Log("starting server failed: " + err.Error())
	}

	server := gowebserver.New(gowebserver.WebServerOptions{
		Port:           addr,
		StaticFilesURL: "/static/",
		StaticFilesDir: "public",
		Cert:           i.Config.Cert,
		CertPrvKey:     i.Config.CertPrvKey,
	}, i.NotFound, "/login")

	for _, appConfig := range i.Config.Apps {
		for _, moduleConfig := range appConfig.Modules {
			for _, moduleInstance := range i.Modules {
				if moduleInstance.Enabled || moduleConfig.ID == moduleInstance.ID {
					for _, r := range moduleInstance.Routes {
						server.Router.AddRoute(r.Path, r.Method, r.Protected, r.Handler, checkerHandler(appConfig.Domain))
					}
				}
			}
		}
	}

	server.AddDataSource(i.DataKey, i.Persistence)

	return &App{
		server,
		i,
	}
}

// Run - runs WebServer process
func (app *App) Run() {
	app.server.Run()
}
