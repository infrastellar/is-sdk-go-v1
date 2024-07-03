package space

import (
	"github.com/infrastellar/is-sdk-go-v1/program"
	"github.com/infrastellar/is-sdk-go-v1/space"

	"github.com/urfave/cli/v2"
)

var (
	flagSpace     string
	flagRegion    string
	flagAccountID string
	flagDirector  string
	flagProvider  string

	SubCmdNew = &cli.Command{
		Name:        "new",
		Usage:       "XXXXX",
		Description: "XXXXX",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "space",
				Aliases:     []string{"s"},
				Required:    true,
				Destination: &flagSpace,
			},
			&cli.StringFlag{
				Name:        "region",
				Aliases:     []string{"r"},
				Required:    true,
				Destination: &flagRegion,
			},
			&cli.StringFlag{
				Name:        "account-id",
				Aliases:     []string{"a"},
				Required:    true,
				Destination: &flagAccountID,
			},
			&cli.StringFlag{
				Name:        "director",
				Aliases:     []string{"d"},
				Required:    true,
				Destination: &flagDirector,
			},
			&cli.StringFlag{
				Name:        "provider",
				Aliases:     []string{"p"},
				Value:       "aws",
				Destination: &flagProvider,
			},
		},
		Action: func(cCtx *cli.Context) error {
			prg, err := program.RetrieveActiveProgram()
			if err != nil {
				return err
			}

			spc := space.NewSpace(
				flagSpace,
				flagRegion,
				flagAccountID,
				flagDirector,
				flagProvider,
				prg.Root,
			)

			err = spc.RenderToDisk()
			if err != nil {
				return err
			}
			return nil
		},
	}

	SubCmdValidate = &cli.Command{
		Name:        "validate",
		Usage:       "XXXXX",
		Description: "XXXXX",
		Action: func(cCtx *cli.Context) error {
			return nil
		},
	}
)
