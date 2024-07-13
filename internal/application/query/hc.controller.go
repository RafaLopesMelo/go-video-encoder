package query

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/common"
)

type HcController struct{}

func (c *HcController) Handle(request *common.Request) *common.Response {
	response := common.NewResponse()
	data := map[string]string{
		"status": "OK",
	}

	response.Json(data)

	return response
}

func NewHcController() *HcController {
	return &HcController{}
}
