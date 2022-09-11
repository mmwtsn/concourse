package gc

import (
	"context"

	"code.cloudfoundry.org/lager/lagerctx"
	"github.com/concourse/concourse/v7/atc/db"
)

type checksCollector struct {
	lifecycle db.CheckLifecycle
}

func NewChecksCollector(lifecycle db.CheckLifecycle) *checksCollector {
	return &checksCollector{
		lifecycle: lifecycle,
	}
}

func (c *checksCollector) Run(ctx context.Context) error {
	logger := lagerctx.FromContext(ctx).Session("check-collector")

	logger.Debug("start")
	defer logger.Debug("done")

	err := c.lifecycle.DeleteCompletedChecks(logger)
	if err != nil {
		logger.Error("failed-to-delete-completed-checks", err)
		return err
	}

	return nil
}
