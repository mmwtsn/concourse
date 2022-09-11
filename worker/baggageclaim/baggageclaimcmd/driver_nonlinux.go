//go:build !linux
// +build !linux

package baggageclaimcmd

import (
	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/v7/worker/baggageclaim/volume"
	"github.com/concourse/concourse/v7/worker/baggageclaim/volume/driver"
)

func (cmd *BaggageclaimCommand) driver(logger lager.Logger) (volume.Driver, error) {
	return &driver.NaiveDriver{}, nil
}
