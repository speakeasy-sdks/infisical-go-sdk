package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
)

type UpdateWorkspaceMembershipPathParams struct {
	MembershipID string `pathParam:"style=simple,explode=false,name=membershipId"`
	WorkspaceID  string `pathParam:"style=simple,explode=false,name=workspaceId"`
}

type UpdateWorkspaceMembershipRequest struct {
	PathParams UpdateWorkspaceMembershipPathParams
	Request    shared.Membership `request:"mediaType=application/json"`
}

type UpdateWorkspaceMembershipResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	Membership    *shared.Membership
	StatusCode    int
}
