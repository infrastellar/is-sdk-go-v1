// Package config handles the cli configuration
package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/infrastellar/is-sdk-go-v1/is"
)

const (
	ConfigDirName  string = "infrastellar"
	ConfigFileName string = "infrastellar.json"
)

type ConfigOptionsFunc func(*ConfigOptions) error

type ConfigOptions struct {
	LogsDirectory     string                       `json:"logs"`
	KeyringsDirectory string                       `json:"keyrings"`
	ProgramsDirectory string                       `json:"programs"`
	Templates         map[string]string            `json:"templates,omitempty"`
	Root              *is.Root                     `json:"root,omitempty"`
	Commands          map[string]map[string]string `json:"commands"`

	configPath string
}

func UserConfigPath() string {
	_, ok := os.LookupEnv("XDG_CONFIG_HOME")
	if ok {
		return xdg.ConfigHome
	}
	return filepath.Join(os.Getenv("HOME"), ".config")
}

func ConfigDirectory() string {
	return fmt.Sprintf("%s/%s", UserConfigPath(), ConfigDirName)
}

func ConfigFilePath() string {
	return fmt.Sprintf("%s/%s", ConfigDirectory(), ConfigFileName)
}

func WithLogsDirectory(dir string) ConfigOptionsFunc {
	return func(o *ConfigOptions) error {
		o.LogsDirectory = dir
		return nil
	}
}

func WithKeyringsDirectory(dir string) ConfigOptionsFunc {
	return func(o *ConfigOptions) error {
		o.KeyringsDirectory = dir
		return nil
	}
}

func WithProgramsDirectory(dir string) ConfigOptionsFunc {
	return func(o *ConfigOptions) error {
		o.ProgramsDirectory = dir
		return nil
	}
}

func WithCommand(cmd string, cmdOption map[string]string) ConfigOptionsFunc {
	return func(o *ConfigOptions) error {
		o.Commands[cmd] = cmdOption
		return nil
	}
}

func WithRootConfig(root *is.Root) ConfigOptionsFunc {
	return func(o *ConfigOptions) error {
		o.Root = root
		return nil
	}
}

func WithProgramTemplate(giturl string) ConfigOptionsFunc {
	return func(o *ConfigOptions) error {
		templates := o.Templates
		templates["program"] = giturl
		o.Templates = templates
		return nil
	}
}

func WithDefaults() ConfigOptionsFunc {
	return func(o *ConfigOptions) error {
		templates := map[string]string{
			"program": is.ProgramTemplateURL,
		}
		o.configPath = ConfigFilePath()
		o.Templates = templates
		o.KeyringsDirectory = filepath.Join(ConfigDirectory(), "keyrings")
		o.LogsDirectory = filepath.Join(ConfigDirectory(), "logs")
		o.ProgramsDirectory = filepath.Join(ConfigDirectory(), "programs")
		o.Commands = map[string]map[string]string{}
		return nil
	}
}

func (cfg *ConfigOptions) UpdateConfig(optFns ...ConfigOptionsFunc) error {
	for _, optFn := range optFns {
		if err := optFn(cfg); err != nil {
			return err
		}
	}

	err := cfg.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func InitializeConfig(optFns ...ConfigOptionsFunc) error {
	var options ConfigOptions

	defaults := []func(*ConfigOptions) error{
		WithDefaults(),
	}

	// Default configuration options
	for _, d := range defaults {
		if err := d(&options); err != nil {
			return err
		}
	}

	for _, optFn := range optFns {
		if err := optFn(&options); err != nil {
			return err
		}
	}

	for _, dir := range []string{
		options.LogsDirectory,
		options.KeyringsDirectory,
		options.ProgramsDirectory,
	} {
		err := os.MkdirAll(dir, 0o750)
		if err != nil {
			return err
		}
	}

	err := options.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}

func ReadConfig() (*ConfigOptions, error) {
	var o ConfigOptions

	cf := ConfigFilePath()
	jf, err := os.Open(cf)
	if err != nil {
		return &o, err
	}
	defer jf.Close()

	jsonb, err := io.ReadAll(jf)
	if err != nil {
		return &o, err
	}
	err = json.Unmarshal(jsonb, &o)
	if err != nil {
		return &o, err
	}

	return &o, nil
}

func (cfg *ConfigOptions) WriteConfig() error {
	json, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(ConfigFilePath(), json, 0o640)
	if err != nil {
		return err
	}
	return nil
}
