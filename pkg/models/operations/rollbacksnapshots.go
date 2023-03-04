package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
	"net/http"
)

type RollbackSnapshotsPathParams struct {
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspaceId"`
}

type RollbackSnapshotsRequest struct {
	PathParams RollbackSnapshotsPathParams
	Request    int64 `request:"mediaType=application/json"`
}

type RollbackSnapshotsResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	Secrets       []shared.Secret
	StatusCode    int
	RawResponse   *http.Response
}
