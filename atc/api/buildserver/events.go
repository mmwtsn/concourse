package buildserver

import (
	"net/http"

	"github.com/concourse/concourse/v7/atc/db"
)

func (s *Server) BuildEvents(build db.BuildForAPI) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		streamDone := make(chan struct{})

		go func() {
			defer close(streamDone)

			s.eventHandlerFactory(s.logger, build).ServeHTTP(w, r)
		}()

		<-streamDone
	})
}
