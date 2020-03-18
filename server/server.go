package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/amsipe/nr-span-example/mysql"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Server struct {
	store *mysql.Store
}

func New(store *mysql.Store) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) HandleMakeSpans() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		users, err := s.store.GetUsers(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		time.Sleep(10 * time.Millisecond)
		s.serviceMethod(ctx)

		time.Sleep(10 * time.Millisecond)
		s.anotherMethod(ctx)

		json.NewEncoder(w).Encode(users)
	}
}

func (s *Server) serviceMethod(ctx context.Context) {
	seg := newrelic.FromContext(ctx).StartSegment("serviceMethod")

	_, _ = s.store.GetUser(ctx, 1)
	seg.End()
}

func (s *Server) anotherMethod(ctx context.Context) {
	seg := newrelic.FromContext(ctx).StartSegment("anotherMethod")

	_, _ = s.store.GetUser(ctx, 2)

	s.serviceMethod(ctx)
	seg.End()
}
