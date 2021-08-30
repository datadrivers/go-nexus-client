package nexus3

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BlobStoreS3Service service

func NewBlobStoreS3Service(c *client) *BlobStoreS3Service {

	s := &BlobStoreS3Service{
		client: c,
	}
	return s
}

type S3BlobStore struct {
	// The name of the S3 blob store
	Name string `json:"name"`
	// Settings to control the soft quota
	SoftQuota *BlobStoreSoftQuota `json:"softQuota,omitempty"`
	// The S3 specific configuration details for the S3 object that'll contain the blob store.
	BucketConfiguration BlobStoreS3BucketConfiguration `json:"bucketConfiguration"`
}

type BlobStoreS3BucketConfiguration struct {
	// Details of the S3 bucket such as name and region
	Bucket BlobStoreS3Bucket `json:"bucket"`

	// The type of encryption to use if any
	Encryption *BlobStoreS3Encryption `json:"encryption,omitempty"`

	// Security details for granting access the S3 API
	BucketSecurity *BlobStoreS3BucketSecurity `json:"bucketSecurity,omitempty"`

	// A custom endpoint URL, signer type and whether path style access is enabled
	AdvancedBucketConnection *BlobStoreS3AdvancedBucketConnection `json:"advancedBucketConnection,omitempty"`
}

type BlobStoreS3Bucket struct {
	// The AWS region to create a new S3 bucket in or an existing S3 bucket's region
	Region string `json:"region"`

	// The name of the S3 bucket
	Name string `json:"name,omitempty"`

	// The S3 blob store (i.e S3 object) key prefix
	Prefix string `json:"prefix,omitempty"`

	// How many days until deleted blobs are finally removed from the S3 bucket (-1 to disable)
	Expiration int32 `json:"expiration"`
}

type BlobStoreS3Encryption struct {
	// The encryption key
	Key string `json:"encryptionKey"`

	// The type of S3 server side encryption to use
	Type string `json:"encryptionType"`
}

type BlobStoreS3BucketSecurity struct {
	// An IAM access key ID for granting access to the S3 bucket
	AccessKeyID string `json:"accessKeyId"`

	// An IAM role to assume in order to access the S3 bucket
	Role string `json:"role"`

	// The secret access key associated with the specified IAM access key ID
	SecretAccessKey string `json:"secretAccessKey"`

	// An AWS STS session token associated with temporary security credentials which grant access to the S3 bucket
	SessionToken string `json:"sessionToken"`
}

type BlobStoreS3AdvancedBucketConnection struct {
	// A custom endpoint URL for third party object stores using the S3 API
	Endpoint string `json:"endpoint"`

	// An API signature version which may be required for third party object stores using the S3 API
	SignerType string `json:"signerType"`

	// Setting this flag will result in path-style access being used for all requests
	ForcePathStyle bool `json:"forcePathStyle"`

	// Setting this value will override the default connection pool size of Nexus of the s3 client for this blobstore.
	MaxConnectionPoolSize int32 `json:"maxConnectionPoolSize"`
}

func (s *BlobStoreS3Service) Create(bs *S3BlobStore) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Post(fmt.Sprintf("%s/s3", blobstoreAPIEndpoint), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create blobstore \"%s\": HTTP: %d, %s", bs.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *BlobStoreS3Service) Get(name string) (*S3BlobStore, error) {
	body, resp, err := s.client.Get(fmt.Sprintf("%s/s3/%s", blobstoreAPIEndpoint, name), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read file blobstores: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var bs S3BlobStore
	if err := json.Unmarshal(body, &bs); err != nil {
		return nil, fmt.Errorf("could not unmarshal blobstore \"%s\": %v", name, err)
	}
	return &bs, nil
}

func (s *BlobStoreS3Service) Update(name string, bs *S3BlobStore) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(fmt.Sprintf("%s/s3/%s", blobstoreAPIEndpoint, name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update blobstore \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *BlobStoreS3Service) Delete(name string) error {
	return deleteBlobstore(s.client, name)
}

func (s *BlobStoreS3Service) GetQuotaStatus(name string) error {
	return getBlobstoreQuotaStatus(s.client, name)
}
