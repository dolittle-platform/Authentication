package nonces_test

import (
	"dolittle.io/pascal/sessions/nonces"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

var _ = Describe("Generator", func() {
	When("generating a nonce with 18 bytes", func() {
		var (
			generated nonces.Nonce
			err       error
		)
		BeforeEach(func() {
			logger := zap.NewNop()
			config := &configuration{size: 18}
			generator := nonces.NewGenerator(config, logger)

			generated, err = generator.Generate()
		})
		It("should generate a string of length 24", func() {
			Expect(generated).To(HaveLen(24))
		})
		It("should generate a string with only valid URL characters", func() {
			Expect(string(generated)).To(OnlyIncludeValidURLCharacters())
		})
		It("should not fail", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})

	When("generating a nonce with 24 bytes", func() {
		var (
			generated nonces.Nonce
			err       error
		)
		BeforeEach(func() {
			logger := zap.NewNop()
			config := &configuration{size: 24}
			generator := nonces.NewGenerator(config, logger)

			generated, err = generator.Generate()
		})
		It("should not fail", func() {
			Expect(err).ToNot(HaveOccurred())
		})
		It("should generate a string of length 32", func() {
			Expect(generated).To(HaveLen(32))
		})
		It("should generate a string with only valid URL characters", func() {
			Expect(string(generated)).To(OnlyIncludeValidURLCharacters())
		})
	})

	When("generating three nonces", func() {
		var (
			generated1 nonces.Nonce
			generated2 nonces.Nonce
			generated3 nonces.Nonce
			err1       error
			err2       error
			err3       error
		)
		BeforeEach(func() {
			logger := zap.NewNop()
			config := &configuration{size: 24}
			generator := nonces.NewGenerator(config, logger)

			generated1, err1 = generator.Generate()
			generated2, err2 = generator.Generate()
			generated3, err3 = generator.Generate()
		})
		It("should not fail", func() {
			Expect(err1).ToNot(HaveOccurred())
			Expect(err2).ToNot(HaveOccurred())
			Expect(err3).ToNot(HaveOccurred())
		})
		It("should generate three unique strings", func() {
			Expect(generated1).NotTo(Equal(generated2))
			Expect(generated1).NotTo(Equal(generated3))
			Expect(generated2).NotTo(Equal(generated3))
		})
	})
})
