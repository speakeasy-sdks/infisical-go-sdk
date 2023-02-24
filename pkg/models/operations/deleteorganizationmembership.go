package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
)

type DeleteOrganizationMembershipPathParams struct {
	MembershipID   string `pathParam:"style=simple,explode=false,name=membershipId"`
	OrganizationID string `pathParam:"style=simple,explode=false,name=organizationId"`
}

type DeleteOrganizationMembershipRequest struct {
	PathParams DeleteOrganizationMembershipPathParams
}

type DeleteOrganizationMembershipResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	Membership    *shared.Membership
	StatusCode    int
}
