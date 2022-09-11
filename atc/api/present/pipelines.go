package present

import (
	"github.com/concourse/concourse/v7/atc"
	"github.com/concourse/concourse/v7/atc/db"
)

func Pipelines(savedPipelines []db.Pipeline) []atc.Pipeline {
	pipelines := make([]atc.Pipeline, len(savedPipelines))

	for i := range savedPipelines {
		pipelines[i] = Pipeline(savedPipelines[i])
	}

	return pipelines
}
