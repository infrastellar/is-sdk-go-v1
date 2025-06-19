package mission

import (
	"context"

	"github.com/infrastellar/is-sdk-go-v1/mission"

	"github.com/urfave/cli/v3"
)

var (
	flagMission string

	SubCmdNew = &cli.Command{
		Name:        "new",
		Usage:       "XXXXX",
		Description: "XXXXX",
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
		Usage:       "XXXXX",
		Description: "XXXXX",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "mission",
				Aliases:     []string{"m"},
				Required:    true,
				Destination: &flagMission,
			},
		},
		Action: func(ctx context.Context, cli *cli.Command) error {
			return nil
		},
	}

	SubCmdList = &cli.Command{
		Name:        "list",
		Usage:       "XXXXX",
		Description: "XXXXX",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "mission",
				Aliases:     []string{"m"},
				Required:    true,
				Destination: &flagMission,
			},
		},
		Action: func(ctx context.Context, cli *cli.Command) error {
			return nil
		},
	}
)
