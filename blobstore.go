package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	blobstoreAPIEndpoint = "service/rest/beta/blobstores"

	BlobstoreTypeFile = "File"
	BlobstoreTypeS3   = "S3"
)

// Blobstore data
type Blobstore struct {
	AvailableSpaceInBytes int    `json:"availableSpaceInBytes"`
	BlobCount             int    `json:"blobCount"`
	Name                  string `json:"name"`
	Path                  string `json:"path,omitempty"` // only if type File
	TotalSizeInBytes      int    `json:"totalSizeInBytes"`
	Type                  string `json:"type"`

	*BlobstoreS3BucketConfiguration `json:"bucketConfiguration,omitempty"`
	*BlobstoreSoftQuota             `json:"softQuota,omitempty"`
}

// BlobstoreSoftQuota data
type BlobstoreSoftQuota struct {
	Limit int    `json:"limit"`
	Type  string `json:"type"`
}

// BlobstoreS3BucketConfiguration data
type BlobstoreS3BucketConfiguration struct {
	*BlobstoreS3Bucket                   `json:"bucket,omitempty"`
	*BlobstoreS3Encryption               `json:"encryption,omitempty"`
	*BlobstoreS3BucketSecurity           `json:"bucketSecurity,omitempty"`
	*BlobstoreS3AdvancedBucketConnection `json:"advancedBucketConnection,omitempty"`
}

// BlobstoreS3Bucket data
type BlobstoreS3Bucket struct {
	Expiration int    `json:"expiration"`
	Name       string `json:"name"`
	Prefix     string `json:"prefix"`
	Region     string `json:"region"`
}

// BlobstoreS3Encryption data
type BlobstoreS3Encryption struct {
	Key  string `json:"encryptionKey"`
	Type string `json:"encryptionType"`
}

// BlobstoreS3BucketSecurity data
type BlobstoreS3BucketSecurity struct {
	AccessKeyID     string `json:"accessKeyId"`
	Role            string `json:"role"`
	SecretAccessKey string `json:"secretAccessKey"`
	SessionToken    string `json:"sessionToken"`
}

// BlobstoreS3AdvancedBucketConnection data
type BlobstoreS3AdvancedBucketConnection struct {
	Endpoint       string `json:"endpoint"`
	SignerType     string `json:"signerType"`
	ForcePathStyle bool   `json:"forcePathStyle"`
}

type BlobstoreSpecified struct {
	Path					 string `json:"path,omitempty"`
	*BlobstoreS3BucketConfiguration `json:"bucketConfiguration,omitempty"`
	*BlobstoreSoftQuota 			`json:"softQuota,omitempty"`
}

func (c client) BlobstoreCreate(bs Blobstore) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := c.Post(fmt.Sprintf("%s/%s", blobstoreAPIEndpoint, strings.ToLower(bs.Type)), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create blobstore \"%s\": HTTP: %d, %s", bs.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (c client) BlobstoreReadSpecified(id string) (*Blobstore, error) {
	body, err := c.RequestWrapper(c.Get, blobstoreAPIEndpoint, nil)

	if err != nil {
		return nil, err
	}

	var blobstores []Blobstore
	if err := json.Unmarshal(body, &blobstores); err != nil {
		return nil, fmt.Errorf("could not unmarshal blobstore: %v", err)
	}

	for _, bs := range blobstores {
		if bs.Name == id {
			innerBody, innerErr := c.RequestWrapper(c.Get, fmt.Sprintf("%s/%s/%s", blobstoreAPIEndpoint, strings.ToLower(bs.Type), bs.Name), nil)

			if innerErr != nil {
				return nil, innerErr
			}

			var blobstoreSpecified BlobstoreSpecified
			if err := json.Unmarshal(innerBody, &blobstoreSpecified); err != nil {
				return nil, fmt.Errorf("could not unmarshal blobstoreSpecified: %v", err)
			}

			bs.Path = blobstoreSpecified.Path
			bs.BlobstoreS3BucketConfiguration = blobstoreSpecified.BlobstoreS3BucketConfiguration
			return &bs, nil
		}
	}

	return nil, nil
}

func (c client) BlobstoreUpdate(id string, bs Blobstore) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(bs)
	if err != nil {
		return err
	}

	body, resp, err := c.Put(fmt.Sprintf("%s/%s/%s", blobstoreAPIEndpoint, strings.ToLower(bs.Type), id), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update blobstore \"%s\": HTTP %d, %s", id, resp.StatusCode, string(body))
	}

	return nil
}

func (c client) BlobstoreDelete(id string) error {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", blobstoreAPIEndpoint, id), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete blobstore \"%s\": HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}
