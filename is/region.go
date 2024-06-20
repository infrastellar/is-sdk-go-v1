package is

import (
	"slices"
	"strings"
)

var (
	// Convert provider regions from their standard name to a shorthand
	// recognized by Infrastellar
	regionToID = map[string]map[string]string{
		"aws": {
			"us-east-1": "use1",
			"us-east-2": "use2",
			"us-west-2": "usw2",
		},
	}

	// Convert provider regions to their standard name from a shorthand
	// recognized by Infrastellar
	regionFromID = map[string]map[string]string{
		"aws": {
			"use1": "us-east-1",
			"use2": "us-east-2",
			"usw2": "us-west-2",
		},
	}
)

func NewRegion(provider, region, acctID string) *Region {
	r := &Region{
		Name: region,
	}

	id, ok := withRegionID(provider, region)
	if ok {
		r.ID = id
	}
	p, ok := withProvider(provider)
	if ok {
		r.Provider = p
	}

	aid, ok := withAccountID(acctID)
	if ok {
		r.AccountID = aid
	}

	return r
}

func providers() (keys []string) {
	for k := range regionToID {
		keys = append(keys, k)
	}
	return keys
}

func withRegionID(provider, region string) (string, bool) {
	rid, ok := ConvertRegionNameToID(provider, region)
	return rid, ok
}

func withProvider(provider string) (string, bool) {
	if slices.Contains(providers(), provider) {
		return provider, true
	}
	return "", false
}

func withAccountID(acctID string) (string, bool) {
	return acctID, true
}

func ConvertRegionNameToID(provider, region string) (id string, verified bool) {
	id, verified = regionToID[provider][region]
	return id, verified
}

func ConvertRegionIDToName(provider, id string) (name string, verified bool) {
	name, verified = regionFromID[provider][id]
	return name, verified
}

func RetrieveRegion(provider, provided string) (region, id string, verified bool) {
	if strings.Contains(provided, "-") {
		id, verified = regionToID[provider][provided]
		region = provided
		return region, id, verified
	}

	if len(provided) == 4 {
		region, verified = regionFromID[provider][provided]
		id = provided
		return region, id, verified
	}

	return "", "", false
}
