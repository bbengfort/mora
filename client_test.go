package mora_test

import (
	"crypto/sha256"
	"net/http"
	"net/http/httputil"

	. "github.com/bbengfort/mora"
	"github.com/tent/hawk-go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Client", func() {

	var (
		server *ghttp.Server
		sonar  *Sonar
		err    error
	)

	BeforeEach(func() {
		sonar, err = New()
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("should authenticate a request with credentials", func() {
		request, err := http.NewRequest(http.MethodGet, sonar.Config.ScriboURL, nil)
		Ω(err).ShouldNot(HaveOccurred())

		Ω(request.Header.Get("Authorization")).Should(Equal(""))
		sonar.Scribo.Authenticate(request)
		Ω(request.Header.Get("Authorization")).ShouldNot(Equal(""))
	})

	It("should only create credentials from config once", func() {
		Ω(sonar.Scribo.Credentials).Should(BeNil(), "Credentials created already?!")

		request, err := http.NewRequest(http.MethodGet, sonar.Config.ScriboURL, nil)
		Ω(err).ShouldNot(HaveOccurred())
		sonar.Scribo.Authenticate(request)

		Ω(sonar.Scribo.Credentials).ShouldNot(BeNil(), "Credentials not stored on client.")
	})

	It("should create an endpoint from a resource", func() {
		endpoint, err := sonar.Scribo.Endpoint(NODES)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(endpoint).Should(Equal(TEST_MORA_SCRIBO_URL + "nodes"))
	})

	It("should create an endpoint from a resource and a detail", func() {
		endpoint, err := sonar.Scribo.Endpoint(NODES, "1")
		Ω(err).ShouldNot(HaveOccurred())
		Ω(endpoint).Should(Equal(TEST_MORA_SCRIBO_URL + "nodes/1"))
	})

	It("should create an endpoint from arbitrarily long detail endpoints", func() {
		endpoint, err := sonar.Scribo.Endpoint(NODES, "1", "set-password")
		Ω(err).ShouldNot(HaveOccurred())
		Ω(endpoint).Should(Equal(TEST_MORA_SCRIBO_URL + "nodes/1/set-password"))
	})

	Describe("client requests to a test server", func() {

		BeforeEach(func() {
			server = ghttp.NewServer()
			server.AllowUnhandledRequests = false
			server.Writer = GinkgoWriter

			sonar.Config.ScriboURL = server.URL()
		})

		AfterEach(func() {
			server.Close()
		})

		Describe("fetching a node list", func() {

			BeforeEach(func() {

				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/nodes"),
						ghttp.RespondWith(204, ""),
						ghttp.VerifyContentType("application/json"),
						VerifyAuth(TEST_MORA_SCRIBO_KEY),
						ghttp.RespondWith(200, `[{"id":3,"name":"dropbox","address":"108.160.172.238","dns":"dropbox.com","created":"2016-05-13T00:11:05.043484-04:00","updated":"2016-05-13T00:11:05.043484-04:00"},{"id":2,"name":"github","address":"192.30.252.122","dns":"github.com","created":"2016-05-13T00:10:20.734114-04:00","updated":"2016-05-13T00:10:20.734114-04:00"},{"id":1,"name":"apollo","address":"108.51.64.223","dns":"bryant.bengfort.com","created":"2016-05-12T23:38:13.930893-04:00","updated":"2016-05-12T23:38:13.930893-04:00"}]`),
					),
				)
			})

			It("should be able to fetch a node list", func() {
				Skip("Test server is not responding correctly.")
				response, err := sonar.Scribo.Get(NODES)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(response).ShouldNot(BeNil(), "No response was returned from nodes.")

				data, err := httputil.DumpResponse(response, true)
				Ω(err).ShouldNot(HaveOccurred())
				GinkgoWriter.Write(data)

				Ω(response.StatusCode).Should(Equal(200))
			})

		})
	})

})

// Test server handlers for mocking the Scribo app
// A closure to wrap a handler with an authentication stub
func VerifyAuth(key string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// Get the authentication from the request
		auth, err := hawk.NewAuthFromRequest(r, func(c *hawk.Credentials) error {
			c.Key = key
			c.Hash = sha256.New
			return nil
		}, nil)

		Ω(err).ShouldNot(HaveOccurred())
		Ω(auth.Valid()).ShouldNot(HaveOccurred(), "Authentication failed")
	}

}
