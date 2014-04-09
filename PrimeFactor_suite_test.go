package PrimeFactor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPrimeFactor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PrimeFactor Suite")
}
