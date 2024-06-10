package videos

type ValidatedVideo struct {
    video Video
}

func (vv ValidatedVideo) Video() Video {
    return vv.video;
}

func NewValidatedVideo(video *Video) *ValidatedVideo {
    err := video.validate()

    return &ValidatedVideo{
        video: video
    }
}
