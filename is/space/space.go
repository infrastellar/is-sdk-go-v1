package space

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/infrastellar/is-sdk-go-v1/is"
	"github.com/infrastellar/is-sdk-go-v1/is/program"
)

const (
	BaseStationDir string = "base-station"
)

type Space struct {
	Name     string        `json:"name"`
	Config   BackendConfig `json:"config"`
	Director string        `json:"director"`
	Region   *is.Region    `json:"region"`
	Root     *is.Root      `json:"root"`
}

type BackendConfig struct {
	Bucket        string `json:"bucket"`
	DynamodbTable string `json:"dynamodb_table"`
	Encrypt       bool   `json:"encrypt"`
	KmsKeyID      string `json:"kms_key_id"`
	Region        string `json:"region"`
	Key           string `json:"key,omitempty"`
	RoleArn       string `json:"role_arn,omitempty"`
}

type BaseStation struct {
	SpaceName      string
	SpaceRegion    string
	SpaceAccountID string
	RootAccountID  string
	RootRegion     string
	Director       string
	DirectorArn    string
}

func NewSpace(name, region, acctID, director, provider string, root *is.Root) *Space {
	sr := is.NewRegion(provider, region, acctID)
	return &Space{
		Name:     name,
		Director: director,
		Region:   sr,
		Root:     root,
	}
}

func (s *Space) RenderToDisk() error {
	spcsdir := filepath.Join(".", program.SpacesDir)
	_, err := os.Stat(spcsdir)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("directory %s doesn't exist: %v", spcsdir, err)
		} else {
			return err
		}
	}

	fmt.Printf("::: Creating new space: %s\n", s.Name)
	newspcdir := filepath.Join(spcsdir, s.Name)
	newspcmoddir := filepath.Join(spcsdir, "modules")
	basestationdir := filepath.Join(newspcdir, BaseStationDir)

	err = os.MkdirAll(newspcmoddir, 0o750)
	if err != nil {
		return err
	}

	err = os.MkdirAll(basestationdir, 0o750)
	if err != nil {
		return err
	}

	bs := BaseStation{
		SpaceName:      s.Name,
		SpaceRegion:    s.Region.Name,
		SpaceAccountID: s.Region.AccountID,
		RootAccountID:  s.Root.Region.AccountID,
		RootRegion:     s.Root.Region.Name,
		Director:       s.Director,
		DirectorArn:    "",
	}

	t := template.Must(template.New("maintf").Parse(BaseStationMainTfTmpl))
	maintf, err := os.Create(filepath.Join(basestationdir, "main.tf"))
	if err != nil {
		return err
	}
	defer maintf.Close()

	err = t.ExecuteTemplate(maintf, "maintf", bs)
	if err != nil {
		return err
	}

	templates := map[string]string{
		"outputs.tf":   BaseStationOutputsTfTmpl,
		"versions.tf":  BaseStationVersionsTfTmpl,
		"providers.tf": BaseStationProvidersTfTmpl,
	}

	for f, tmpl := range templates {
		if err = os.WriteFile(
			filepath.Join(basestationdir, f),
			[]byte(tmpl),
			0o660,
		); err != nil {
			return err
		}
	}

	return nil
}
