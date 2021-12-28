package blobstore

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/blobstore"
)

type BlobStoreLegacyService client.Service

func NewBlobStoreLegacyService(c *client.Client) *BlobStoreLegacyService {

	s := &BlobStoreLegacyService{
		Client: c,
	}
	return s
}

func (s *BlobStoreLegacyService) Create(bs *blobstore.Legacy) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Post(fmt.Sprintf("%s/%s", blobstoreAPIEndpoint, strings.ToLower(bs.Type)), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create blobstore \"%s\": HTTP: %d, %s", bs.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *BlobStoreLegacyService) Get(name string) (*blobstore.Legacy, error) {
	body, resp, err := s.Client.Get(blobstoreAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read file blobstores: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var blobstores []blobstore.Legacy
	if err := json.Unmarshal(body, &blobstores); err != nil {
		return nil, fmt.Errorf("could not unmarshal blobstore \"%s\": %v", name, err)
	}

	for _, bs := range blobstores {
		if bs.Name == name {
			bsDetailed, err := s.ReadDetails(name, bs.Type)
			if err != nil {
				return nil, err
			}

			bsDetailed.AvailableSpaceInBytes = bs.AvailableSpaceInBytes
			bsDetailed.BlobCount = bs.BlobCount
			bsDetailed.Name = bs.Name
			bsDetailed.TotalSizeInBytes = bs.TotalSizeInBytes
			bsDetailed.Type = bs.Type

			return bsDetailed, nil
		}
	}

	return nil, nil
}

func (s *BlobStoreLegacyService) Update(id string, bs blobstore.Legacy) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Put(fmt.Sprintf("%s/%s/%s", blobstoreAPIEndpoint, strings.ToLower(bs.Type), id), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update blobstore \"%s\": HTTP %d, %s", id, resp.StatusCode, string(body))
	}

	return nil
}

func (s *BlobStoreLegacyService) Delete(name string) error {
	return deleteBlobstore(s.Client, name)
}

func (s *BlobStoreLegacyService) GetQuotaStatus(name string) error {
	return getBlobstoreQuotaStatus(s.Client, name)
}

func (s *BlobStoreLegacyService) ReadDetails(id string, bsType string) (*blobstore.Legacy, error) {
	body, resp, err := s.Client.Get(fmt.Sprintf("%s/%s/%s", blobstoreAPIEndpoint, strings.ToLower(bsType), id), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read blobstore \"%s\" of type \"%s\": HTTP: %d, %s", id, bsType, resp.StatusCode, string(body))
	}

	blobstore := &blobstore.Legacy{}
	if err := json.Unmarshal(body, blobstore); err != nil {
		return nil, fmt.Errorf("could not unmarshal details of blobstore \"%s\": %v", id, err)
	}

	return blobstore, nil
}
