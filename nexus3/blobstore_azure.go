package nexus3

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/blobstore"
)

type BlobStoreAzureService service

func NewBlobStoreAzureService(c *client) *BlobStoreAzureService {

	s := &BlobStoreAzureService{
		client: c,
	}
	return s
}

func (s *BlobStoreAzureService) Create(bs *blobstore.Azure) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Post(fmt.Sprintf("%s/azure", blobstoreAPIEndpoint), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create blobstore \"%s\": HTTP: %d, %s", bs.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *BlobStoreAzureService) Get(name string) (*blobstore.Azure, error) {
	body, resp, err := s.client.Get(fmt.Sprintf("%s/azure/%s", blobstoreAPIEndpoint, name), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read azure blobstores: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var bs blobstore.Azure
	if err := json.Unmarshal(body, &bs); err != nil {
		return nil, fmt.Errorf("could not unmarshal blobstore \"%s\": %v", name, err)
	}
	bs.Name = name
	return &bs, nil
}

func (s *BlobStoreAzureService) Update(name string, bs *blobstore.Azure) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(fmt.Sprintf("%s/azure/%s", blobstoreAPIEndpoint, name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update blobstore \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *BlobStoreAzureService) Delete(name string) error {
	return deleteBlobstore(s.client, name)
}

func (s *BlobStoreAzureService) GetQuotaStatus(name string) error {
	return getBlobstoreQuotaStatus(s.client, name)
}

func (s *BlobStoreAzureService) TestConnection(bs *blobstore.Azure) error {
	con := &blobstore.AzureConnection{
		AccountName:          bs.BucketConfiguration.AccountName,
		ContainerName:        bs.BucketConfiguration.ContainerName,
		AuthenticationMethod: bs.BucketConfiguration.Authentication.AuthenticationMethod,
		AccountKey:           bs.BucketConfiguration.Authentication.AccountKey,
	}
	ioReader, err := jsonMarshalInterfaceToIOReader(con)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Post(fmt.Sprintf("%sv1/azureblobstore/test-connection", basePath), ioReader)
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusBadRequest:
		return fmt.Errorf("could not connect to azure storage account container: Azure Blob Store connection failed")
	case http.StatusUnauthorized:
		return fmt.Errorf("could not connect to azure storage account container: Authentication required")
	case http.StatusForbidden:
		return fmt.Errorf("could not connect to azure storage account container: Insufficient permissions")
	default:
		return fmt.Errorf("could not connect to azure storage account container: HTTP: %d, %s", resp.StatusCode, string(body))
	}

}
