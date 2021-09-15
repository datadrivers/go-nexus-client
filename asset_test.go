package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssetRead(t *testing.T) {
	client := getTestClient()

	id := ""

	asset, err := client.AssetRead(id)
	assert.Nil(t, err)
	assert.NotNil(t, asset)

	if asset != nil {
		assert.Equal(t, id, asset.ID)
		assert.Equal(t, "maven-central", asset.Repository)
	}
}
