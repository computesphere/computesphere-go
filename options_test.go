package computesphereapi

import (
	"context"
	"strings"
	"testing"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func TestWithSkipRedeploy(t *testing.T) {
	id := openapi_types.UUID{1}
	req, err := NewDeleteDeploymentEnvVarRequest("https://api.example.test/v2", id, "KEY")
	if err != nil {
		t.Fatal(err)
	}
	if err := WithSkipRedeploy()(context.Background(), req); err != nil {
		t.Fatal(err)
	}
	if got := req.URL.Query().Get("skip_redeploy"); got != "true" {
		t.Fatalf("skip_redeploy = %q, want true (url %s)", got, req.URL)
	}
	if !strings.HasSuffix(req.URL.Path, "/deployments/"+id.String()+"/env-vars/KEY") {
		t.Fatalf("path = %s", req.URL.Path)
	}
}
