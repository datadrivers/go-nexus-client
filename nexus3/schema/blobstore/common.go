package blobstore

type SoftQuota struct {
	// The type to use such as spaceRemainingQuota, or spaceUsedQuota
	Type string `json:"type,omitempty"`
	// The limit in MB.
	Limit int64 `json:"limit,omitempty"`
}

type QuotaStatus struct {
	IsViolation   bool   `json:"isViolation"`
	Message       string `json:"message,omitempty"`
	BlobStoreName string `json:"blobStoreName"`
}
