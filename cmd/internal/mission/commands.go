package mission

import (
	"context"

	"github.com/infrastellar/is-sdk-go-v1/mission"

	"github.com/urfave/cli/v3"
)

var (
	flagEnvironment string
	flagMission     string

	SubCmdNew = &cli.Command{
		Name:        "new",
		Usage:       "Add a new mission setup",
		Description: "Add a new mission setup",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "mission",
				Aliases:     []string{"m"},
				Required:    true,
				Destination: &flagMission,
			},
		},
		Action: func(ctx context.Context, cli *cli.Command) error {
			m := mission.Mission{Name: flagMission}
			err := m.RenderToDisk()
			if err != nil {
				return err
			}
			return nil
		},
	}

	SubCmdAdd = &cli.Command{
		Name:        "add",
		Usage:       "Add provided mission to the provided environment",
		Description: "Add provided mission to the provided environment",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "mission",
				Aliases:     []string{"m"},
				Required:    true,
				Destination: &flagMission,
			},
			&cli.StringFlag{
				Name:        "environment",
				Aliases:     []string{"e"},
				Required:    true,
				Destination: &flagEnvironment,
			},
		},
		Action: func(ctx context.Context, cli *cli.Command) error {
			return nil
		},
	}

	SubCmdList = &cli.Command{
		Name:        "list",
		Usage:       "List missions for the provided environment",
		Description: "List missions for the provided environment",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "environment",
				Aliases:     []string{"e"},
				Required:    true,
				Destination: &flagEnvironment,
			},
		},
		Action: func(ctx context.Context, cli *cli.Command) error {
			return nil
		},
	}
)
