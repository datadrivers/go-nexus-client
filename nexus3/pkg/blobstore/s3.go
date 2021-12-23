package blobstore

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/blobstore"
)

type BlobStoreS3Service client.Service

func NewBlobStoreS3Service(c *client.Client) *BlobStoreS3Service {

	s := &BlobStoreS3Service{
		Client: c,
	}
	return s
}

func (s *BlobStoreS3Service) Create(bs *blobstore.S3) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Post(fmt.Sprintf("%s/s3", blobstoreAPIEndpoint), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create blobstore \"%s\": HTTP: %d, %s", bs.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *BlobStoreS3Service) Get(name string) (*blobstore.S3, error) {
	body, resp, err := s.Client.Get(fmt.Sprintf("%s/s3/%s", blobstoreAPIEndpoint, name), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read file blobstores: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var bs blobstore.S3
	if err := json.Unmarshal(body, &bs); err != nil {
		return nil, fmt.Errorf("could not unmarshal blobstore \"%s\": %v", name, err)
	}
	return &bs, nil
}

func (s *BlobStoreS3Service) Update(name string, bs *blobstore.S3) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Put(fmt.Sprintf("%s/s3/%s", blobstoreAPIEndpoint, name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update blobstore \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *BlobStoreS3Service) Delete(name string) error {
	return deleteBlobstore(s.Client, name)
}

func (s *BlobStoreS3Service) GetQuotaStatus(name string) error {
	return getBlobstoreQuotaStatus(s.Client, name)
}
