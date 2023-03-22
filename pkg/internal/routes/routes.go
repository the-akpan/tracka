package routes

import (
	"internal/middlewares"
	"internal/routes/admin"
	"log"
	"strings"

	"github.com/gorilla/mux"
)

func Init(debug bool) *mux.Router {
	router := mux.NewRouter()

	router.Use(middlewares.LoggingMiddleware)
	router.Use(middlewares.AcceptedContentTypeMiddleware)
	router.Use(middlewares.ResponseContentTypeMiddleware)

	admin.Register(router)

	if debug {
		walk(router)
	}

	return router
}

func walk(router *mux.Router) {
	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			log.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			log.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			log.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			log.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			log.Println("Methods:", strings.Join(methods, ","))
		}
		log.Println()
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
