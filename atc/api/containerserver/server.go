package containerserver

import (
	"context"
	"time"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/v7/atc/db"
	"github.com/concourse/concourse/v7/atc/gc"
	"github.com/concourse/concourse/v7/atc/runtime"
)

type Pool interface {
	LocateContainer(ctx context.Context, teamID int, handle string) (runtime.Container, runtime.Worker, bool, error)
}

type Server struct {
	logger lager.Logger

	workerPool              Pool
	interceptTimeoutFactory InterceptTimeoutFactory
	interceptUpdateInterval time.Duration
	containerRepository     db.ContainerRepository
	destroyer               gc.Destroyer
	clock                   clock.Clock
}

func NewServer(
	logger lager.Logger,
	workerPool Pool,
	interceptTimeoutFactory InterceptTimeoutFactory,
	interceptUpdateInterval time.Duration,
	containerRepository db.ContainerRepository,
	destroyer gc.Destroyer,
	clock clock.Clock,
) *Server {
	return &Server{
		logger:                  logger,
		workerPool:              workerPool,
		interceptTimeoutFactory: interceptTimeoutFactory,
		interceptUpdateInterval: interceptUpdateInterval,
		containerRepository:     containerRepository,
		destroyer:               destroyer,
		clock:                   clock,
	}
}
