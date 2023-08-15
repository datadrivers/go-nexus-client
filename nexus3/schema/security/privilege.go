package security

const (
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

	// See https://help.sonatype.com/repomanager3/nexus-repository-administration/access-control/privileges#Privileges-PrivilegeActions for reference
	ActionCreate = "CREATE"
	ActionAdd    = "ADD"
	ActionRead   = "READ"
	ActionUpdate = "UPDATE"
	ActionDelete = "DELETE"
	ActionBrowse = "BROWSE"
	ActionRun    = "RUN"
	ActionEdit   = "EDIT"

	SecurityPrivilegeApplicationActionCreate SecurityPrivilegeApplicationActions = ActionCreate
	SecurityPrivilegeApplicationActionRead   SecurityPrivilegeApplicationActions = ActionRead
	SecurityPrivilegeApplicationActionUpdate SecurityPrivilegeApplicationActions = ActionUpdate
	SecurityPrivilegeApplicationActionDelete SecurityPrivilegeApplicationActions = ActionDelete

	SecurityPrivilegeRepositoryAdminBrowse SecurityPrivilegeRepositoryAdminActions = ActionBrowse
	SecurityPrivilegeRepositoryAdminRead   SecurityPrivilegeRepositoryAdminActions = ActionRead
	SecurityPrivilegeRepositoryAdminEdit   SecurityPrivilegeRepositoryAdminActions = ActionEdit
	SecurityPrivilegeRepositoryAdminAdd    SecurityPrivilegeRepositoryAdminActions = ActionAdd
	SecurityPrivilegeRepositoryAdminDelete SecurityPrivilegeRepositoryAdminActions = ActionDelete

	SecurityPrivilegeRepositoryContentSelectorBrowse SecurityPrivilegeRepositoryContentSelectorActions = ActionBrowse
	SecurityPrivilegeRepositoryContentSelectorRead   SecurityPrivilegeRepositoryContentSelectorActions = ActionRead
	SecurityPrivilegeRepositoryContentSelectorEdit   SecurityPrivilegeRepositoryContentSelectorActions = ActionEdit
	SecurityPrivilegeRepositoryContentSelectorAdd    SecurityPrivilegeRepositoryContentSelectorActions = ActionAdd
	SecurityPrivilegeRepositoryContentSelectorDelete SecurityPrivilegeRepositoryContentSelectorActions = ActionDelete

	SecurityPrivilegeRepositoryViewBrowse SecurityPrivilegeRepositoryViewActions = ActionBrowse
	SecurityPrivilegeRepositoryViewRead   SecurityPrivilegeRepositoryViewActions = ActionRead
	SecurityPrivilegeRepositoryViewEdit   SecurityPrivilegeRepositoryViewActions = ActionEdit
	SecurityPrivilegeRepositoryViewAdd    SecurityPrivilegeRepositoryViewActions = ActionAdd
	SecurityPrivilegeRepositoryViewDelete SecurityPrivilegeRepositoryViewActions = ActionDelete

	SecurityPrivilegeScriptAdd    SecurityPrivilegeScriptActions = ActionAdd
	SecurityPrivilegeScriptRead   SecurityPrivilegeScriptActions = ActionRead
	SecurityPrivilegeScriptDelete SecurityPrivilegeScriptActions = ActionDelete
	SecurityPrivilegeScriptBrowse SecurityPrivilegeScriptActions = ActionBrowse
	SecurityPrivilegeScriptRun    SecurityPrivilegeScriptActions = ActionRun
	SecurityPrivilegeScriptEdit   SecurityPrivilegeScriptActions = ActionEdit
)

type SecurityPrivilegeApplicationActions string
type SecurityPrivilegeRepositoryAdminActions string
type SecurityPrivilegeRepositoryContentSelectorActions string
type SecurityPrivilegeRepositoryViewActions string
type SecurityPrivilegeScriptActions string

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

// ToDo: Refactor privilege structs?
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

type PrivilegeScript struct {
	Name        string                           `json:"name"`
	Description string                           `json:"description"`
	Actions     []SecurityPrivilegeScriptActions `json:"actions,omitempty"`
	ScriptName  string                           `json:"scriptName,omitempty"`
}

type PrivilegeRepositoryView struct {
	Name        string                                   `json:"name"`
	Description string                                   `json:"description,omitempty"`
	Actions     []SecurityPrivilegeRepositoryViewActions `json:"actions"`
	Format      string                                   `json:"format"`
	Repository  string                                   `json:"repository"`
}

type PrivilegeWildcard struct {
	Name        string `json:"name"`
	Pattern     string `json:"pattern"`
	Description string `json:"description,omitempty"`
}

type PrivilegeRepositoryAdmin struct {
	Name        string                                    `json:"name"`
	Description string                                    `json:"description,omitempty"`
	Actions     []SecurityPrivilegeRepositoryAdminActions `json:"actions"`
	Format      string                                    `json:"format"`
	Repository  string                                    `json:"repository"`
}

type PrivilegeRepositoryContentSelector struct {
	Name            string                                              `json:"name"`
	Description     string                                              `json:"description,omitempty"`
	Actions         []SecurityPrivilegeRepositoryContentSelectorActions `json:"actions"`
	Format          string                                              `json:"format"`
	Repository      string                                              `json:"repository"`
	ContentSelector string                                              `json:"contentSelector"`
}

type PrivilegeApplication struct {
	Name        string                                `json:"name"`
	Description string                                `json:"description,omitempty"`
	Actions     []SecurityPrivilegeApplicationActions `json:"actions"`
	Domain      string                                `json:"domain"`
}
