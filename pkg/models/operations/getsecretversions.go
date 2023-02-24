package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
)

type GetSecretVersionsPathParams struct {
	SecretID string `pathParam:"style=simple,explode=false,name=secretId"`
}

type GetSecretVersionsQueryParams struct {
	Limit  *string `queryParam:"style=form,explode=true,name=limit"`
	Offset *string `queryParam:"style=form,explode=true,name=offset"`
}

type GetSecretVersionsRequest struct {
	PathParams  GetSecretVersionsPathParams
	QueryParams GetSecretVersionsQueryParams
}

type GetSecretVersionsResponse struct {
	ContentType    string
	ErrorResponse  *shared.ErrorResponse
	SecretVersions []shared.SecretVersion
	StatusCode     int
}
