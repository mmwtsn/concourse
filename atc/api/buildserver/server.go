package buildserver

import (
	"net/http"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/v7/atc/api/auth"
	"github.com/concourse/concourse/v7/atc/db"
)

type EventHandlerFactory func(lager.Logger, db.BuildForAPI) http.Handler

type Server struct {
	logger lager.Logger

	externalURL string

	teamFactory         db.TeamFactory
	buildFactory        db.BuildFactory
	eventHandlerFactory EventHandlerFactory
	rejector            auth.Rejector
}

func NewServer(
	logger lager.Logger,
	externalURL string,
	teamFactory db.TeamFactory,
	buildFactory db.BuildFactory,
	eventHandlerFactory EventHandlerFactory,
) *Server {
	return &Server{
		logger: logger,

		externalURL: externalURL,

		teamFactory:         teamFactory,
		buildFactory:        buildFactory,
		eventHandlerFactory: eventHandlerFactory,

		rejector: auth.UnauthorizedRejector{},
	}
}
