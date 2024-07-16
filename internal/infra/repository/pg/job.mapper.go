package pg

import (
	"database/sql"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type PersistenceJobDto struct {
	id            string
	status        entity.JobStatus
	kind          entity.JobKind
	video_id      string
	resource_id   sql.NullString
	depends_on_id sql.NullString
	error         string
}

func (dto PersistenceJobDto) ToEntity() *entity.Job {
	ID := vo.NewIDFromValue(dto.id)
	videoID := vo.NewIDFromValue(dto.video_id)

	var resourceID *vo.UniqueEntityID = nil
	if dto.resource_id.Valid {
		resourceID = vo.NewIDFromValue(dto.resource_id.String)
	}

	var dependsOnID *vo.UniqueEntityID = nil
	if dto.depends_on_id.Valid {
		dependsOnID = vo.NewIDFromValue(dto.depends_on_id.String)
	}

	return entity.LoadJob(entity.LoadJobDto{
		Status:      dto.status,
		Kind:        dto.kind,
		VideoID:     videoID,
		Error:       dto.error,
		ResourceID:  resourceID,
		DependsOnID: dependsOnID,
	}, ID)
}
