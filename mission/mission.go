package mission

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/infrastellar/is-sdk-go-v1/program"
)

const (
	MissionBaseDir string = "missions"
)

type Mission struct {
	Name string
}

func (m *Mission) RenderToDisk() error {
	program, err := program.RetrieveActiveProgram()
	if err != nil {
		return err
	}

	mpath := filepath.Join(program.Path, MissionBaseDir)
	_, err = os.Stat(mpath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(mpath, 0o750)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	nmpath := filepath.Join(mpath, m.Name)
	if _, err = os.Stat(nmpath); os.IsNotExist(err) {
		err = os.MkdirAll(nmpath, 0o750)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Mission with name %s already exists at %s", m.Name, nmpath)
	}

	err = os.Chdir(nmpath)
	if err != nil {
		return err
	}

	for _, dir := range []string{"modules", "stage000", "stage100"} {
		nd := filepath.Join(".", dir)
		err = os.MkdirAll(nd, 0o750)
		if err != nil {
			return err
		}
	}

	mtf, err := os.Create(fmt.Sprintf("%s.tf", m.Name))
	if err != nil {
		return err
	}
	mtf.Close()

	mex, err := os.Create(fmt.Sprintf("%s.tfvars.example", m.Name))
	if err != nil {
		return err
	}
	mex.Close()

	fmt.Printf("Mission '%s' created.\n", m.Name)

	return nil
}
