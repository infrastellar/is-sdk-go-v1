// Package root
package root

import "github.com/infrastellar/is-sdk-go-v1/is"

func NewRoot(region *is.Region) *is.Root {
	return &is.Root{
		Region: region,
	}
}
