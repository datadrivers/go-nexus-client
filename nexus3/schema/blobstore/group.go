package blobstore

const (
	GroupFillPolicyRoundRobin   = "roundRobin"
	GroupFillPolicyWriteToFirst = "writeToFirst"
)

type Group struct {
	// The name of the Group blob store
	Name string `json:"name"`
	// Settings to control the soft quota
	SoftQuota *SoftQuota `json:"softQuota,omitempty"`

	// List of the names of blob stores that are members of this group
	Members []string `json:"members"`

	// Possible values: roundRobin,writeToFirst
	FillPolicy string `json:"fillPolicy"`
}
