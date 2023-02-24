package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
)

type GetWorkspaceSnapshotsPathParams struct {
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspaceId"`
}

type GetWorkspaceSnapshotsRequest struct {
	PathParams GetWorkspaceSnapshotsPathParams
}

type GetWorkspaceSnapshotsResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	Snapshots     []interface{}
	StatusCode    int
}
