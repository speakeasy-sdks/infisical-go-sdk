# github.com/speakeasy-sdks/infisical-go-sdk

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/infisical-go-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/infisical-go-sdk"
    "github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/infisical-go-sdk/pkg/models/operations"
)

func main() {
    opts := []sdk.SDKOption{
        sdk.WithSecurity(
            shared.Security{
                BearerAuth: shared.SchemeBearerAuth{
                    Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
                },
            },
        ),
    }

    s := sdk.New(opts...)
    
    req := operations.GetWorkspaceKeysRequest{
        PathParams: operations.GetWorkspaceKeysPathParams{
            WorkspaceID: "unde",
        },
    }

    ctx := context.Background()
    res, err := s.Key.GetWorkspaceKeys(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Keys != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations


### Key

* `GetWorkspaceKeys` - Get workspace encrypted key

### Log

* `GetWorkspaceLogs` - Get workspace logs

### Membership

* `DeleteOrganizationMembership` - Delete organization membership
* `DeleteWorkspaceMembership` - Delete workspace membership
* `GetOrganizationMemberships` - Get organization memberships
* `GetWorkspaceMemberships` - Get workspace memberships
* `UpdateOrganizationMembership` - Update organization membership
* `UpdateWorkspaceMembership` - Update workspace membership

### Organization

* `DeleteOrganizationMembership` - Delete organization membership
* `GetOrganizationMemberships` - Get organization memberships
* `GetOrganizationWorkspaces` - Get organization workspaces
* `UpdateOrganizationMembership` - Update organization membership

### Secret

* `CreateSecret` - Create secret
* `DeleteSecrets` - Delete secrets
* `GetSecretVersions` - Get secret versions
* `GetSecrets` - Get secrets
* `RollbackSecretVersions` - Rollback secret versions
* `UpdateSecrets` - Update secrets

### Snapshot

* `GetWorkspaceSnapshots` - Get workspace snapshots
* `RollbackSnapshots` - Rollback workspace secret snapshots

### User

* `MyOrganizations` - Get current user organizations
* `MyUser` - Get current user

### Workspace

* `DeleteWorkspaceMembership` - Delete workspace membership
* `GetOrganizationWorkspaces` - Get organization workspaces
* `GetWorkspaceKeys` - Get workspace encrypted key
* `GetWorkspaceLogs` - Get workspace logs
* `GetWorkspaceMemberships` - Get workspace memberships
* `GetWorkspaceSnapshots` - Get workspace snapshots
* `RollbackSnapshots` - Rollback workspace secret snapshots
* `UpdateWorkspaceMembership` - Update workspace membership
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
