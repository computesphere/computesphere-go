package computesphereapi

import (
	"context"
	"net/http"
)

// WithSkipRedeploy adds skip_redeploy=true to a request's query, for the
// operations that accept it as a query parameter, such as
// DeleteDeploymentEnvVar:
//
//	client.DeleteDeploymentEnvVar(ctx, id, key, api.WithSkipRedeploy())
func WithSkipRedeploy() RequestEditorFn {
	return func(_ context.Context, req *http.Request) error {
		q := req.URL.Query()
		q.Set("skip_redeploy", "true")
		req.URL.RawQuery = q.Encode()
		return nil
	}
}
