// Package region is an important type throughout the Infrastellar universe and both
// maps to provider regions and Infrastellar regions.
package region

import (
	"slices"

	"github.com/infrastellar/is-sdk-go-v1/is"
)

func NewRegion(provider, region, acctID string) *is.Region {
	r := &is.Region{
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
