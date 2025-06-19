package is

// Region is an important type throughout the Infrastellar universe and both
// maps to provider rregions and infrastellar regions. We use a base type here
// to cover the translation from provider to our regions, but also to store
// common data between them all.
type Region struct {
	// Name represents the typical provider region, for example in AWS: us-east-1
	Name string `json:"name"`
	// Id represents the Infrastellar translation of the provider region, for
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

// EnvironmentStatus is a container for the dynamic status of an environment
type EnvironmentStatus struct {
	Condition string `json:"condition"`
	Recovery  string `json:"recovery,omitempty"`
}

type Environment struct {
	Name     string            `json:"name"`
	ID       string            `json:"id"`
	Tier     string            `json:"tier,omitempty"`
	Audience string            `json:"audience,omitempty"`
	Tags     map[string]string `json:"tags,omitempty"`
	Features map[string]string `json:"features,omitempty"`
	Status   EnvironmentStatus `json:"status,omitempty"`
}

// Program represents a program object and data we need to interact with it
type Program struct {
	Name    string `json:"name"`
	Path    string `json:"path,omitempty"`
	Url     string `json:"url,omitempty"`
	Version struct {
		Commit string `json:"commit,omitempty"`
	}
	Environments []struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"environments,omitempty"`
}
