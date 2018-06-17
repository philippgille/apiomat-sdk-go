package dto

// Backend struct generated with gojson.
//
// "Backend" is also called "Application" in ApiOmat.
// Field "Configuration" was changed manually.
type Backend struct {
	AdminName    string `json:"adminName"`
	AnalyticsIds struct {
		Live    string `json:"LIVE"`
		Staging string `json:"STAGING"`
		Test    string `json:"TEST"`
	} `json:"analyticsIds"`
	APIKeys struct {
		LiveAPIKey    string `json:"liveApiKey"`
		StagingAPIKey string `json:"stagingApiKey"`
		TestAPIKey    string `json:"testApiKey"`
	} `json:"apiKeys"`
	ApplicationDisplayName string `json:"applicationDisplayName"`
	ApplicationName        string `json:"applicationName"`
	ApplicationStatus      struct {
		Live    string `json:"LIVE"`
		Staging string `json:"STAGING"`
		Test    string `json:"TEST"`
	} `json:"applicationStatus"`
	AutoUpgradePlan bool `json:"autoUpgradePlan"`
	Configuration   struct {
		LiveConfig    map[string]map[string]string `json:"liveConfig"`
		StagingConfig map[string]map[string]string `json:"stagingConfig"`
		TestConfig    map[string]map[string]string `json:"testConfig"`
	} `json:"configuration"`
	CreatedAt           int64  `json:"createdAt"`
	CustomerHref        string `json:"customerHref"`
	Description         string `json:"description"`
	Href                string `json:"href"`
	LastDeployCustomers struct {
		Live    string `json:"LIVE"`
		Staging string `json:"STAGING"`
		Test    string `json:"TEST"`
	} `json:"lastDeployCustomers"`
	LastDeployTimes struct {
		Live    int64 `json:"LIVE"`
		Staging int64 `json:"STAGING"`
		Test    int64 `json:"TEST"`
	} `json:"lastDeployTimes"`
	LastModifiedAt             int64  `json:"lastModifiedAt"`
	LiveClusterBaseURL         string `json:"liveClusterBaseURL"`
	Name                       string `json:"name"`
	SelectedPlanHref           string `json:"selectedPlanHref"`
	SelectedPlanJSON           string `json:"selectedPlanJSON"`
	StagingClusterBaseURL      string `json:"stagingClusterBaseURL"`
	TestClusterBaseURL         string `json:"testClusterBaseURL"`
	UserRoleToCreateRessources string `json:"userRoleToCreateRessources"`
	UsesModulesHref            string `json:"usesModulesHref"`
	UsesUisHref                string `json:"usesUisHref"`
}
