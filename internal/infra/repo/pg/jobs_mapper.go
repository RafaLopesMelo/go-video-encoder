package pg

import (
	"database/sql"

	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/entity"
	"github.com/RafaLopesMelo/go-video-encoder/internal/domain/vo"
)

type jobMapper struct {
}

func (m jobMapper) ToEntity(dto persistenceJobDto) entity.Job {
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

func (m jobMapper) ToPersistence(entity entity.Job) persistenceJobDto {
	dto := persistenceJobDto{}

	dto.id = entity.ID.Value()
	dto.status = entity.Status
	dto.kind = entity.Kind
	dto.video_id = entity.VideoID.Value()
	dto.error = entity.Error

	dto.resource_id = sql.NullString{
		String: "",
		Valid:  false,
	}
	if entity.ResourceID != nil {
		dto.resource_id.Valid = true
		dto.resource_id.String = entity.ResourceID.Value()
	}

	dto.depends_on_id = sql.NullString{
		String: "",
		Valid:  false,
	}
	if entity.DependsOnID != nil {
		dto.depends_on_id.Valid = true
		dto.depends_on_id.String = entity.DependsOnID.Value()
	}

	return dto
}

func newJobMapper() *jobMapper {
	return &jobMapper{}
}
