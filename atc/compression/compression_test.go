package compression_test

import (
	"github.com/concourse/concourse/v7/atc/compression"
	"github.com/concourse/concourse/v7/worker/baggageclaim"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Compression", func() {
	var (
		comp compression.Compression
	)

	Describe("Gzip", func() {
		BeforeEach(func() {
			comp = compression.NewGzipCompression()
		})

		It("returns gzip", func() {
			Expect(comp.Encoding()).To(Equal(baggageclaim.GzipEncoding))
		})
	})

	Describe("Zstd", func() {
		BeforeEach(func() {
			comp = compression.NewZstdCompression()
		})

		It("returns zstd", func() {
			Expect(comp.Encoding()).To(Equal(baggageclaim.ZstdEncoding))
		})
	})
})
