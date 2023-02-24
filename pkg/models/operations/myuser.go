package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
)

type MyUserResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	StatusCode    int
	User          *shared.User
}
