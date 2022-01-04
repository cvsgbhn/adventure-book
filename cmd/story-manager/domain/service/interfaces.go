package service

import "stories/domain/entity"

type ChapterStorager interface {
	ByArc(arc string) (entity.Chapter, error)
}
