package pauser_test

import (
	"context"
	"testing"

	"github.com/concourse/concourse/v7/atc/component"
	"github.com/concourse/concourse/v7/atc/db/dbfakes"
	"github.com/concourse/concourse/v7/atc/pauser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPauser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pauser Suite")
}

var _ = Describe("PipelinePauser", func() {
	var pauseComp component.Runnable
	var fakePipelinePauser *dbfakes.FakePipelinePauser

	BeforeEach(func() {
		fakePipelinePauser = new(dbfakes.FakePipelinePauser)
	})

	Describe("Run", func() {
		It("tells the pipeline pauser to pause pipelines older than 10 days", func() {
			pauseComp = pauser.NewPipelinePauser(fakePipelinePauser, 10)
			err := pauseComp.Run(context.TODO())
			Expect(err).NotTo(HaveOccurred())

			Expect(fakePipelinePauser.PausePipelinesCallCount()).To(Equal(1))
			_, givenDays := fakePipelinePauser.PausePipelinesArgsForCall(0)
			Expect(givenDays).To(Equal(10))
		})
		It("it short circuts if days is zero", func() {
			pauseComp = pauser.NewPipelinePauser(fakePipelinePauser, 0)
			err := pauseComp.Run(context.TODO())
			Expect(err).NotTo(HaveOccurred())

			Expect(fakePipelinePauser.PausePipelinesCallCount()).To(Equal(0), "should not call the db.PipelinePauser")
		})
	})
})
