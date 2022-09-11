package volume

import (
	"io"

	"github.com/concourse/concourse/v7/worker/baggageclaim/uidgid"
)

//go:generate counterfeiter . Streamer

type Streamer interface {
	In(io.Reader, string, bool) (bool, error)
	Out(io.Writer, string, bool) error
}

type tarZstdStreamer struct {
	namespacer uidgid.Namespacer
}

type tarGzipStreamer struct {
	namespacer uidgid.Namespacer
}
