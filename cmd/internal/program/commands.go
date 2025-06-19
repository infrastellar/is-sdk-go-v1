package program

import (
	"context"
	"fmt"

	"github.com/infrastellar/is-sdk-go-v1/program"

	"github.com/urfave/cli/v3"
)

var (
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
)
