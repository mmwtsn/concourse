package artifactserver

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/concourse/concourse/v7/atc/api/present"
	"github.com/concourse/concourse/v7/atc/compression"
	"github.com/concourse/concourse/v7/atc/db"
	"github.com/concourse/concourse/v7/atc/worker"
)

func (s *Server) CreateArtifact(team db.Team) http.Handler {
	hLog := s.logger.Session("create-artifact")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		w.Header().Set("Content-Type", "application/json")

		// TODO: can probably check if fly sent us an etag header
		// which we can lookup in the checksum field
		// that way we don't have to create another volume.

		workerSpec := worker.Spec{
			TeamID:   team.ID(),
			Platform: r.FormValue("platform"),
			Tags:     r.Form["tags"],
		}

		volume, artifact, err := s.workerPool.CreateVolumeForArtifact(ctx, workerSpec)
		if err != nil {
			hLog.Error("failed-to-create-volume", err)
			http.Error(w, fmt.Sprintf("%v",err), http.StatusInternalServerError)
			return
		}

		err = volume.StreamIn(r.Context(), "/", compression.NewGzipCompression(), r.Body)
		if err != nil {
			hLog.Error("failed-to-stream-volume-contents", err)
			http.Error(w, fmt.Sprintf("%v",err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(present.WorkerArtifact(artifact))
	})
}
