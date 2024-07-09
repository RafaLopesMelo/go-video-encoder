package router

import "github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/common"

type HttpController interface {
	Handle(commonuest *common.Request) *common.Response
}
