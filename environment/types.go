package environment

import (
	"github.com/infrastellar/is-sdk-go-v1/is"
)

type Environments []Environment

type Environment struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	// EID represents a unique identifier (UUIDv7) for the environment generated at
	// creation.
	EID  string `json:"eid"`
	Tier string `json:"tier"`
	// Audience represents whether those accessing the environment are
	// "internal" or "external" and allow other resources to react accordingly
	Audience string            `json:"audience"`
	Tags     map[string]string `json:"tags"`
	Features map[string]string `json:"features"`
	Status   struct {
		// Condition represents the state of a region, whether in "draft" or
		// "published". Again this is similar to audience in that it allows
		// downstream resources to behave accordingly.
		Condition string `json:"condition"`
	} `json:"status"`
}

type EnvironmentRegion struct {
	is.Region
	// Arrangement represents the order in which regions are added to the
	// environment. We keep track of this to avoid conflicts down the line when
	// running procedures.
	Arrangement int `json:"arrangement"`
	// Network represents a network that has a name and a CIDR. These are
	// typically very static structures, and important to the configuration of
	// an environment region so we manage these here.
	Network map[string]EnvironmentRegionNetwork `json:"network,omitempty"`
	Status  struct {
		Enabled bool `json:"enabled"`
		// Designation typically refers to the behavior of the region,
		// typically this means "publisher" or "subscriber"
		Designation string `json:"designation"`
	} `json:"status"`
}

type EnvironmentRegionNetwork struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	CIDR        string            `json:"cidr"`
	Features    map[string]string `json:"features"`
}
