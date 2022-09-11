package artifactserver

import (
	"context"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/v7/atc/db"
	"github.com/concourse/concourse/v7/atc/runtime"
	"github.com/concourse/concourse/v7/atc/worker"
)

type Pool interface {
	LocateVolume(ctx context.Context, teamID int, handle string) (runtime.Volume, runtime.Worker, bool, error)
	CreateVolumeForArtifact(ctx context.Context, spec worker.Spec) (runtime.Volume, db.WorkerArtifact, error)
}

type Server struct {
	logger     lager.Logger
	workerPool Pool
}

func NewServer(
	logger lager.Logger,
	workerPool Pool,
) *Server {
	return &Server{
		logger:     logger,
		workerPool: workerPool,
	}
}
