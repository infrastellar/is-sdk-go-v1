// Package is
package is

// Region We use a base type here to cover the translation from provider to
// our regions, but also to store common data between them all.
type Region struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	Provider  string `json:"provider"`
	AccountID string `json:"account_id"`
}

type RegionStatus struct {
	Enabled     bool   `json:"enabled"`
	Designation string `json:"designation"`
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

type Environments []Environment

// EnvironmentRegionNetwork prepares a network configuration for a region
type EnvironmentRegionNetwork struct {
	Name        string            `json:"name,omitempty"`
	Cidr        string            `json:"cidr"`
	Description string            `json:"description,omitempty"`
	Features    map[string]string `json:"features,omitempty"`
}

// EnvironmentRegionStatus is a container for the dynamic status of an environment
type EnvironmentRegionStatus struct {
	Enabled bool `json:"enabled"`
	// Designation typically refers to the behavior of the region,
	// typical values might be "publisher" or "subscriber"
	Designation string `json:"designation"`
}

type EnvironmentRegion struct {
	Region

	// Arrangement represents the order in which regions are added to the
	// environment. We keep track of this to avoid conflicts down the line when
	// running procedures.
	Arrangement int `json:"arrangement"`
	// Network represents a network that has a name and a CIDR. These are
	// typically very static structures, and important to the configuration of
	// an environment region so we manage these here.
	Network map[string]EnvironmentRegionNetwork `json:"network,omitempty"`
	Status  EnvironmentRegionStatus
}

// Program represents a program object and data we need to interact with it
type Program struct {
	Name    string `json:"name"`
	Path    string `json:"path,omitempty"`
	URL     string `json:"url,omitempty"`
	Version struct {
		Commit string `json:"commit,omitempty"`
	}
	Environments Environments
}
