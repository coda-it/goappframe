package goappframe

import (
	"errors"
	"net/http"

	"github.com/coda-it/goappframe/config"
	"github.com/coda-it/goappframe/module"
	"github.com/coda-it/goutils/logger"
	"github.com/coda-it/gowebserver"
	"github.com/coda-it/gowebserver/router"
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
		if domain == "" || (domain != "" && req.Host == domain) {
			return true
		}

		return false
	}
}

func registerModuleRoutes(addRoute router.AddRouteFunc, cnf config.Config, modules []module.Module) {
	for _, appConfig := range cnf.Apps {
		for _, moduleInstance := range modules {
			enabled := moduleInstance.Enabled

			if !enabled {
				for _, moduleConfig := range appConfig.Modules {
					if moduleConfig.ID == moduleInstance.ID {
						enabled = true
						break
					}
				}
			}

			if enabled {
				for _, r := range moduleInstance.Routes {
					addRoute(r.Path, r.Method, r.Protected, r.Handler, checkerHandler(appConfig.Domain))
				}
			}
		}
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

	registerModuleRoutes(server.Router.AddRoute, i.Config, i.Modules)

	server.AddDataSource(i.DataKey, i.Persistence)

	return &App{
		server,
		i,
	}
}

// Reconfigure - applies a new config at runtime, atomically rebuilding all
// routes; server-level options (port, certificates) still require a restart
func (app *App) Reconfigure(cnf config.Config) {
	app.internals.Config = cnf

	app.server.Router.ReplaceRoutes(func(addRoute router.AddRouteFunc) {
		registerModuleRoutes(addRoute, cnf, app.internals.Modules)
	})
}

// Run - runs WebServer process
func (app *App) Run() {
	app.server.Run()
}
