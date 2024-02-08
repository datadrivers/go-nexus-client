package readonly

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/stretchr/testify/assert"
)

var (
	testClient *client.Client = nil
)

func getTestClient() *client.Client {
	if testClient != nil {
		return testClient
	}
	return client.NewClient(getDefaultConfig())
}

func getTestService() *ReadOnlyService {
	return NewReadOnlyService(getTestClient())
}

func getDefaultConfig() client.Config {
	return client.Config{
		Insecure: tools.GetEnv("NEXUS_INSECURE_SKIP_VERIFY", true).(bool),
		Password: tools.GetEnv("NEXUS_PASSWORD", "admin123").(string),
		URL:      tools.GetEnv("NEXUS_URL", "http://127.0.0.1:8081").(string),
		Username: tools.GetEnv("NEXUS_USRNAME", "admin").(string),
		Timeout:  tools.GetEnv("NEXUS_TIMEOUT", 30).(int),
	}
}

func TestFreezeAndReleaseReadOnlyState(t *testing.T) {
	s := getTestService()

	state, err := s.GetState()
	if err != nil {
		assert.Failf(t, "fail to retreive readonly state", err.Error())
		return
	}

	if state.Frozen {
		// try release first
		if err = s.Release(); err != nil && err != ErrNoChangeToReadOnlyState {
			assert.Failf(t, "unexpected error when try to release", err.Error())
		}
	}

	// freeze
	if err = s.Freeze(); err != nil && err != ErrNoChangeToReadOnlyState {
		assert.Failf(t, "unexpected error when try to freeze", err.Error())
	}

	// release
	if err = s.Release(); err != nil && err != ErrNoChangeToReadOnlyState {
		assert.Failf(t, "unexpected error when try to release", err.Error())
	}

	// freeze again
	if err = s.Freeze(); err != nil && err != ErrNoChangeToReadOnlyState {
		assert.Failf(t, "unexpected error when try to freeze again", err.Error())
	}

	// force release
	if err = s.ForceRelease(); err != nil && err != ErrNoChangeToReadOnlyState {
		assert.Failf(t, "unexpected error when try to force release", err.Error())
	}
}
