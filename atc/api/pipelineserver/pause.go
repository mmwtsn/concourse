package pipelineserver

import (
	"net/http"

	"github.com/concourse/concourse/v7/atc/api/accessor"
	"github.com/concourse/concourse/v7/atc/db"
)

func (s *Server) PausePipeline(pipelineDB db.Pipeline) http.Handler {
	logger := s.logger.Session("pause-pipeline")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		acc := accessor.GetAccessor(r)
		user := acc.UserInfo().DisplayUserId

		err := pipelineDB.Pause(user)
		if err != nil {
			logger.Error("failed-to-pause-pipeline", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
