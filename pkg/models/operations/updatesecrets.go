package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
)

type UpdateSecretsRequestBody struct {
	Secrets []shared.Secret `json:"secrets,omitempty"`
}

type UpdateSecretsRequest struct {
	Request UpdateSecretsRequestBody `request:"mediaType=application/json"`
}

type UpdateSecretsResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	Secrets       []shared.Secret
	StatusCode    int
}
