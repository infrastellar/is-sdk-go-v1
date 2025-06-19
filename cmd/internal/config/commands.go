// Package config handles the cli configuration
package config

import (
	"context"

	"github.com/infrastellar/is-sdk-go-v1/config"
	"github.com/infrastellar/is-sdk-go-v1/is"

	"github.com/urfave/cli/v3"
)

var (
	flagRegion    string
	flagAccountID string
	flagProvider  string

	SubCmdInit = &cli.Command{
		Name:        "init",
		Usage:       "XXXXX",
		Description: "XXXXX",
		Action: func(ctx context.Context, cli *cli.Command) error {
			err := config.InitializeConfig()
			if err != nil {
				return err
			}
			return nil
		},
	}

	SubCmdRoot = &cli.Command{
		Name:        "root",
		Usage:       "Set the root account configuration",
		Description: "Set the root account configuration",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "region",
				Usage:       "Root account region",
				Required:    true,
				Destination: &flagRegion,
			},
			&cli.StringFlag{
				Name:        "account",
				Usage:       "Root account id",
				Required:    true,
				Destination: &flagAccountID,
			},
			&cli.StringFlag{
				Name:        "provider",
				Usage:       "Root account provider",
				Value:       "aws",
				Destination: &flagProvider,
			},
		},
		Action: func(ctx context.Context, cli *cli.Command) error {
			newRegion := is.NewRegion(flagProvider, flagRegion, flagAccountID)
			newRoot := is.NewRoot(newRegion)
			cfg, err := config.ReadConfig()
			if err != nil {
				return err
			}

			err = cfg.UpdateConfig(
				config.WithRootConfig(newRoot),
			)
			if err != nil {
				return err
			}

			return nil
		},
	}
)
