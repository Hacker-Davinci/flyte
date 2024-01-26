package agent

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	agentMocks "github.com/flyteorg/flyte/flyteplugins/go/tasks/plugins/webapi/agent/mocks"
)

func mockGetBadAsyncClientFunc() *agentMocks.AsyncAgentServiceClient {
	return nil
}

func TestInitializeClientFunc(t *testing.T) {
	cfg := defaultConfig
	ctx := context.Background()
	err := SetConfig(&cfg)
	assert.NoError(t, err)
	cs, err := initializeClients(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, cs)
}
