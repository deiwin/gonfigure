package gonfigure_test

import (
	"os"

	// "github.com/deiwin/gonfigure"

	"github.com/deiwin/gonfigure"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Go-n-figure", func() {
	BeforeEach(func() {
		err := os.Unsetenv(envVar)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("EnvProperty", func() {
		var (
			envProp  gonfigure.Property
			fallback = "the fallback value"
		)

		BeforeEach(func() {
			envProp = gonfigure.NewEnvProperty(envVar, fallback)
		})

		It("should return the fallback value", func() {
			val := envProp.Value()
			Expect(val).To(Equal(fallback))
		})

		Context("with environment variable set", func() {
			var currentValue = "something different"
			BeforeEach(func() {
				err := os.Setenv(envVar, currentValue)
				Expect(err).NotTo(HaveOccurred())
			})

			It("should return the set environment variable", func() {
				val := envProp.Value()
				Expect(val).To(Equal(currentValue))
			})
		})
	})

	Describe("RequiredEnvProperty", func() {
		var (
			envProp gonfigure.Property
		)

		BeforeEach(func() {
			envProp = gonfigure.NewRequiredEnvProperty(envVar)
		})

		It("should panic", func() {
			Expect(func() {
				_ = envProp.Value()
			}).To(Panic())
		})

		Context("with environment variable set", func() {
			var currentValue = "something differenter"
			BeforeEach(func() {
				err := os.Setenv(envVar, currentValue)
				Expect(err).NotTo(HaveOccurred())
			})

			It("should return the set environment variable", func() {
				val := envProp.Value()
				Expect(val).To(Equal(currentValue))
			})
		})
	})
})
