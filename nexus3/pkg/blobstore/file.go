package blobstore

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/blobstore"
)

type BlobStoreFileService client.Service

func NewBlobStoreFileService(c *client.Client) *BlobStoreFileService {

	s := &BlobStoreFileService{
		Client: c,
	}
	return s
}

func (s *BlobStoreFileService) Create(bs *blobstore.File) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Post(fmt.Sprintf("%s/file", blobstoreAPIEndpoint), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create blobstore \"%s\": HTTP: %d, %s", bs.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *BlobStoreFileService) Get(name string) (*blobstore.File, error) {
	body, resp, err := s.Client.Get(fmt.Sprintf("%s/file/%s", blobstoreAPIEndpoint, name), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read file blobstores: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var bs blobstore.File
	if err := json.Unmarshal(body, &bs); err != nil {
		return nil, fmt.Errorf("could not unmarshal blobstore \"%s\": %v", name, err)
	}
	bs.Name = name
	return &bs, nil
}

func (s *BlobStoreFileService) Update(name string, bs *blobstore.File) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Put(fmt.Sprintf("%s/file/%s", blobstoreAPIEndpoint, name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update blobstore \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *BlobStoreFileService) Delete(name string) error {
	return deleteBlobstore(s.Client, name)
}

func (s *BlobStoreFileService) GetQuotaStatus(name string) (*blobstore.QuotaStatus, error) {
	return getBlobstoreQuotaStatus(s.Client, name)
}
