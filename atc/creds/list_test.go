package creds_test

import (
	"github.com/concourse/concourse/v7/atc/creds"
	"github.com/concourse/concourse/v7/vars"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Evaluate", func() {
	It("interpolates the full list", func() {
		list, err := creds.NewList(
			vars.StaticVariables{"list": []string{"foo", "bar"}},
			"((list))",
		).Evaluate()
		Expect(err).ToNot(HaveOccurred())
		Expect(list).To(Equal([]interface{}{"foo", "bar"}))
	})

	It("interpolates within a list", func() {
		list, err := creds.NewList(
			vars.StaticVariables{"element": "blah"},
			[]interface{}{"abc", "((element))"},
		).Evaluate()
		Expect(err).ToNot(HaveOccurred())
		Expect(list).To(Equal([]interface{}{"abc", "blah"}))
	})
})
