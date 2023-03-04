package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateOrganizationMembershipPathParams struct {
	MembershipID   string `pathParam:"style=simple,explode=false,name=membershipId"`
	OrganizationID string `pathParam:"style=simple,explode=false,name=organizationId"`
}

type UpdateOrganizationMembershipRequest struct {
	PathParams UpdateOrganizationMembershipPathParams
	Request    shared.Membership `request:"mediaType=application/json"`
}

type UpdateOrganizationMembershipResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	Membership    *shared.Membership
	StatusCode    int
	RawResponse   *http.Response
}
