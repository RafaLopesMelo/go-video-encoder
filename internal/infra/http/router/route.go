package router

import (
	"strings"

	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/common"
)

type Route struct {
	path       string
	controller HttpController
}

func (r *Route) Match(request *common.Request) bool {
	targetParts := filter(strings.Split(request.Path, "/"))
	currentParts := filter(strings.Split(r.path, "/"))

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

func NewRoute(path string, controller HttpController) *Route {
	return &Route{
		path:       path,
		controller: controller,
	}
}
