package blobstore

type S3 struct {
	// The name of the S3 blob store
	Name string `json:"name"`
	// Settings to control the soft quota
	SoftQuota *SoftQuota `json:"softQuota,omitempty"`
	// The S3 specific configuration details for the S3 object that'll contain the blob store.
	BucketConfiguration S3BucketConfiguration `json:"bucketConfiguration"`
}

type S3BucketConfiguration struct {
	// Details of the S3 bucket such as name and region
	Bucket S3Bucket `json:"bucket"`

	// The type of encryption to use if any
	Encryption *S3Encryption `json:"encryption,omitempty"`

	// Security details for granting access the S3 API
	BucketSecurity *S3BucketSecurity `json:"bucketSecurity,omitempty"`

	// A custom endpoint URL, signer type and whether path style access is enabled
	AdvancedBucketConnection *S3AdvancedBucketConnection `json:"advancedBucketConnection,omitempty"`
}

type S3Bucket struct {
	// The AWS region to create a new S3 bucket in or an existing S3 bucket's region
	Region string `json:"region"`

	// The name of the S3 bucket
	Name string `json:"name,omitempty"`

	// The S3 blob store (i.e S3 object) key prefix
	Prefix string `json:"prefix,omitempty"`

	// How many days until deleted blobs are finally removed from the S3 bucket (-1 to disable)
	Expiration int32 `json:"expiration"`
}

type S3Encryption struct {
	// The encryption key
	Key string `json:"encryptionKey"`

	// The type of S3 server side encryption to use
	Type string `json:"encryptionType"`
}

type S3BucketSecurity struct {
	// An IAM access key ID for granting access to the S3 bucket
	AccessKeyID string `json:"accessKeyId"`

	// An IAM role to assume in order to access the S3 bucket
	Role string `json:"role"`

	// The secret access key associated with the specified IAM access key ID
	SecretAccessKey string `json:"secretAccessKey"`

	// An AWS STS session token associated with temporary security credentials which grant access to the S3 bucket
	SessionToken string `json:"sessionToken"`
}

type S3AdvancedBucketConnection struct {
	// A custom endpoint URL for third party object stores using the S3 API
	Endpoint string `json:"endpoint"`

	// An API signature version which may be required for third party object stores using the S3 API
	SignerType string `json:"signerType"`

	// Setting this flag will result in path-style access being used for all requests
	ForcePathStyle bool `json:"forcePathStyle"`

	// Setting this value will override the default connection pool size of Nexus of the s3 client for this blobstore.
	MaxConnectionPoolSize int32 `json:"maxConnectionPoolSize"`
}
