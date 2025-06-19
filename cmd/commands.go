package iscmd

import (
	"github.com/infrastellar/is-sdk-go-v1/cmd/internal/config"
	"github.com/infrastellar/is-sdk-go-v1/cmd/internal/mission"
	"github.com/infrastellar/is-sdk-go-v1/cmd/internal/program"
	"github.com/infrastellar/is-sdk-go-v1/cmd/internal/space"

	"github.com/urfave/cli/v3"
)

var (
	Config = &cli.Command{
		Name:        "config",
		Usage:       "TODO",
		Description: "TODO",
		Commands: []*cli.Command{
			config.SubCmdInit,
			config.SubCmdRoot,
		},
	}

	Program = &cli.Command{
		Name:        "program",
		Usage:       "TODO",
		Description: "TODO",
		Commands: []*cli.Command{
			program.SubCmdCreate,
			program.SubCmdActive,
		},
	}

	Space = &cli.Command{
		Name:        "space",
		Usage:       "TODO",
		Description: "TODO",
		Commands: []*cli.Command{
			space.SubCmdNew,
		},
	}

	Mission = &cli.Command{
		Name:        "mission",
		Usage:       "TODO",
		Description: "TODO",
		Commands: []*cli.Command{
			mission.SubCmdNew,
			mission.SubCmdAdd,
			mission.SubCmdList,
		},
	}
)
