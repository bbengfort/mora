package mora

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/tent/hawk-go"
)

// Resource definitions for RESTful endpoints in the Scribo App.
const (
	NODES = "nodes"
	PINGS = "pings"
)

// ScriboClient connects to and interacts with the Scribo RESTful API.
type ScriboClient struct {
	config      *Configuration    // A pointer to the configuration in the app.
	credentials *hawk.Credentials // A pointer to a HAWK credential struct
	client      *http.Client      // The inner HTTP client for making requests
}

// Get makes a GET http request to the given resource.
// Note that this method must have a pointer receiver because Authenticate
// must check and modify the credentials of the scribo instance.
func (scribo *ScriboClient) Get(resource string) (interface{}, error) {
	endpoint, err := scribo.constructEndpoint(resource)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	scribo.Authenticate(request)

	response, err := scribo.client.Do(request)
	if err != nil {
		return nil, err
	}

	_, err = httputil.DumpResponse(response, true)
	if err != nil {
		return nil, err
	}

	// fmt.Println(string(data))

	return nil, nil
}

// Authenticate a request using the configuration provided.
func (scribo *ScriboClient) Authenticate(request *http.Request) {

	// Construct a set of credentials if they don't already exist.
	// Note that all callers must also have the pointer receiver to modify scribo.
	if scribo.credentials == nil {
		scribo.credentials = new(hawk.Credentials)
		scribo.credentials.ID = scribo.config.Name
		scribo.credentials.Key = scribo.config.ScriboKey
		scribo.credentials.Hash = sha256.New

		fmt.Println("Credentials created!")
	} else {
		fmt.Println("Credentials not created!")
	}

	// Add the credentials to the request
	auth := hawk.NewRequestAuth(request, scribo.credentials, time.Duration(0))
	request.Header.Add("Authorization", auth.RequestHeader())
}

// Construct an endpoint from a resource, should be one of NODES or PINGS,
// but I have refrained from error checking to make things more adaptable.
func (scribo ScriboClient) constructEndpoint(resource string) (string, error) {

	u, err := url.Parse(scribo.config.ScriboURL)
	if err != nil {
		return "", err
	}

	u.Path = resource
	return u.String(), nil
}
