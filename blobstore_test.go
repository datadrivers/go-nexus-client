package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlobstoreCreate(t *testing.T) {

}

func TestBlobstoreRead(t *testing.T) {
	client := NewClient(getDefaultConfig())

	bsName := "default"

	bs, err := client.BlobstoreRead(bsName)
	assert.Nil(t, err)
	assert.NotNil(t, bs)

	if bs != nil {
		assert.Equal(t, bsName, bs.Name)
	}
}

func TestBlobstoreUpdate(t *testing.T) {

}

func TestBlobstoreDelete(t *testing.T) {

}
