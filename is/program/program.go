package program

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/infrastellar/is-sdk-go-v1/is"
	"github.com/infrastellar/is-sdk-go-v1/is/config"
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

func NewProgramFromTemplate(name string) (*is.Program, error) {
	cfg, err := config.ReadConfig()
	if err != nil {
		return nil, err
	}

	ptmpl, ok := cfg.Templates["program"]
	if !ok {
		return nil, fmt.Errorf("unable to find program template")
	}

	pName := fmt.Sprintf("%s-program", strings.ToLower(name))
	pPath, err := filepath.Abs(fmt.Sprintf("./%s", pName))
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("git", "clone", ptmpl, pName)
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	program := &is.Program{
		Name: name,
		Path: pPath,
	}

	return program, nil
}

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

func RenderProgramManifestToDisk(p *is.Program) error {
	cfg, err := config.ReadConfig()
	if err != nil {
		return err
	}
	json, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}

	manifest := fmt.Sprintf("%s.program.json", p.Name)
	prgPath := filepath.Join(cfg.ProgramsDirectory, manifest)
	err = os.WriteFile(prgPath, json, 0o640)
	if err != nil {
		return err
	}

	return nil
}
