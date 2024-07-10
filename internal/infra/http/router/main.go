package router

import (
	"net/http"

	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/common"
)

type Router struct {
	routes map[string][]Route
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	request := common.NewRequestFromVanilla(req)
	controller := r.matchHandler(request)

	if controller == nil {
		return
	}

	response := (*controller).Handle(request)
	response.Send(w)
}

func (r *Router) matchHandler(request *common.Request) *HttpController {
	routes := r.routes[request.Method]

	if routes == nil {
		return nil
	}

	for _, route := range routes {
		match := route.Match(request)

		if match == false {
			continue
		}

		return &route.controller
	}

	return nil
}

func (r *Router) Get(path string, controller HttpController) {
	r.routes["GET"] = append(r.routes["GET"], *NewRoute(path, controller))
}

func (r *Router) Post(path string, controller HttpController) {
	r.routes["POST"] = append(r.routes["POST"], *NewRoute(path, controller))
}

func (r *Router) Put(path string, controller HttpController) {
	r.routes["PUT"] = append(r.routes["PUT"], *NewRoute(path, controller))
}

func (r *Router) Delete(path string, controller HttpController) {
	r.routes["DELETE"] = append(r.routes["DELETE"], *NewRoute(path, controller))
}

func (r *Router) Patch(path string, controller HttpController) {
	r.routes["PATCH"] = append(r.routes["PATCH"], *NewRoute(path, controller))
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string][]Route),
	}
}
