package environment

import (
	"github.com/google/uuid"
)

func NewEnvironment(e *Environment) (*Environment, error) {
	uuidv7, err := genID()
	if err != nil {
		return e, err
	}

	e.EID = uuidv7

	return e, nil
}

func ReadEnvironmentsFromDisk(path string) ([]*Environment, error) {
	return []*Environment{}, nil
}

func (e *Environment) RenderToJSON() error {
	return nil
}

func (e *Environment) RenderToDisk() error {
	return nil
}

func genID() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
