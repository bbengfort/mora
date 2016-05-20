package mora

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/tent/hawk-go"
)

// Resource definitions for RESTful endpoints in the Scribo App.
const (
	NODES = "nodes"
	PINGS = "pings"
)

// Request Header Keys and Values
const (
	ContentTypeKey = "Content-Type"
	ContentTypeVal = "application/json;charset=UTF-8"
	XScriboKey     = "X-Scribo-Application"
	XScriboVal     = "Mora-Scio/v%s"
)

// Go 1.6 net/http does not include HTTP Method Constants, so we include them.
const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
)

// ScriboClient connects to and interacts with the Scribo RESTful API.
type ScriboClient struct {
	Credentials *hawk.Credentials // A pointer to a HAWK credential struct
	config      *Configuration    // A pointer to the configuration in the app.
	client      *http.Client      // The inner HTTP client for making requests
}

// Do executes a request with the internal client, ensuring that it is
// authenticated and that any default headers are added to the request.
// Note that this method must have a pointer receiver in order to ensure that
// Authenticate can check and modify the credentials of the scribo instance.
func (scribo *ScriboClient) Do(request *http.Request) (*http.Response, error) {
	// Authenticate the request and add headers
	scribo.Authenticate(request)
	request.Header.Set(XScriboKey, fmt.Sprintf(XScriboVal, Version))

	// Execute the request
	return scribo.client.Do(request)
}

// Get makes a GET http request to the given resource.
// Note that this method must have a pointer receiver because Authenticate
// must check and modify the credentials of the scribo instance.
func (scribo *ScriboClient) Get(resource string, detail ...string) (*http.Response, error) {
	endpoint, err := scribo.Endpoint(resource, detail...)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return scribo.Do(request)
}

// Post makes a POST http request to the given resource with associated data.
// Note that this method must have a pointer receiver because Authenticate
// must check and modify the credentials of the scribo instance.
func (scribo *ScriboClient) Post(data interface{}, resource string, detail ...string) (*http.Response, error) {
	// Construct the endpoint
	endpoint, err := scribo.Endpoint(resource, detail...)
	if err != nil {
		return nil, err
	}

	// Marshal the JSON data into a byte buffer
	buff := new(bytes.Buffer)
	if err = json.NewEncoder(buff).Encode(data); err != nil {
		return nil, err
	}

	// Construct the request, execute it and return the response
	request, err := http.NewRequest(MethodPost, endpoint, buff)
	if err != nil {
		return nil, err
	}

	// Set the Content-Type Header
	request.Header.Set(ContentTypeKey, ContentTypeVal)

	return scribo.Do(request)

}

// Put makes a PUT http request to update a resource/detail with new data.
// Note that this method must have a pointer receiver because Authenticate
// must check and modify the credentials of the scribo instance.
func (scribo *ScriboClient) Put(data interface{}, resource string, detail string) (*http.Response, error) {
	// Construct the endpoint
	endpoint, err := scribo.Endpoint(resource, detail)
	if err != nil {
		return nil, err
	}

	// Marshal the JSON data into a byte buffer
	buff := new(bytes.Buffer)
	if err = json.NewEncoder(buff).Encode(data); err != nil {
		return nil, err
	}

	// Construct the request, execute it and return the response
	request, err := http.NewRequest(MethodPut, endpoint, buff)
	if err != nil {
		return nil, err
	}

	// Set the Content-Type Header
	request.Header.Set(ContentTypeKey, ContentTypeVal)

	return scribo.Do(request)
}

// Delete makes a DELETE http request to delete a resource/detail.
// Note that this method must have a pointer receiver because Authenticate
// must check and modify the credentials of the scribo instance.
func (scribo *ScriboClient) Delete(resource string, detail string) (*http.Response, error) {
	// Construct the endpoint
	endpoint, err := scribo.Endpoint(resource, detail)
	if err != nil {
		return nil, err
	}

	// Construct the request, execute it and return the response
	request, err := http.NewRequest(MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return scribo.Do(request)
}

// Authenticate a request using the configuration provided.
func (scribo *ScriboClient) Authenticate(request *http.Request) {

	// Construct a set of credentials if they don't already exist.
	// Note that all callers must also have the pointer receiver to modify scribo.
	if scribo.Credentials == nil {
		scribo.Credentials = new(hawk.Credentials)
		scribo.Credentials.ID = scribo.config.Name
		scribo.Credentials.Key = scribo.config.ScriboKey
		scribo.Credentials.Hash = sha256.New
	}

	// Add the credentials to the request
	auth := hawk.NewRequestAuth(request, scribo.Credentials, time.Duration(0))
	request.Header.Add("Authorization", auth.RequestHeader())
}

// Endpoint constructs a URL from a resource, e.g. NODES or PINGS, and can
// optionally pass a detail string to get a specific element from the resource.
// TODO: this was exported for testing, but would rather have it be internal.
func (scribo ScriboClient) Endpoint(resource string, detail ...string) (string, error) {
	var scriboPath string

	if len(detail) > 0 {
		parts := append([]string{resource}, detail...)
		scriboPath = path.Join(parts...)
	} else {
		scriboPath = resource
	}

	u, err := url.Parse(scribo.config.ScriboURL)
	if err != nil {
		return "", err
	}

	u.Path = scriboPath
	return u.String(), nil
}
