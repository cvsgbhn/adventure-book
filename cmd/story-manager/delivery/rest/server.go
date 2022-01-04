package rest

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	srv *http.Server
	ch  *ChapterHandler
	h   Cors
}

type Cors mux.MiddlewareFunc

func NewServer(port string, ch *ChapterHandler) *Server {
	return &Server{
		srv: &http.Server{
			Addr:    ":" + port,
			Handler: nil,
		},
		ch: ch,
		h:  CreateHttpHandler(),
	}
}

func CreateHttpHandler() Cors {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler.ServeHTTP(w, r)
		})
	}
}

func (s *Server) setUsersRoutes(r *mux.Router) {
	users := r.PathPrefix("/next-arc").
		Subrouter()

	users.HandleFunc("/", s.ch.GetChapter).
		Methods(http.MethodGet)
}

func (s *Server) Serve() error {
	r := mux.NewRouter()
	r.Use(mux.MiddlewareFunc(s.h))

	s.setUsersRoutes(r)

	s.srv.Handler = r

	return s.srv.ListenAndServe()
}
