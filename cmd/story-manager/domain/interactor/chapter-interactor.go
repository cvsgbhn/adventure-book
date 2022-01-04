package interactor

import (
	"stories/domain/entity"
	"stories/domain/service"
)

type ChapterInteractor struct {
	cs *service.ChapterService
}

func NewChapterInteractor(cs *service.ChapterService) *ChapterInteractor {
	return &ChapterInteractor{cs: cs}
}

func (ci *ChapterInteractor) ChapterByArc(arc string) (entity.Chapter, error) {
	chapter, err := ci.cs.ByArc(arc)
	if err != nil {
		return entity.Chapter{}, err
	}

	return chapter, nil
}
