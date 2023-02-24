package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
)

type DeleteSecretsRequest struct {
	Request interface{} `request:"mediaType=application/json"`
}

type DeleteSecretsResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	Secrets       []shared.Secret
	StatusCode    int
}
