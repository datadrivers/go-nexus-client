package security

type SSLCertificate struct {
	Id                      string `json:"id"`
	Fingerprint             string `json:"fingerprint"`
	SerialNumber            string `json:"serialNumber"`
	IssuerCommonName        string `json:"issuerCommonName"`
	IssuerOrganization      string `json:"issuerOrganization"`
	IssuerOrganizationUnit  string `json:"issuerOrganizationalUnit"`
	SubjectCommonName       string `json:"subjectCommonName"`
	SubjectOrganization     string `json:"subjectOrganization"`
	SubjectOrganizationUnit string `json:"subjectOrganizationalUnit"`
	Pem                     string `json:"pem"`
	IssuedOn                int64  `json:"issuedOn"`
	ExpiresOn               int64  `json:"expiresOn"`
}
