package hc

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/common"
)

type HcController struct{}

func (c *HcController) Handle(request *common.Request) *common.Response {
	response := common.NewResponse()
	data := map[string]string{
		"message": "Hello World!",
	}

	response.Json(data)

	return response
}

func NewHcController() *HcController {
	return &HcController{}
}
