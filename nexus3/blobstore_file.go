package nexus3

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BlobStoreFileService service

func NewBlobStoreFileService(c *client) *BlobStoreFileService {

	s := &BlobStoreFileService{
		client: c,
	}
	return s
}

type FileBlobStore struct {
	// Name of the BlobStore
	Name string `json:"name"`
	// Settings to control the soft quota
	SoftQuota *BlobStoreSoftQuota `json:"softQuota,omitempty"`
	// The path to the blobstore contents. This can be an absolute path to anywhere on the system Nexus Repository Manager has access to or it can be a path relative to the sonatype-work directory.
	Path string `json:"path,omitempty"`
}

func (s *BlobStoreFileService) Create(bs *FileBlobStore) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Post(fmt.Sprintf("%s/file", blobstoreAPIEndpoint), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create blobstore \"%s\": HTTP: %d, %s", bs.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *BlobStoreFileService) Get(name string) (*FileBlobStore, error) {
	body, resp, err := s.client.Get(fmt.Sprintf("%s/file/%s", blobstoreAPIEndpoint, name), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read file blobstores: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var bs FileBlobStore
	if err := json.Unmarshal(body, &bs); err != nil {
		return nil, fmt.Errorf("could not unmarshal blobstore \"%s\": %v", name, err)
	}
	bs.Name = name
	return &bs, nil
}

func (s *BlobStoreFileService) Update(name string, bs *FileBlobStore) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(fmt.Sprintf("%s/file/%s", blobstoreAPIEndpoint, name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update blobstore \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *BlobStoreFileService) Delete(name string) error {
	return deleteBlobstore(s.client, name)
}

func (s *BlobStoreFileService) GetQuotaStatus(name string) error {
	return getBlobstoreQuotaStatus(s.client, name)
}
