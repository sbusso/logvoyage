package main

import (
	"html/template"
	"reflect"
	"runtime"
	"time"

	"github.com/Unknwon/com"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"

	"github.com/firstrow/logvoyage/web/context"
	"github.com/firstrow/logvoyage/web/middleware"
	"github.com/firstrow/logvoyage/web/routers/home"
	"github.com/firstrow/logvoyage/web/routers/profile"
	"github.com/firstrow/logvoyage/web/routers/sources"
	"github.com/firstrow/logvoyage/web/routers/users"
	"github.com/firstrow/logvoyage/web/widgets"
	"github.com/firstrow/logvoyage/web_socket"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Template methods
	templateFunc := template.FuncMap{
		"FormatTimeToHuman": func(s ...string) string {
			if len(s) > 0 {
				t, _ := time.Parse(time.RFC3339Nano, s[0])
				return t.Format("2006-01-02 15:04:05") + " UTC"
			} else {
				return "Unknown"
			}
		},
		"isEmpty": func(i interface{}) bool {
			switch reflect.TypeOf(i).Kind() {
			case reflect.Slice:
				v := reflect.ValueOf(i)
				return v.Len() == 0
			}
			return true
		},
		"eq":                       reflect.DeepEqual,
		"isSliceContainsStr":       com.IsSliceContainsStr,
		"renderSourceGroupsWidget": widgets.NewSourceGroups,
	}

	m := martini.Classic()
	// Template
	m.Use(render.Renderer(render.Options{
		Funcs:  []template.FuncMap{templateFunc},
		Layout: "layouts/main",
	}))
	// Serve static files
	m.Use(martini.Static("../static"))
	// Sessions
	store := sessions.NewCookieStore([]byte("super_secret_key"))
	m.Use(sessions.Sessions("default", store))

	m.Use(context.Contexter)

	// Routes
	m.Any("/register", middleware.RedirectIfAuthorized, users.Register)
	m.Any("/login", middleware.RedirectIfAuthorized, users.Login)
	// Auth routes
	m.Get("/dashboard", middleware.Authorize, home.Index)
	m.Get("/view", middleware.Authorize, home.View)
	m.Any("/profile", middleware.Authorize, profile.Index)
	// Sources
	m.Group("/sources", func(r martini.Router) {
		r.Any("", sources.Index)
		r.Any("/new", sources.New)
		r.Any("/edit/:id", sources.Edit)
		r.Any("/delete/:id", sources.Delete)
		// Types
		r.Any("/types", sources.Types)
		r.Any("/types/delete/:name", sources.DeleteType)
	}, middleware.Authorize)

	go web_socket.StartServer()
	m.Run()
}