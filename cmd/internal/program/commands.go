package program

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/infrastellar/is-sdk-go-v1/program"

	"github.com/urfave/cli/v2"
)

var (
	flagProvider  string
	flagRegion    string
	flagAccountID string

	SubCmdActive = &cli.Command{
		Name:        "active",
		Usage:       "XXXXX",
		Description: "XXXXX",
		Action: func(cCtx *cli.Context) error {
			pf, err := program.RetrieveActiveProgram()
			if err != nil {
				return err
			}
			fmt.Println(pf.Path)

			return nil
		},
	}

	SubCmdDescribe = &cli.Command{
		Name:        "describe",
		Aliases:     []string{"desc"},
		Usage:       "XXXXX",
		Description: "XXXXX",
		Action: func(cCtx *cli.Context) error {
			pf, err := program.RetrieveActiveProgram()
			if err != nil {
				return err
			}

			b, err := json.MarshalIndent(pf, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(b))

			return nil
		},
	}

	SubCmdNew = &cli.Command{
		Name:        "new",
		Usage:       "XXXXX",
		Description: "XXXXX",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "provider",
				Aliases:     []string{"p"},
				Required:    true,
				Usage:       "XXXXX",
				Destination: &flagProvider,
			},
			&cli.StringFlag{
				Name:        "region",
				Aliases:     []string{"r"},
				Required:    true,
				Usage:       "XXXXX",
				Destination: &flagRegion,
			},
			&cli.StringFlag{
				Name:        "account-id",
				Aliases:     []string{"a"},
				Required:    true,
				Usage:       "XXXXX",
				Destination: &flagAccountID,
			},
		},
		Action: func(cCtx *cli.Context) error {
			name := cCtx.Args().Get(0)

			p := program.NewProgram(name, flagProvider, flagRegion, flagAccountID)
			err := p.RenderToDisk()
			if err != nil {
				return err
			}
			return nil
		},
	}

	SubCmdSet = &cli.Command{
		Name:        "set",
		Usage:       "XXXXX",
		Description: "XXXXX",
		Action: func(cCtx *cli.Context) error {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}

			err = program.SetProgram(cwd)
			if err != nil {
				return err
			}

			return nil
		},
	}

	SubCmdUnSet = &cli.Command{
		Name:        "unset",
		Usage:       "XXXXX",
		Description: "XXXXX",
		Action: func(cCtx *cli.Context) error {
			err := program.UnsetProgram()
			if err != nil {
				return err
			}

			return nil
		},
	}
)
