# computesphere-go

The official Go SDK for the [ComputeSphere](https://computesphere.com) API.

## Install

```bash
go get github.com/computesphere/computesphere-go
```

## Authentication

The client authenticates with a **ComputeSphere access token** — either an
**API token** (create one in the [console](https://console.computesphere.com)
under **Settings → API Tokens**) or an **OAuth access token**. Pass it however
you obtain it; the client treats it as an opaque bearer credential.

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	computesphere "github.com/computesphere/computesphere-go"
)

func main() {
	token := os.Getenv("COMPUTESPHERE_API_TOKEN")
	accountID := os.Getenv("COMPUTESPHERE_ACCOUNT_ID")
	bearer := func(_ context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	}

	client, err := computesphere.NewClient(
		"https://api.computesphere.com/v2",
		computesphere.WithRequestEditorFn(bearer),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.ListRegionsWithResponse(context.Background(), &computesphere.ListRegionsParams{
		XAccountId: accountID,
	})
	if err != nil {
		log.Fatal(err)
	}
	if resp.JSON200 == nil {
		log.Fatalf("HTTP %d", resp.StatusCode())
	}
	for _, r := range resp.JSON200.Items {
		fmt.Println(r.Name, r.DisplayName)
	}
}
```

## Documentation

Full API reference: <https://docs.computesphere.com>.

## License

[Apache-2.0](./LICENSE). ComputeSphere is a trademark of ComputeSphere LLC.
