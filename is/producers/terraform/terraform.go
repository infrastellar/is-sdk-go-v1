// Package terraform
package terraform

import (
	"github.com/infrastellar/is-sdk-go-v1/is"
	"github.com/infrastellar/is-sdk-go-v1/is/space"
)

// Example:
// producer := terraform.NewProducer(
//     WithManifest(manifest),
//     WithMission("test"),
//     WithStages([]string{"stage000"}),
//     WithProcedures([]string{"proc1", "proc2"}),
//     WithCommands([]string{"plan"}),
// )

// allowedCommands are the commands that this producer works with
var allowedCommands = []string{"init", "plan", "apply", "fmt", "validate", "console", "output", "destroy", "manual"}

type Producer struct {
	BaseCommand string
	Commands    []string
	Environment is.Environment
	Space       space.Space
	Mission     is.Mission
	Stage       string
}

func NewProducer(options ...func(*Producer)) *Producer {
	producer := &Producer{
		BaseCommand: "terraform",
	}
	for _, opt := range options {
		opt(producer)
	}
	return producer
}

func (p *Producer) SetupProcedure() error {
	return nil
}

func (p *Producer) RunProcedure() error {
	return nil
}

func (p *Producer) CleanupProcedure() error {
	return nil
}

func WithCommands(cmds []string) func(*Producer) {
	return func(p *Producer) {
		ok := validateCommands(cmds)

		if ok {
			p.Commands = cmds
		}
	}
}

func validateCommands(cmds []string) bool {
	verified := make(map[string]bool, len(allowed))
	for _, valid := range allowedCommands {
		verified[valid] = true
	}
	for _, cmd := range cmds {
		if !verified[cmd] {
			return false
		}
	}
	return true
}
