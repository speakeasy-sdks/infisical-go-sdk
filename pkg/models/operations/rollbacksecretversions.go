package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
)

type RollbackSecretVersionsPathParams struct {
	SecretID string `pathParam:"style=simple,explode=false,name=secretId"`
}

type RollbackSecretVersionsRequest struct {
	PathParams RollbackSecretVersionsPathParams
	Request    int64 `request:"mediaType=application/json"`
}

type RollbackSecretVersionsResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	Secret        *shared.Secret
	StatusCode    int
}
