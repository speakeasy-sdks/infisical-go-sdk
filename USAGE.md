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