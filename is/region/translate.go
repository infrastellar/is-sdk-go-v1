package region

import "strings"

var (
	// Convert provider regions from their standard name to a shorthand
	// recognized by Infrastellar
	regionToID = map[string]map[string]string{
		"aws": {
			"us-east-1": "use1",
			"us-east-2": "use2",
			"us-west-2": "usw2",
			"eu-west-1": "euw1",
			"eu-west-2": "euw2",
		},
	}

	// Convert provider regions to their standard name from a shorthand
	// recognized by Infrastellar
	regionFromID = map[string]map[string]string{
		"aws": {
			"use1": "us-east-1",
			"use2": "us-east-2",
			"usw2": "us-west-2",
			"euw1": "eu-west-1",
			"euw2": "eu-west-2",
		},
	}
)

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
