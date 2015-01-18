package gonfigure_test

import (
	"log"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

const envVar = "TEST_VARIABLE"

var originalValue string

func TestGonfigure(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gonfigure Suite")
}

var _ = BeforeSuite(func() {
	originalValue = os.Getenv(envVar)
})

var _ = AfterSuite(restoreOriginalValue)

func restoreOriginalValue() {
	err := os.Setenv(envVar, originalValue)
	if err != nil {
		log.Print(err)
	}
}
