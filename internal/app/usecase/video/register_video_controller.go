package video

import (
	"encoding/json"
	"net/http"
)

type RegisterVideoController struct {
	uc RegisterUseCase
}

type RegisterVideoResponse struct {
	ID        string `json:"id"`
	UploadURL string `json:"uploadUrl"`
}

func (c RegisterVideoController) Handle(w http.ResponseWriter, r *http.Request) {
	video, err := c.uc.Execute()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RegisterVideoResponse{
		ID:        video.ID,
		UploadURL: video.UploadURL,
	})
}

func NewRegisterVideoController(uc RegisterUseCase) RegisterVideoController {
	return RegisterVideoController{
		uc: uc,
	}
}
