package ops_test

import (
	"testing"
	"time"

	"github.com/concourse/concourse/v7/integration/internal/flytest"
	"github.com/stretchr/testify/require"
)

func setupUpgradeDowngrade(t *testing.T, fly flytest.Cmd) {
	t.Run("set up test pipeline", func(t *testing.T) {
		fly.Run(t, "set-pipeline", "-p", "test", "-c", "pipelines/smoke-pipeline.yml", "-n")
		fly.Run(t, "unpause-pipeline", "-p", "test")

		fly.Run(t, "trigger-job", "-j", "test/say-hello", "-w")
	})
}

func verifyUpgradeDowngrade(t *testing.T, fly flytest.Cmd) {
	t.Run("pipeline and build still exists", func(t *testing.T) {
		fly.Run(t, "get-pipeline", "-p", "test")

		out := fly.Output(t, "watch", "-j", "test/say-hello", "--ignore-event-parsing-errors")
		require.Contains(t, out, "Hello, world!")
	})

	t.Run("can still run pipeline builds", func(t *testing.T) {
		fly.Run(t, "trigger-job", "-j", "test/say-hello", "-w")
	})

	t.Run("can still run checks", func(t *testing.T) {
		fly.Run(t, "check-resource", "-r", "test/mockery")
	})

	t.Run("can still run one-off builds", func(t *testing.T) {
		out := fly.Output(t, "execute", "-c", "../tasks/hello.yml")
		require.Contains(t, out, "hello")
	})
}

func waitForVolumesGC(t *testing.T, fly flytest.Cmd) {
	require.Eventually(t, func() bool {
		volumes := fly.Table(t, "volumes")
		return len(volumes) == 0
	}, 2*time.Minute, 5*time.Second)
}
