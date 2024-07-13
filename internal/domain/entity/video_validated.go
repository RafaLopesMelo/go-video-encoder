package entity

type ValidatedVideo struct {
	video Video
}

func (vv ValidatedVideo) Video() Video {
	return vv.video
}

func NewValidatedVideo(video Video) (*ValidatedVideo, error) {
	err := video.validate()

	if err != nil {
		return nil, err
	}

	return &ValidatedVideo{
		video: video,
	}, nil
}
