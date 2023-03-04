package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
	"net/http"
)

type MyOrganizationsResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	Organizations []shared.Organization
	StatusCode    int
	RawResponse   *http.Response
}
