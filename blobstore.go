package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	blobstoreAPIEndpoint = "service/rest/beta/blobstores"

	BlobstoreTypeFile = "file"
	BlobstoreTypeS3   = "s3"
)

// Blobstore data
type Blobstore struct {
	AvailableSpaceInBytes int    `json:"availableSpaceInBytes"`
	BlobCount             int    `json:"blobCount"`
	Name                  string `json:"name"`
	Path                  string `json:"path"`
	TotalSizeInBytes      int    `json:"totalSizeInBytes"`
	Type                  string `json:"type"`

	*BlobstoreSoftQuota `json:"softQuota,omitempty"`
}

// BlobstoreSoftQuota data
type BlobstoreSoftQuota struct {
	Limit int    `json:"limit"`
	Type  string `json:"type"`
}

func (c client) BlobstoreCreate(bs Blobstore, bsType string) error {
	return nil
}

func (c client) BlobstoreRead(id string) (*Blobstore, error) {
	body, resp, err := c.Get(blobstoreAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(body))
	}

	var blobstores []Blobstore
	if err := json.Unmarshal(body, &blobstores); err != nil {
		return nil, fmt.Errorf("could not unmarshal blobstore: %v", err)
	}

	for _, bs := range blobstores {
		if bs.Name == id {
			return &bs, nil
		}
	}

	return nil, nil
}

func (c client) BlobstoreUpdate(id string, bs Blobstore, bsType string) error {
	return nil
}

func (c client) BlobstoreDelete(id string) error {
	return nil
}
