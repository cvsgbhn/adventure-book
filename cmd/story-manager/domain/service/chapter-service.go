package service

import "stories/domain/entity"

type ChapterService struct {
	repo ChapterStorager
}

func NewChapterService(repo ChapterStorager) *ChapterService {
	return &ChapterService{repo: repo}
}

func (cs *ChapterService) ByArc(arc string) (entity.Chapter, error) {
	return cs.repo.ByArc(arc)
}
