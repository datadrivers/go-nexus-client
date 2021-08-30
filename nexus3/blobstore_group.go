package nexus3

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	BlobstoreGroupFillPolicyRoundRobin   = "roundRobin"
	BlobstoreGroupFillPolicyWriteToFirst = "writeToFirst"
)

type BlobStoreGroupService service

func NewBlobStoreGroupService(c *client) *BlobStoreGroupService {

	s := &BlobStoreGroupService{
		client: c,
	}
	return s
}

type GroupBlobStore struct {
	// The name of the Group blob store
	Name string `json:"name"`
	// Settings to control the soft quota
	SoftQuota *BlobStoreSoftQuota `json:"softQuota,omitempty"`

	// List of the names of blob stores that are members of this group
	Members []string `json:"members,omitempty"`

	// Possible values: roundRobin,writeToFirst
	FillPolicy string `json:"fillPolicy"`
}

func (s *BlobStoreGroupService) Create(bs *GroupBlobStore) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Post(fmt.Sprintf("%s/group", blobstoreAPIEndpoint), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create blobstore \"%s\": HTTP: %d, %s", bs.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *BlobStoreGroupService) Get(name string) (*GroupBlobStore, error) {
	body, resp, err := s.client.Get(fmt.Sprintf("%s/group/%s", blobstoreAPIEndpoint, name), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read file blobstores: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var bs GroupBlobStore
	if err := json.Unmarshal(body, &bs); err != nil {
		return nil, fmt.Errorf("could not unmarshal blobstore \"%s\": %v", name, err)
	}
	bs.Name = name
	return &bs, nil
}

func (s *BlobStoreGroupService) Update(name string, bs *GroupBlobStore) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(fmt.Sprintf("%s/group/%s", blobstoreAPIEndpoint, name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update blobstore \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *BlobStoreGroupService) Delete(name string) error {
	return deleteBlobstore(s.client, name)
}

func (s *BlobStoreGroupService) GetQuotaStatus(name string) error {
	return getBlobstoreQuotaStatus(s.client, name)
}
