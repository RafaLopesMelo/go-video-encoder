package router

import (
	"net/http"
	"strings"

	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/common"
)

type Router struct {
	routes map[string]map[string]HttpController
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

	for route, handler := range routes {
		match := r.doesRoutesMatch(route, request)

		if match == false {
			continue
		}

		return &handler
	}

	return nil
}

func (r *Router) doesRoutesMatch(current string, request *common.Request) bool {
	targetParts := filter(strings.Split(request.Path, "/"))
	currentParts := filter(strings.Split(current, "/"))

	if len(currentParts) != len(targetParts) {
		return false
	}

	params := make(map[string]string)

	for i := 0; i < len(targetParts); i++ {
		targetPart := targetParts[i]
		currentPart := currentParts[i]

		if currentPart[0] == ':' {
			key := strings.Replace(currentPart, ":", "", 1)
			params[key] = targetPart
			continue
		}

		if targetPart != currentPart {
			return false
		}
	}

	for key, value := range params {
		request.SetParam(key, value)
	}

	return true
}

func filter(parts []string) []string {
	var result []string
	for _, v := range parts {
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}

func (r *Router) Get(path string, controller HttpController) {
	if r.routes["GET"] == nil {
		r.routes["GET"] = make(map[string]HttpController)
	}

	r.routes["GET"][path] = controller
}

func (r *Router) Post(path string, controller HttpController) {
	if r.routes["POST"] == nil {
		r.routes["POST"] = make(map[string]HttpController)
	}

	r.routes["POST"][path] = controller
}

func (r *Router) Put(path string, controller HttpController) {
	if r.routes["PUT"] == nil {
		r.routes["PUT"] = make(map[string]HttpController)
	}

	r.routes["PUT"][path] = controller
}

func (r *Router) Delete(path string, controller HttpController) {
	if r.routes["DELETE"] == nil {
		r.routes["DELETE"] = make(map[string]HttpController)
	}

	r.routes["DELETE"][path] = controller
}

func (r *Router) Patch(path string, controller HttpController) {
	if r.routes["PATCH"] == nil {
		r.routes["PATCH"] = make(map[string]HttpController)
	}

	r.routes["PATCH"][path] = controller
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]HttpController),
	}
}
