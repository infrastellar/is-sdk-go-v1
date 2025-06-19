package program

import (
	"context"
	"fmt"

	"github.com/infrastellar/is-sdk-go-v1/program"

	"github.com/urfave/cli/v3"
)

var (
	flagName      string
	flagProvider  string
	flagRegion    string
	flagAccountID string

	SubCmdActive = &cli.Command{
		Name:        "active",
		Usage:       "Return the path to the active program",
		Description: "Return the path to the active program",
		Action: func(ctx context.Context, cli *cli.Command) error {
			program, err := program.RetrieveActiveProgram()
			if err != nil {
				return err
			}
			fmt.Println(program.Path)

			return nil
		},
	}

	SubCmdCreate = &cli.Command{
		Name:        "create",
		Usage:       "Create a new program using the default template",
		Description: "Create a new program using the default template",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Usage:       "Name of new program",
				Aliases:     []string{"n"},
				Required:    true,
				Destination: &flagName,
			},
		},
		Action: func(ctx context.Context, cli *cli.Command) error {
			prg, err := program.NewProgramFromTemplate(flagName)
			if err != nil {
				return err
			}

			err = program.RenderProgramManifestToDisk(prg)
			if err != nil {
				return err
			}

			return nil
		},
	}
)
