package is

// Region is an important type throughout the Infrastellar universe and both
// maps to provider rregions and infrastellar regions. We use a base type here
// to cover the translation from provider to our regions, but also to store
// common data between them all.
type Region struct {
	// Name represents the typical provider region, for example in AWS: us-east-1
	Name string `json:"name"`
	// Id represents the infrastellar translation of the provider region, for
	// example when the provider is AWS: use1
	ID string `json:"id"`
	// Provider name, for example: "aws" (is all that is supported currently)
	Provider string `json:"provider"`
	// AccountID is the provider account id
	AccountID string `json:"account_id"`
}

// Root represents the very root of the entire program, including the root
// account at the provider. The root account is never intended to house
// infrastructure resources, but is used for resource management, such as
// spaces.
type Root struct {
	Region *Region `json:"region"`
}

type EnvRegion struct {
	Region
	Network     map[string]EnvRegionNetwork `json:"network,omitempty"`
	Arrangement int                         `json:"arrangement"`
	Status      EnvRegionStatus             `json:"status"`
}

type EnvRegionNetwork struct{}

type EnvRegionStatus struct{}

type Environment struct {
	Name     string            `json:"name"`
	Id       string            `json:"id"`
	Tier     string            `json:"tier"`
	Tags     map[string]string `json:"tags"`
	Features map[string]bool   `json:"features"`
}
