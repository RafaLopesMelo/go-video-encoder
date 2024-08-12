package video

import (
	"net/http"

	httperror "github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/error"
)

type RegisterVideoController struct {
	uc RegisterUseCase
}

type RegisterVideoResponse struct {
	ID        string `json:"id"`
	UploadURL string `json:"uploadUrl"`
}

func (c RegisterVideoController) Handle(w http.ResponseWriter, r *http.Request) any {
	video, err := c.uc.Execute()

	if err != nil {
		return httperror.NewFromDomain(err)
	}

	return RegisterVideoResponse{
		ID:        video.ID,
		UploadURL: video.UploadURL,
	}
}

func NewRegisterVideoController(uc RegisterUseCase) RegisterVideoController {
	return RegisterVideoController{
		uc: uc,
	}
}
