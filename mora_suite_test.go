package mora_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMora(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mora Suite")
}

// Test Environment Variables
const (
	TEST_MORA_NODE_NAME  = "tester"
	TEST_MORA_SCRIBO_URL = "http://localhost:5080/"
	TEST_MORA_SCRIBO_KEY = "JEVxHZKVwhsa6asUY1DCiaGOhJ11sMUHgHeO+50sZgQ="
)

// Set the test environment before the suite
var _ = BeforeSuite(func() {
	os.Setenv("MORA_NODE_NAME", TEST_MORA_NODE_NAME)
	os.Setenv("MORA_SCRIBO_URL", TEST_MORA_SCRIBO_URL)
	os.Setenv("MORA_SCRIBO_KEY", TEST_MORA_SCRIBO_KEY)
})
