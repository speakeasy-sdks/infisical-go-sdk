package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
	"net/http"
)

type GetWorkspaceKeysPathParams struct {
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspaceId"`
}

type GetWorkspaceKeysRequest struct {
	PathParams GetWorkspaceKeysPathParams
}

type GetWorkspaceKeysResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	Keys          []shared.Key
	StatusCode    int
	RawResponse   *http.Response
}
