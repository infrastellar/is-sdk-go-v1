package iscmd

import (
	"github.com/infrastellar/is-sdk-go-v1/cmd/internal/config"
	"github.com/infrastellar/is-sdk-go-v1/cmd/internal/program"

	"github.com/urfave/cli/v2"
)

var (
	Config = &cli.Command{
		Name:        "config",
		Usage:       "TODO",
		Description: "TODO",
		Subcommands: []*cli.Command{
			config.SubCmdInit,
		},
	}

	Program = &cli.Command{
		Name:        "program",
		Usage:       "TODO",
		Description: "TODO",
		Subcommands: []*cli.Command{
			program.SubCmdActive,
			program.SubCmdNew,
			program.SubCmdSet,
			program.SubCmdUnSet,
		},
	}
)
