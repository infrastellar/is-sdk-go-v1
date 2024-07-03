package iscmd

import (
	"github.com/infrastellar/is-sdk-go-v1/cmd/internal/config"
	"github.com/infrastellar/is-sdk-go-v1/cmd/internal/program"
	"github.com/infrastellar/is-sdk-go-v1/cmd/internal/space"

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

	Space = &cli.Command{
		Name:        "space",
		Usage:       "TODO",
		Description: "TODO",
		Subcommands: []*cli.Command{
			space.SubCmdNew,
			space.SubCmdValidate,
		},
	}

	Mission = &cli.Command{
		Name:        "mission",
		Usage:       "TODO",
		Description: "TODO",
		Subcommands: []*cli.Command{
			mission.SubCmdNew,
			mission.SubCmdAdd,
			mission.SubCmdList,
		},
	}
)
