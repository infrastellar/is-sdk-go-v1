package program

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/infrastellar/is-sdk-go-v1/config"
	"github.com/infrastellar/is-sdk-go-v1/is"
)

const (
	ActiveProgramFileName = "PROGRAM"

	ignored string = `**/.terraform
**/.terraform.lock.hcl
**/*.plan
**/config.s3.tfbackend
**/env.tf
**/env.tfvars.json
**/mission.tfvars
**/mission.tf
`
)

var ActivePath string

// Program represents a program object and data we need to interact with it
type Program struct {
	Name string   `json:"name"`
	Root *is.Root `json:"root,omitempty"`
	Path string   `json:"path,omitempty"`
}

// BeforeProgram is used prior to any actions against the program to ensure the
// active program is being worked on
func BeforeProgram() error {
	program, err := RetrieveActiveProgram()
	if err != nil {
		return err
	}

	err = os.Chdir(program)
	if err != nil {
		return err
	}

	ActivePath = program

	return nil
}

func RetrieveActiveProgram() (program string, err error) {
	program, ok := os.LookupEnv(is.EnvVarPROGRAM)
	if !ok {
		cfgdir := config.ConfigDirectory()
		pf, err := os.ReadFile(filepath.Join(cfgdir, ActiveProgramFileName))
		if errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("active program not found, variable %s not set", is.EnvVarPROGRAM)
		}
		program = string(pf[:])
	}

	_, err = os.Stat(program)
	if err != nil {
		return program, err
	}

	return program, nil
}

func NewProgram(name, provider, region, acctID string) *Program {
	nr := is.NewRegion(provider, region, acctID)
	return &Program{
		Name: name,
		Root: &is.Root{
			Region: nr,
		},
	}
}

func SetProgram(path string) error {
	cfgdir := config.ConfigDirectory()
	pf := filepath.Join(cfgdir, ActiveProgramFileName)
	err := os.WriteFile(pf, []byte(path), 0o640)
	if err != nil {
		return err
	}

	return nil
}

func UnsetProgram() error {
	cfgdir := config.ConfigDirectory()
	pf := filepath.Join(cfgdir, ActiveProgramFileName)
	err := os.RemoveAll(pf)
	if err != nil {
		return err
	}
	return nil
}

func (p *Program) RenderToDisk() error {
	prgpath := filepath.Join(".", p.Name)

	if _, err := os.Stat(prgpath); os.IsNotExist(err) {
		err = os.MkdirAll(prgpath, 0o750)
		if err != nil {
			return err
		}
	}

	err := os.Chdir(prgpath)
	if err != nil {
		return err
	}

	fullpath, err := os.Getwd()
	if err != nil {
		return err
	}

	p.Path = fullpath

	for _, dir := range []string{"environments", "spaces", "missions", "root", "modules"} {
		nd := filepath.Join(".", dir)
		err = os.MkdirAll(nd, 0o750)
		if err != nil {
			return err
		}
	}

	readme := []byte(fmt.Sprintf("# Infrastellar Space Program: %s\n", p.Name))
	err = os.WriteFile("README.md", readme, 0o640)
	if err != nil {
		return err
	}

	err = os.WriteFile(".gitignore", []byte(ignored), 0o640)
	if err != nil {
		return err
	}

	cfg, err := config.Read()
	if err != nil {
		return err
	}

	prgcf := filepath.Join(cfg.ProgramsDirectory, fmt.Sprintf("%s.json", p.Name))
	json, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(prgcf, json, 0o640)
	if err != nil {
		return err
	}

	fmt.Println("::: Infrastellar Space Program created.")
	fmt.Println("")
	fmt.Println("... In order to start using this program, set the following:")
	fmt.Printf("\texport %s=%s\n", is.EnvVarPROGRAM, p.Path)

	return nil
}
