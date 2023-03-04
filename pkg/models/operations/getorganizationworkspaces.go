package operations

import (
	"github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
	"net/http"
)

type GetOrganizationWorkspacesPathParams struct {
	OrganizationID string `pathParam:"style=simple,explode=false,name=organizationId"`
}

type GetOrganizationWorkspacesRequest struct {
	PathParams GetOrganizationWorkspacesPathParams
}

type GetOrganizationWorkspacesResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	StatusCode    int
	RawResponse   *http.Response
	Workspaces    []shared.Workspace
}
