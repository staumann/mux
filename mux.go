package mux

import "net/http"

type Mux interface {
	HandleFunc(pattern string, handler func(writer http.ResponseWriter, request *http.Request))
	HasHandlerForPattern(pattern string) bool
	Handle(pattern string, handler http.Handler)
	GetAllRoutes() []string
}

// ExtendedMux is a struct extending http.ServerMux additionally it contains a map of all routes so it is easier to check all registered routes
type extendedMux struct {
	http.ServeMux
	routes map[string]bool
}

// HandleFunc does the same as http.ServerMux.HandleFuc. Additionally it stores the given path to a map so we can check all routes via HasHandlerForPattern
func (mux *extendedMux) HandleFunc(pattern string, handler func(writer http.ResponseWriter, request *http.Request)) {
	mux.ServeMux.HandleFunc(pattern, handler)
	mux.routes[pattern] = true
}

// HasHandlerForPattern is a helper method to check if the containing mux has a handler for the given path
func (mux *extendedMux) HasHandlerForPattern(pattern string) bool {
	_, ok := mux.routes[pattern]
	return ok
}

// Handle extends the handle function of http.ServerMux
func (mux *extendedMux) Handle(pattern string, handler http.Handler) {
	mux.ServeMux.Handle(pattern, handler)
	mux.routes[pattern] = true
}

// GetAllRoutes returns all registered routes
func (mux *extendedMux) GetAllRoutes() []string {
	routes := make([]string, len(mux.routes))
	counter := 0
	for k, _ := range mux.routes {
		routes[counter] = k
		counter = counter + 1
	}
	return routes
}

// GetNewMux return a new mux instance with no routes registered
func GetNewMux() *extendedMux {
	return &extendedMux{
		routes: make(map[string]bool),
	}
}
