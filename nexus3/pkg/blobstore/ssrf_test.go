package blobstore

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
)

// allowSSRFTo adds the host of rawURL to the SSRF protection allow list.
//
// Since Nexus 3.93.0 an S3 blob store pointed at a host that resolves to a
// private address is rejected with "Host resolves to private/local IP
// address(es) ... update the SSRF protection settings via the API at
// /v1/security/ssrf-protection". The MinIO container the tests run against is
// exactly such a host, so the allow list has to name it before the blob store
// can be created.
//
// Older Nexus versions do not serve the endpoint; a 404 is ignored so the
// tests keep working against them.
func allowSSRFTo(t *testing.T, c *client.Client, rawURL string) {
	t.Helper()

	if rawURL == "" {
		return
	}

	parsed, err := url.Parse(rawURL)
	if err != nil || parsed.Hostname() == "" {
		return
	}

	payload, err := tools.JsonMarshalInterfaceToIOReader(map[string]any{
		"enabled":        true,
		"allowedIPs":     []string{},
		"allowedDomains": []string{parsed.Hostname()},
	})
	if err != nil {
		t.Fatalf("could not build SSRF protection payload: %v", err)
	}

	body, resp, err := c.Put(client.BasePath+"v1/security/ssrf-protection", payload)
	if err != nil {
		t.Fatalf("could not update SSRF protection settings: %v", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		t.Fatalf("could not update SSRF protection settings: HTTP: %d, %s", resp.StatusCode, string(body))
	}
}
