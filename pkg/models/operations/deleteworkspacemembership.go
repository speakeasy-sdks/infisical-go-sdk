package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
	"net/http"
)

type DeleteWorkspaceMembershipPathParams struct {
	MembershipID string `pathParam:"style=simple,explode=false,name=membershipId"`
	WorkspaceID  string `pathParam:"style=simple,explode=false,name=workspaceId"`
}

type DeleteWorkspaceMembershipRequest struct {
	PathParams DeleteWorkspaceMembershipPathParams
}

type DeleteWorkspaceMembershipResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	Membership    *shared.Membership
	StatusCode    int
	RawResponse   *http.Response
}
