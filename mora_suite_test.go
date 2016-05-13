package mora_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMora(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mora Suite")
}
