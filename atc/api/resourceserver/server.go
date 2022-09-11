package resourceserver

import (
	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/v7/atc/creds"
	"github.com/concourse/concourse/v7/atc/db"
)

type Server struct {
	logger                lager.Logger
	secretManager         creds.Secrets
	varSourcePool         creds.VarSourcePool
	checkFactory          db.CheckFactory
	resourceFactory       db.ResourceFactory
	resourceConfigFactory db.ResourceConfigFactory
}

func NewServer(
	logger lager.Logger,
	secretManager creds.Secrets,
	varSourcePool creds.VarSourcePool,
	checkFactory db.CheckFactory,
	resourceFactory db.ResourceFactory,
	resourceConfigFactory db.ResourceConfigFactory,
) *Server {
	return &Server{
		logger:                logger,
		secretManager:         secretManager,
		varSourcePool:         varSourcePool,
		checkFactory:          checkFactory,
		resourceFactory:       resourceFactory,
		resourceConfigFactory: resourceConfigFactory,
	}
}
