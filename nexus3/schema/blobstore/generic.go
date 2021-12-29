package blobstore

// Generic data
type Generic struct {
	Name string `json:"name"`
	Type string `json:"type"`

	Unavailable           bool `json:"unavailable"`
	AvailableSpaceInBytes int  `json:"availableSpaceInBytes"`
	BlobCount             int  `json:"blobCount"`
	TotalSizeInBytes      int  `json:"totalSizeInBytes"`

	*SoftQuota `json:"softQuota,omitempty"`
}
