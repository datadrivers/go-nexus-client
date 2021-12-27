package repository

// Group contains repository group configuration data
type Group struct {
	MemberNames []string `json:"memberNames,omitempty"`
}

// GroupDeploy
type GroupDeploy struct {
	// Member repositories' names
	MemberNames []string `json:"memberNames,omitempty"`
	// Pro-only: This field is for the Group Deployment feature available in NXRM Pro.
	WritableMember *string `json:"writableMember,omitempty"`
}
