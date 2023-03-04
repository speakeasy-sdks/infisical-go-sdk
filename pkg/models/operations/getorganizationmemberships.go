package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
	"net/http"
)

type GetOrganizationMembershipsPathParams struct {
	OrganizationID string `pathParam:"style=simple,explode=false,name=organizationId"`
}

type GetOrganizationMembershipsRequest struct {
	PathParams GetOrganizationMembershipsPathParams
}

type GetOrganizationMembershipsResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	Memberships   []shared.Membership
	StatusCode    int
	RawResponse   *http.Response
}
