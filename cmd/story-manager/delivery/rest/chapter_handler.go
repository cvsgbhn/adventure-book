package rest

import (
	"fmt"
	"net/http"
	"stories/domain/interactor"
)

type ChapterHandler struct {
	ci *interactor.ChapterInteractor
}

func NewChapterHandler(ci *interactor.ChapterInteractor) *ChapterHandler {
	return &ChapterHandler{
		ci: ci,
	}
}

func (h *ChapterHandler) GetChapter(w http.ResponseWriter, r *http.Request) {
	chapter, err := h.ci.ChapterByArc(r.URL.Query()["arc"][0])
	if err != nil {
		fmt.Println(err)
		return
	}

	SendOK(w, 200, chapter)
}
