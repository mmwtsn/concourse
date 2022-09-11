package commands

import (
	"fmt"

	"github.com/concourse/concourse/v7/fly/commands/internal/displayhelpers"
	"github.com/concourse/concourse/v7/fly/commands/internal/flaghelpers"
	"github.com/concourse/concourse/v7/fly/rc"
	"github.com/concourse/concourse/v7/go-concourse/concourse"
)

type HidePipelineCommand struct {
	Pipeline flaghelpers.PipelineFlag `short:"p"   long:"pipeline" required:"true" description:"Pipeline to hide"`
	Team     flaghelpers.TeamFlag     `long:"team"                                 description:"Name of the team to which the pipeline belongs, if different from the target default"`
}

func (command *HidePipelineCommand) Validate() error {
	_, err := command.Pipeline.Validate()
	return err
}

func (command *HidePipelineCommand) Execute(args []string) error {
	err := command.Validate()
	if err != nil {
		return err
	}

	target, err := rc.LoadTarget(Fly.Target, Fly.Verbose)
	if err != nil {
		return err
	}

	err = target.Validate()
	if err != nil {
		return err
	}

	var team concourse.Team
	team, err = command.Team.LoadTeam(target)
	if err != nil {
		return err
	}

	pipelineRef := command.Pipeline.Ref()
	found, err := team.HidePipeline(pipelineRef)
	if err != nil {
		return err
	}

	if found {
		fmt.Printf("hid '%s'\n", pipelineRef.String())
	} else {
		displayhelpers.Failf("pipeline '%s' not found\n", pipelineRef.String())
	}

	return nil
}
