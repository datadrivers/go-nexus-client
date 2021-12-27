package blobstore

const (
	BlobstoreTypeFile = "File"
	BlobstoreTypeS3   = "S3"
)

// Legacy data
type Legacy struct {
	AvailableSpaceInBytes int    `json:"availableSpaceInBytes"`
	BlobCount             int    `json:"blobCount"`
	Name                  string `json:"name"`
	Path                  string `json:"path,omitempty"` // only if type File
	TotalSizeInBytes      int    `json:"totalSizeInBytes"`
	Type                  string `json:"type"`

	*S3BucketConfiguration `json:"bucketConfiguration,omitempty"`
	*SoftQuota             `json:"softQuota,omitempty"`
}
