package space

import (
	"github.com/urfave/cli/v2"
)

var (
	flagRegion    string
	flagAccountID string
	flagProvider  string

	SubCmdNew = &cli.Command{
		Name:        "new",
		Usage:       "XXXXX",
		Description: "XXXXX",
		Action: func(cCtx *cli.Context) error {
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
