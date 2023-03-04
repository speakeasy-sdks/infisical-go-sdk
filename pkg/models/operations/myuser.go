package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
	"net/http"
)

type MyUserResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	StatusCode    int
	RawResponse   *http.Response
	User          *shared.User
}
