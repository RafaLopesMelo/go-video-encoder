package router

import "github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/common"

type HttpController interface {
	Handle(request *common.Request) *common.Response
}
