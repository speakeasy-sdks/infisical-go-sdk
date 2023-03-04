package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
	"net/http"
)

type GetWorkspaceLogsPathParams struct {
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspaceId"`
}

type GetWorkspaceLogsRequest struct {
	PathParams GetWorkspaceLogsPathParams
}

type GetWorkspaceLogsResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	Logs          []shared.Log
	StatusCode    int
	RawResponse   *http.Response
}
