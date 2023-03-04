package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateSecretRequestBody struct {
	Environment *string         `json:"environment,omitempty"`
	Secrets     []shared.Secret `json:"secrets,omitempty"`
	WorkspaceID *string         `json:"workspaceId,omitempty"`
}

type CreateSecretRequest struct {
	Request CreateSecretRequestBody `request:"mediaType=application/json"`
}

type CreateSecretResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	Secrets       []shared.Secret
	StatusCode    int
	RawResponse   *http.Response
}
