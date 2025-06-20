package space

import (
	"context"
	"fmt"

	"github.com/infrastellar/is-sdk-go-v1/is/program"

	"github.com/urfave/cli/v3"
)

var (
	flagSpace     string
	flagRegion    string
	flagAccountID string
	flagDirector  string
	flagProvider  string

	SubCmdNew = &cli.Command{
		Name:  "new",
		Usage: "Create a new space for storing configuration",
		Description: `Create a new space for storing configuration.

Uses the provided configuration to create the space. If a region,
account-id, and provider are not specified the Root configuration
from the configuration is used.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "space",
				Aliases:     []string{"s"},
				Required:    true,
				Destination: &flagSpace,
			},
			&cli.StringFlag{
				Name:        "director",
				Aliases:     []string{"d"},
				Required:    true,
				Destination: &flagDirector,
			},
			&cli.StringFlag{
				Name:        "region",
				Aliases:     []string{"r"},
				Destination: &flagRegion,
			},
			&cli.StringFlag{
				Name:        "account-id",
				Aliases:     []string{"a"},
				Destination: &flagAccountID,
			},
			&cli.StringFlag{
				Name:        "provider",
				Aliases:     []string{"p"},
				Value:       "aws",
				Destination: &flagProvider,
			},
		},
		Action: func(ctx context.Context, cli *cli.Command) error {
			prg, err := program.RetrieveActiveProgram()
			if err != nil {
				return err
			}
			fmt.Println(prg.Path)
			return nil
		},
	}
)
