package buildserver

import (
	"encoding/json"
	"net/http"

	"github.com/concourse/concourse/v7/atc/api/accessor"
	"github.com/concourse/concourse/v7/atc/api/present"
	"github.com/concourse/concourse/v7/atc/db"
)

func (s *Server) GetBuild(build db.BuildForAPI) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.Session("get-build")

		job, _, _ := build.Job()
		access := accessor.GetAccessor(r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(present.Build(build, job, access))
		if err != nil {
			logger.Error("failed-to-encode-build", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
