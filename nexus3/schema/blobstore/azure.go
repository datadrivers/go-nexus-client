package blobstore

const (
	AzureAuthenticationMethodAccountKey      AzureAuthenticationMethod = "ACCOUNTKEY"
	AzureAuthenticationMethodManagedIdentity AzureAuthenticationMethod = "MANAGEDIDENTITY"
)

type AzureAuthenticationMethod string

type Azure struct {
	// Name of the BlobStore
	Name string `json:"name"`
	// Settings to control the soft quota
	SoftQuota *SoftQuota `json:"softQuota,omitempty"`
	// The Azure specific configuration details for the Azure object that'll contain the blob store.
	BucketConfiguration AzureBucketConfiguration `json:"bucketConfiguration"`
}

type AzureBucketConfiguration struct {
	// Account name found under Access keys for the storage account.
	AccountName string `json:"accountName"`

	// The Azure specific authentication details.
	Authentication AzureBucketConfigurationAuthentication `json:"authentication"`

	// The name of an existing container to be used for storage.
	ContainerName string `json:"containerName"`
}

type AzureConnection struct {
	// Account name found under Access keys for the storage account.
	AccountName string `json:"accountName"`

	// The type of Azure authentication to use.
	AuthenticationMethod AzureAuthenticationMethod `json:"authenticationMethod"`

	// The account key.
	AccountKey string `json:"accountKey,omitempty"`
	// The name of an existing container to be used for storage.
	ContainerName string `json:"containerName"`
}

type AzureBucketConfigurationAuthentication struct {
	// The type of Azure authentication to use.
	AuthenticationMethod AzureAuthenticationMethod `json:"authenticationMethod"`

	// The account key.
	AccountKey string `json:"accountKey,omitempty"`
}
