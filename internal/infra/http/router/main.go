package router

import (
	"net/http"

	"github.com/RafaLopesMelo/go-video-encoder/internal/app/usecase/hc"
	"github.com/RafaLopesMelo/go-video-encoder/internal/app/usecase/video"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/gateway/gcp"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/repo/pg"
)

func Setup() {
	http.HandleFunc("GET /hc", hc.NewCheckController().Handle)

	connection := pg.NewConnection()
	rv := pg.NewVideosRepo(connection)
	rr := pg.NewResourcesRepo(connection)
	u := gcp.NewUploader()
	uc := video.NewRegisterUseCase(rv, rr, u)

	http.HandleFunc("POST /videos", video.NewRegisterVideoController(uc).Handle)
}
