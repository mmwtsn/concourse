package worker

import (
	"context"

	"github.com/concourse/concourse/v7/tsa"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . TSAClient
type TSAClient interface {
	Register(context.Context, tsa.RegisterOptions) error

	Land(context.Context) error
	Retire(context.Context) error
	Delete(context.Context) error

	ReportContainers(context.Context, []string) error
	ContainersToDestroy(context.Context) ([]string, error)

	ReportVolumes(context.Context, []string) error
	VolumesToDestroy(context.Context) ([]string, error)
}
