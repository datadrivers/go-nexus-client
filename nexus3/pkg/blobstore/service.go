package blobstore

import (
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
)

const (
	blobstoreAPIEndpoint = client.BasePath + "v1/blobstores"
)

type BlobStoreService struct {
	Client *client.Client

	// API Services
	Azure *BlobStoreAzureService
	File  *BlobStoreFileService
	Group *BlobStoreGroupService
	S3    *BlobStoreS3Service
}

func NewBlobStoreService(c *client.Client) *BlobStoreService {
	return &BlobStoreService{
		Client: c,

		Azure: NewBlobStoreAzureService(c),
		File:  NewBlobStoreFileService(c),
		Group: NewBlobStoreGroupService(c),
		S3:    NewBlobStoreS3Service(c),
	}
}

func (s *BlobStoreService) Delete(name string) error {
	return deleteBlobstore(s.Client, name)
}

func deleteBlobstore(c *client.Client, name string) error {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", blobstoreAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete blobstore \"%s\": HTTP: %d, %s", name, resp.StatusCode, string(body))
	}
	return nil
}

func (s *BlobStoreService) GetQuotaStatus(name string) error {
	return getBlobstoreQuotaStatus(s.Client, name)
}

func getBlobstoreQuotaStatus(c *client.Client, name string) error {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", blobstoreAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete blobstore \"%s\": HTTP: %d, %s", name, resp.StatusCode, string(body))
	}
	return nil
}
