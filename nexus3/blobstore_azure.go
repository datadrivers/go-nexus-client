package nexus3

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	BlobstoreAzureAuthenticationMethodAccountKey      = "ACCOUNTKEY"
	BlobstoreAzureAuthenticationMethodManagedIdentity = "MANAGEDIDENTITY"
)

type BlobStoreAzureService service

func NewBlobStoreAzureService(c *client) *BlobStoreAzureService {

	s := &BlobStoreAzureService{
		client: c,
	}
	return s
}

type AzureBlobStore struct {
	// Name of the BlobStore
	Name string `json:"name"`
	// Settings to control the soft quota
	SoftQuota *BlobStoreSoftQuota `json:"softQuota,omitempty"`
	// The Azure specific configuration details for the Azure object that'll contain the blob store.
	BucketConfiguration BlobStoreAzureBucketConfiguration `json:"bucketConfiguration,omitempty"`
}

type BlobStoreAzureBucketConfiguration struct {
	// Account name found under Access keys for the storage account.
	AccountName string `json:"accountName"`

	// The Azure specific authentication details.
	Authentication BlobStoreAzureBucketConfigurationAuthentication `json:"authentication"`

	// The name of an existing container to be used for storage.
	ContainerName string `json:"containerName"`
}

type blobStoreAzureConnection struct {
	// Account name found under Access keys for the storage account.
	AccountName string `json:"accountName"`

	// The type of Azure authentication to use.
	AuthenticationMethod string `json:"authenticationMethod"`

	// The account key.
	AccountKey string `json:"accountKey,omitempty"`
	// The name of an existing container to be used for storage.
	ContainerName string `json:"containerName"`
}

type BlobStoreAzureBucketConfigurationAuthentication struct {
	// The type of Azure authentication to use.
	AuthenticationMethod string `json:"authenticationMethod"`

	// The account key.
	AccountKey string `json:"accountKey,omitempty"`
}

func (s *BlobStoreAzureService) Create(bs *AzureBlobStore) error {
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

func (s *BlobStoreAzureService) Get(name string) (*AzureBlobStore, error) {
	body, resp, err := s.client.Get(fmt.Sprintf("%s/azure/%s", blobstoreAPIEndpoint, name), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read azure blobstores: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var bs AzureBlobStore
	if err := json.Unmarshal(body, &bs); err != nil {
		return nil, fmt.Errorf("could not unmarshal blobstore \"%s\": %v", name, err)
	}
	bs.Name = name
	return &bs, nil
}

func (s *BlobStoreAzureService) Update(name string, bs *AzureBlobStore) error {
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

func (s *BlobStoreAzureService) TestConnection(bs *AzureBlobStore) error {
	con := &blobStoreAzureConnection{
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
