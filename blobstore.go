package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	blobstoreAPIEndpoint = "service/rest/beta/blobstores"

	BlobstoreTypeFile = "File"
	BlobstoreTypeS3   = "S3"
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

func (c client) BlobstoreCreate(bs Blobstore) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := c.Post(fmt.Sprintf("%s/%s", blobstoreAPIEndpoint, strings.ToLower(bs.Type)), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not create blobstore \"%s\": HTTP: %d, %s", bs.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (c client) BlobstoreRead(id string) (*Blobstore, error) {
	body, resp, err := c.Get(blobstoreAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read blobstores: HTTP: %d, %s", resp.StatusCode, string(body))
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

func (c client) BlobstoreUpdate(id string, bs Blobstore) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := c.Put(fmt.Sprintf("%s/%s/%s", blobstoreAPIEndpoint, strings.ToLower(bs.Type), id), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update blobstore \"%s\": HTTP %d, %s", id, resp.StatusCode, string(body))
	}

	return nil
}

func (c client) BlobstoreDelete(id string) error {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", blobstoreAPIEndpoint, id))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete blobstore \"%s\": HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}
