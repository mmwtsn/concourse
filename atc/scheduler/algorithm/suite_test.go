package algorithm_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tedsuo/ifrit"
	"go.opentelemetry.io/otel/exporters/jaeger"

	"github.com/concourse/concourse/v7/atc/db"
	"github.com/concourse/concourse/v7/atc/db/lock"
	"github.com/concourse/concourse/v7/atc/metric"
	"github.com/concourse/concourse/v7/atc/postgresrunner"
	"github.com/concourse/concourse/v7/tracing"
)

var (
	postgresRunner postgresrunner.Runner
	dbProcess      ifrit.Process

	lockFactory lock.LockFactory
	teamFactory db.TeamFactory

	dbConn db.Conn

	exporter *jaeger.Exporter
)

func TestAlgorithm(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Algorithm Suite")
}

var _ = BeforeSuite(func() {
	jaegerURL := os.Getenv("JAEGER_URL")

	if jaegerURL != "" {
		c := tracing.Config{
			Jaeger: tracing.Jaeger{
				Endpoint: jaegerURL + "/api/traces",
				Service:  "algorithm_test",
			},
		}

		err := c.Prepare()
		Expect(err).ToNot(HaveOccurred())
	}

	postgresrunner.InitializeRunnerForGinkgo(&postgresRunner, &dbProcess)
})

var _ = BeforeEach(func() {
	postgresRunner.CreateTestDBFromTemplate()

	dbConn = postgresRunner.OpenConn()
	db.CleanupBaseResourceTypesCache()

	var lockConns [lock.FactoryCount]*sql.DB
	for i := 0; i < lock.FactoryCount; i++ {
		lockConns[i] = postgresRunner.OpenSingleton()
	}
	lockFactory = lock.NewLockFactory(lockConns, metric.LogLockAcquired, metric.LogLockReleased)
	teamFactory = db.NewTeamFactory(dbConn, lockFactory)
})

var _ = AfterEach(func() {
	err := dbConn.Close()
	Expect(err).NotTo(HaveOccurred())

	postgresRunner.DropTestDB()
})

var _ = AfterSuite(func() {
	postgresrunner.FinalizeRunnerForGinkgo(&postgresRunner, &dbProcess)

	if exporter != nil {
		exporter.Shutdown(context.Background())
	}
})
