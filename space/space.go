package space

import (
	"github.com/infrastellar/is-sdk-go-v1/is"
)

type Space struct {
	Name     string            `json:"name"`
	Config   *is.BackendConfig `json:"config"`
	Director string            `json:"director,omitempty"`
}

func NewSpace(name, region, acctID, director string) *Space {
	return &Space{
		Name:     name,
		Director: director,
	}
}

func (s *Space) RenderToDisk() error {
	return nil
}
