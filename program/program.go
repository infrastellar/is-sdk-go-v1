package program

import (
	"os"
	"path/filepath"

	"github.com/infrastellar/is-sdk-go-v1/is"
)

const (
	ignored string = `**/.terraform
**/.terraform.lock.hcl
**/*.plan
**/config.s3.tfbackend
**/env.tf
**/env.tfvars.json
**/mission.tfvars
**/mission.tf
`

	EnvironmentsDir string = "environments"
	SpacesDir       string = "spaces"
	MissionsDir     string = "missions"
	RootDir         string = "root"
	ModulesDir      string = "modules"
)

func RetrieveActiveProgram() (*is.Program, error) {
	programPath, ok := os.LookupEnv(is.EnvVarPROGRAM)
	if ok {
		err := is.DirExists(programPath)
		if err != nil {
			return nil, err
		}
	}

	program, err := ReadProgram(programPath)
	if err != nil {
		return nil, err
	}
	return program, nil
}

type ActiveProgram string

func ReadProgram(path string) (*is.Program, error) {
	err := is.DirExists(path)
	if err != nil {
		return nil, err
	}

	name := filepath.Base(path)

	program := &is.Program{
		Name: name,
		Path: path,
	}
	return program, nil
}
