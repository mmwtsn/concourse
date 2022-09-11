package runtimetest

import (
	"context"
	"io"

	"github.com/concourse/concourse/v7/atc/compression"
)

type Artifact struct {
	Content VolumeContent
}

func (a Artifact) Handle() string {
	return ""
}

func (a Artifact) Source() string {
	return ""
}

func (a Artifact) StreamOut(ctx context.Context, path string, compression compression.Compression) (io.ReadCloser, error) {
	return a.Content.StreamOut(ctx, path, compression.Encoding())
}
