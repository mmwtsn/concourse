package worker_test

import (
	"context"
	"database/sql"
	"testing"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagertest"
	"github.com/concourse/concourse/v7/atc/db"
	"github.com/concourse/concourse/v7/atc/db/lock"
	"github.com/concourse/concourse/v7/atc/postgresrunner"
	"github.com/concourse/concourse/v7/atc/worker/workertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	postgresRunner postgresrunner.Runner
	dbConn         db.Conn
	lockFactory    lock.LockFactory
)

var logger = lagertest.NewTestLogger("dummy")
var ctx = context.Background()

var _ = postgresrunner.GinkgoRunner(&postgresRunner)

var _ = BeforeEach(func() {
	postgresRunner.CreateTestDBFromTemplate()

	dbConn = postgresRunner.OpenConn()
	db.CleanupBaseResourceTypesCache()

	ignore := func(logger lager.Logger, id lock.LockID) {}
	var lockConns [lock.FactoryCount]*sql.DB
	for i := 0; i < lock.FactoryCount; i++ {
		lockConns[i] = postgresRunner.OpenSingleton()
	}
	lockFactory = lock.NewLockFactory(lockConns, ignore, ignore)
})

var _ = AfterEach(func() {
	Expect(dbConn.Close()).To(Succeed())
	postgresRunner.DropTestDB()
})

func TestWorker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Worker Suite")
}

func Setup(setup ...workertest.SetupFunc) *workertest.Scenario {
	return workertest.Setup(dbConn, lockFactory, setup...)
}

var Test = It
var FTest = FIt
var XTest = XIt
