package config

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	apiclient "github.com/pydio/pydio-sdk-go/client"
)

var (
	ApiResourcePath  = "/pydio/api/v2"
	OidcResourcePath = "/auth/dex"

	grantType = "password"
	scope     = "email profile pydio"

	store = NewTokenStore()
)

// GetPreparedApiClient connects to the Pydio Cells server defined by this config.
// Also returns a context to be used in subsequent requests.
func GetPreparedApiClient(sdkConfig *SdkConfig) (*apiclient.PydioAPIV2, context.Context, error) {

	transport := httptransport.New(sdkConfig.Url, ApiResourcePath, []string{sdkConfig.Protocol})

	basicAuth := httptransport.BasicAuth(sdkConfig.User, sdkConfig.Password)

	transport.DefaultAuthentication = basicAuth

	client := apiclient.New(transport, strfmt.Default)

	return client, context.Background(), nil
}

func GetHttpClient(sdkConfig *SdkConfig) *http.Client {

	if sdkConfig.SkipVerify {
		log.Println("[WARNING] Using SkipVerify for ignoring SSL certificate issues!!")
		return &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		}}
	}
	return http.DefaultClient
}
