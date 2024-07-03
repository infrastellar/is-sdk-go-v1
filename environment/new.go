package environment

import (
	"github.com/infrastellar/is-sdk-go-v1/is"

	"github.com/google/uuid"
)

func NewEnvironment(e *Environment) *Environment {
	e.EID = genID()
	return e
}

func (e *Environment) RenderToJSON() error {
	return nil
}

func (e *Environment) RenderToDisk() error {
	return nil
}

func genID() string {
	id := uuid.NewV7()
	return id.String()
}
