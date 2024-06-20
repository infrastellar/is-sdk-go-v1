package config

import (
	"github.com/infrastellar/is-sdk-go-v1/config"

	"github.com/urfave/cli/v2"
)

var (
	flagRegion    string
	flagAccountID string
	flagProvider  string

	SubCmdInit = &cli.Command{
		Name:        "init",
		Usage:       "XXXXX",
		Description: "XXXXX",
		Action: func(cCtx *cli.Context) error {
			err := config.Initialize()
			if err != nil {
				return err
			}
			return nil
		},
	}
)
