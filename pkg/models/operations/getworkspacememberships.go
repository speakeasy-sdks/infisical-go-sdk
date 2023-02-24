package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
)

type GetWorkspaceMembershipsPathParams struct {
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspaceId"`
}

type GetWorkspaceMembershipsRequest struct {
	PathParams GetWorkspaceMembershipsPathParams
}

type GetWorkspaceMembershipsResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	Memberships   []shared.Membership
	StatusCode    int
}
