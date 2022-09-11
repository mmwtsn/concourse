package commands

import (
	"fmt"
	"strconv"

	"github.com/concourse/concourse/v7/go-concourse/concourse"

	"github.com/concourse/concourse/v7/atc"
	"github.com/concourse/concourse/v7/fly/commands/internal/flaghelpers"
	"github.com/concourse/concourse/v7/fly/rc"
)

type AbortBuildCommand struct {
	Job   flaghelpers.JobFlag  `short:"j" long:"job" value-name:"PIPELINE/JOB"   description:"Name of a job to cancel"`
	Build string               `short:"b" long:"build" required:"true" description:"If job is specified: build number to cancel. If job not specified: build id"`
	Team  flaghelpers.TeamFlag `long:"team" description:"Name of the team to which the pipeline belongs, if different from the target default"`
}

func (command *AbortBuildCommand) Execute([]string) error {
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

	var build atc.Build
	var exists bool
	if command.Job.PipelineRef.Name == "" && command.Job.JobName == "" {
		build, exists, err = target.Client().Build(command.Build)
	} else {
		build, exists, err = team.JobBuild(command.Job.PipelineRef, command.Job.JobName, command.Build)
	}
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("build does not exist")
	}

	if err := target.Client().AbortBuild(strconv.Itoa(build.ID)); err != nil {
		return err
	}

	fmt.Println("build successfully aborted")
	return nil
}
