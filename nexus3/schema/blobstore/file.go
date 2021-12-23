package blobstore

type File struct {
	// Name of the BlobStore
	Name string `json:"name"`
	// Settings to control the soft quota
	SoftQuota *SoftQuota `json:"softQuota,omitempty"`
	// The path to the blobstore contents. This can be an absolute path to anywhere on the system Nexus Repository Manager has access to or it can be a path relative to the sonatype-work directory.
	Path string `json:"path,omitempty"`
}
