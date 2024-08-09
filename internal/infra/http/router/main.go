package router

import (
	"net/http"

	"github.com/RafaLopesMelo/go-video-encoder/internal/app/usecase/hc"
	"github.com/RafaLopesMelo/go-video-encoder/internal/app/usecase/video"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/gateway/gcp"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repo/pg"
)

type Router struct {
	middlewares []Middleware
}

type Handler func(http.ResponseWriter, *http.Request) any
type Middleware func(h Handler) Handler

func (r *Router) Use(middleware Middleware) {
	r.middlewares = append(r.middlewares, middleware)
}

func (r *Router) Setup() {
	r.addRoute("GET /hc", hc.NewCheckController().Handle)

	connection := pg.NewConnection()
	rv := pg.NewVideosRepo(connection)
	rr := pg.NewResourcesRepo(connection)
	u := gcp.NewUploader()
	uc := video.NewRegisterUseCase(rv, rr, u)

	r.addRoute("POST /videos", video.NewRegisterVideoController(uc).Handle)
}

func (r *Router) httpWrapper(handle Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handle(w, r)
	}
}

func (r *Router) addRoute(pattern string, handler Handler) {
	for _, middleware := range r.middlewares {
		handler = middleware(handler)
	}

	wrapped := r.httpWrapper(handler)
	http.HandleFunc(pattern, wrapped)
}

func NewRouter() *Router {
	return &Router{}
}
