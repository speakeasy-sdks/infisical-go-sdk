package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
)

type GetSecretsRequestBody struct {
	Context     *string `json:"context,omitempty"`
	Environment *string `json:"environment,omitempty"`
	WorkspaceID *string `json:"workspaceId,omitempty"`
}

type GetSecretsRequest struct {
	Request GetSecretsRequestBody `request:"mediaType=application/json"`
}

type GetSecretsResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	StatusCode    int
}
