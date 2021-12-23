package nexus3

import (
	"fmt"
	"net/http"
)

const (
	blobstoreAPIEndpoint = basePath + "v1/blobstores"
)

type BlobStoreService struct {
	client *client

	// API Services
	Azure *BlobStoreAzureService
	File  *BlobStoreFileService
	Group *BlobStoreGroupService
	S3    *BlobStoreS3Service
}

func NewBlobStoreService(c *client) *BlobStoreService {
	a := NewBlobStoreAzureService(c)
	f := NewBlobStoreFileService(c)
	g := NewBlobStoreGroupService(c)
	s := NewBlobStoreS3Service(c)
	return &BlobStoreService{
		client: c,
		Azure:  a,
		File:   f,
		Group:  g,
		S3:     s,
	}
}

func (s *BlobStoreService) Delete(name string) error {
	return deleteBlobstore(s.client, name)
}

func deleteBlobstore(c *client, name string) error {
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
	return getBlobstoreQuotaStatus(s.client, name)
}

func getBlobstoreQuotaStatus(c *client, name string) error {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", blobstoreAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete blobstore \"%s\": HTTP: %d, %s", name, resp.StatusCode, string(body))
	}
	return nil
}
