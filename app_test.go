package goappframe

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coda-it/goappframe/config"
	"github.com/coda-it/goappframe/module"
	"github.com/coda-it/goappframe/route"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
)

func contentHandler(content string) router.ControllerHandler {
	return func(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
		w.Write([]byte(content))
	}
}

func configWithModules(moduleIDs ...string) config.Config {
	modules := make([]config.Module, 0)

	for _, moduleID := range moduleIDs {
		modules = append(modules, config.Module{ID: moduleID})
	}

	return config.Config{
		Apps: []config.App{
			{
				ID:      "default",
				Modules: modules,
			},
		},
	}
}

func TestRoutesAreServedForAliasDomains(t *testing.T) {
	postModule := module.Module{
		ID: "post",
		Routes: []route.Route{
			{
				Path:    "/post",
				Method:  "GET",
				Handler: contentHandler("post handler"),
			},
		},
	}

	cnf := config.Config{
		Apps: []config.App{
			{
				ID:      "default",
				Domain:  "app.example.com",
				Domains: []string{"www.example.com"},
				Modules: []config.Module{{ID: "post"}},
			},
		},
	}

	app := New(Internals{
		Port:     "8080",
		Config:   cnf,
		Modules:  []module.Module{postModule},
		NotFound: contentHandler("not found"),
	})

	cases := map[string]string{
		"app.example.com":      "post handler",
		"www.example.com":      "post handler",
		"unlisted.example.com": "not found",
	}

	for host, expectedBody := range cases {
		request, _ := http.NewRequest(http.MethodGet, "/post", nil)
		request.Host = host
		writer := httptest.NewRecorder()
		app.server.Router.Route(writer, request)

		if writer.Body.String() != expectedBody {
			t.Errorf("host %q routed to %q, expected %q", host, writer.Body.String(), expectedBody)
		}
	}
}

func TestReconfigure(t *testing.T) {
	postModule := module.Module{
		ID: "post",
		Routes: []route.Route{
			{
				Path:      "/post",
				Method:    "GET",
				Handler:   contentHandler("post handler"),
				Protected: false,
			},
		},
	}

	app := New(Internals{
		Port:     "8080",
		Config:   configWithModules("post"),
		Modules:  []module.Module{postModule},
		NotFound: contentHandler("not found"),
	})

	request, _ := http.NewRequest(http.MethodGet, "/post", nil)
	writer := httptest.NewRecorder()
	app.server.Router.Route(writer, request)

	if writer.Body.String() != "post handler" {
		t.Errorf("module enabled in config should be routed, got %q", writer.Body.String())
	}

	app.Reconfigure(configWithModules())

	request, _ = http.NewRequest(http.MethodGet, "/post", nil)
	writer = httptest.NewRecorder()
	app.server.Router.Route(writer, request)

	if writer.Body.String() != "not found" {
		t.Errorf("module removed from config should not be routed, got %q", writer.Body.String())
	}

	app.Reconfigure(configWithModules("post"))

	request, _ = http.NewRequest(http.MethodGet, "/post", nil)
	writer = httptest.NewRecorder()
	app.server.Router.Route(writer, request)

	if writer.Body.String() != "post handler" {
		t.Errorf("module re-added to config should be routed again, got %q", writer.Body.String())
	}
}
