package jobserver

import (
	"net/http"

	"github.com/concourse/concourse/v7/atc/api/accessor"
	"github.com/concourse/concourse/v7/atc/db"
	"github.com/tedsuo/rata"
)

func (s *Server) PauseJob(pipeline db.Pipeline) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.Session("pause-job")
		jobName := rata.Param(r, "job_name")

		job, found, err := pipeline.Job(jobName)
		if err != nil {
			logger.Error("failed-to-get-job", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		acc := accessor.GetAccessor(r)
		user := acc.UserInfo().DisplayUserId

		err = job.Pause(user)
		if err != nil {
			logger.Error("failed-to-pause-job", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
