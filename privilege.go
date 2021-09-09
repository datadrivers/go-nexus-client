package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	privilegeAPIEndpoint = basePath + "v1/security/privileges"

	// PrivilegeDomains
	PrivilegeDomainAll                = "*"
	PrivilegeDomainAPIKey             = "apikey"
	PrivilegeDomainAnalytics          = "analytics"
	PrivilegeDomainAtlas              = "atlas"
	PrivilegeDomainBlobstores         = "blobstores"
	PrivilegeDomainBundles            = "bundles"
	PrivilegeDomainCapabilities       = "capabilities"
	PrivilegeDomainComponent          = "component"
	PrivilegeDomainDatastores         = "datastores"
	PrivilegeDomainHealthcheck        = "healthcheck"
	PrivilegeDomainHealthcheckSummary = "healthchecksummary"
	PrivilegeDomainIQViolationSummery = "iq-violation-summary"
	PrivilegeDomainLDAP               = "ldap"
	PrivilegeDomainLicensing          = "licensing"
	PrivilegeDomainLogging            = "logging"
	PrivilegeDomainMetrics            = "metrics"
	PrivilegeDomainPrivileges         = "privileges"
	PrivilegeDomainRoles              = "roles"
	PrivilegeDomainSearch             = "search"
	PrivilegeDomainSelectors          = "selectors"
	PrivilegeDomainSettings           = "settings"
	PrivilegeDomainSSLTruststore      = "ssl-truststore"
	PrivilegeDomainTasks              = "tasks"
	PrivilegeDomainUsers              = "users"
	PrivilegeDomainUsersChangePW      = "userschangepw"
	PrivilegeDomainWonderland         = "wonderland"

	// PrivilegeTypes
	PrivilegeTypeApplication     = "application"
	PrivilegeTypeContentSelector = "repository-content-selector"
	PrivilegeTypeRepositoryAdmin = "repository-admin"
	PrivilegeTypeRepositoryView  = "repository-view"
	PrivilegeTypeScript          = "script"
	PrivilegeTypeWildcard        = "wildcard"
)

var (
	// PrivilegeDomains represents a string slice of supported privilege domains
	PrivilegeDomains []string = []string{
		PrivilegeDomainAll,
		PrivilegeDomainAPIKey,
		PrivilegeDomainAnalytics,
		PrivilegeDomainAtlas,
		PrivilegeDomainBlobstores,
		PrivilegeDomainBundles,
		PrivilegeDomainCapabilities,
		PrivilegeDomainComponent,
		PrivilegeDomainDatastores,
		PrivilegeDomainHealthcheck,
		PrivilegeDomainHealthcheckSummary,
		PrivilegeDomainIQViolationSummery,
		PrivilegeDomainLDAP,
		PrivilegeDomainLicensing,
		PrivilegeDomainLogging,
		PrivilegeDomainMetrics,
		PrivilegeDomainPrivileges,
		PrivilegeDomainRoles,
		PrivilegeDomainSearch,
		PrivilegeDomainSelectors,
		PrivilegeDomainSettings,
		PrivilegeDomainSSLTruststore,
		PrivilegeDomainTasks,
		PrivilegeDomainUsers,
		PrivilegeDomainUsersChangePW,
		PrivilegeDomainWonderland,
	}
	// PrivilegeTypes represents a string slice of possible privilege types
	PrivilegeTypes []string = []string{
		PrivilegeTypeApplication,
		PrivilegeTypeContentSelector,
		PrivilegeTypeRepositoryAdmin,
		PrivilegeTypeRepositoryView,
		PrivilegeTypeScript,
		PrivilegeTypeWildcard,
	}
)

// Privilege data
type Privilege struct {
	Actions         []string `json:"actions,omitempty"`
	ContentSelector string   `json:"contentSelector,omitempty"`
	Description     string   `json:"description"`
	Domain          string   `json:"domain,omitempty"`
	Format          string   `json:"format,omitempty"`
	Name            string   `json:"name"`
	Pattern         string   `json:"pattern,omitempty"`
	ScriptName      string   `json:"scriptName,omitempty"`
	ReadOnly        bool     `json:"readOnly"`
	Repository      string   `json:"repository,omitempty"`
	Type            string   `json:"type"`
}

func (c client) Privileges() ([]Privilege, error) {
	body, resp, err := c.Get(privilegeAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read privileges: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var privileges []Privilege
	if err := json.Unmarshal(body, &privileges); err != nil {
		return nil, fmt.Errorf("could not unmarshal privileges: %v", err)
	}

	return privileges, nil
}

func (c client) PrivilegeCreate(p Privilege) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(p)
	if err != nil {
		return err
	}

	body, resp, err := c.Post(fmt.Sprintf("%s/%s", privilegeAPIEndpoint, strings.ToLower(p.Type)), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create privilege \"%s\": HTTP: %d, %s", p.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (c client) PrivilegeRead(name string) (*Privilege, error) {
	privileges, err := c.Privileges()
	if err != nil {
		return nil, err
	}

	for _, p := range privileges {
		if p.Name == name {
			return &p, nil
		}
	}

	return nil, nil
}

func (c client) PrivilegeUpdate(name string, p Privilege) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(p)
	if err != nil {
		return err
	}

	body, resp, err := c.Put(fmt.Sprintf("%s/%s/%s", privilegeAPIEndpoint, p.Type, name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update privilege \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}

func (c client) PrivilegeDelete(name string) error {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", privilegeAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete privilege \"%s\": HTTP: %d, %s", name, resp.StatusCode, string(body))
	}
	return nil
}
