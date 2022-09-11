package flaghelpers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/concourse/concourse/v7/atc"
	"github.com/concourse/concourse/v7/fly/commands/internal/flaghelpers"
)

var _ = Describe("PipelineFlag", func() {
	Describe("UnmarshalFlag", func() {

		var pipelineFlag *flaghelpers.PipelineFlag

		BeforeEach(func() {
			pipelineFlag = &flaghelpers.PipelineFlag{}
		})

		for _, tt := range []struct {
			desc         string
			flag         string
			name         string
			instanceVars atc.InstanceVars
		}{
			{
				desc: "name",
				flag: "some-pipeline",
				name: "some-pipeline",
			},
			{
				desc:         "instance var",
				flag:         "some-pipeline/branch:master",
				name:         "some-pipeline",
				instanceVars: atc.InstanceVars{"branch": "master"},
			},
		} {
			tt := tt
			It(tt.desc, func() {
				err := pipelineFlag.UnmarshalFlag(tt.flag)
				Expect(err).ToNot(HaveOccurred())
				Expect(pipelineFlag.Name).To(Equal(tt.name))
				Expect(pipelineFlag.InstanceVars).To(Equal(tt.instanceVars))
			})
		}
	})
})
