package commands

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/concourse/concourse/v7/atc"
	"github.com/concourse/concourse/v7/fly/commands/internal/flaghelpers"
	"github.com/concourse/concourse/v7/fly/eventstream"
	"github.com/concourse/concourse/v7/fly/rc"
	"github.com/concourse/concourse/v7/go-concourse/concourse"
)

type WatchCommand struct {
	Job                      flaghelpers.JobFlag  `short:"j" long:"job"         value-name:"PIPELINE/JOB"  description:"Watches builds of the given job"`
	Build                    string               `short:"b" long:"build"                                  description:"Watches a specific build"`
	Url                      string               `short:"u" long:"url"                                    description:"URL for the build or job to watch"`
	Timestamp                bool                 `short:"t" long:"timestamps"                             description:"Print with local timestamp"`
	IgnoreEventParsingErrors bool                 `long:"ignore-event-parsing-errors"                      description:"Ignore event parsing errors"`
	Team                     flaghelpers.TeamFlag `long:"team" description:"Name of the team to which the pipeline belongs, if different from the target default"`
}

func getBuildIDFromURL(target rc.Target, urlParam string, team concourse.Team) (int, error) {
	var buildId int
	client := target.Client()

	u, err := url.Parse(urlParam)
	if err != nil {
		return 0, err
	}

	urlMap := parseUrlPath(u.Path)

	parsedTargetUrl := url.URL{
		Scheme: u.Scheme,
		Host:   u.Host,
	}

	host := parsedTargetUrl.String()
	if host != target.URL() {
		err = fmt.Errorf("URL doesn't match target (%s, %s)", urlParam, target.URL())
		return 0, err
	}

	teamName := urlMap["teams"]
	if teamName != "" && teamName != team.Name() {
		err = fmt.Errorf("Team in URL doesn't match the provided team name or current team of the target  (%s, %s)", urlParam, team.Name())
		return 0, err
	}

	if urlMap["pipelines"] != "" && urlMap["jobs"] != "" {
		pipelineRef := atc.PipelineRef{Name: urlMap["pipelines"]}
		pipelineRef.InstanceVars, err = atc.InstanceVarsFromQueryParams(u.Query())
		if err != nil {
			return 0, err
		}
		build, err := GetBuild(client, team, urlMap["jobs"], urlMap["builds"], pipelineRef)

		if err != nil {
			return 0, err
		}
		buildId = build.ID
	} else if urlMap["builds"] != "" {
		buildId, err = strconv.Atoi(urlMap["builds"])

		if err != nil {
			return 0, err
		}
	} else {
		return 0, fmt.Errorf("No build found in %s", urlParam)
	}
	return buildId, nil
}

func (command *WatchCommand) Execute(args []string) error {
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

	var buildId int
	client := target.Client()
	if command.Job.JobName != "" || command.Build == "" && command.Url == "" {
		build, err := GetBuild(client, team, command.Job.JobName, command.Build, command.Job.PipelineRef)
		if err != nil {
			return err
		}
		buildId = build.ID
	} else if command.Build != "" {
		buildId, err = strconv.Atoi(command.Build)

		if err != nil {
			return err
		}
	} else if command.Url != "" {
		buildId, err = getBuildIDFromURL(target, command.Url, team)

		if err != nil {
			return err
		}
	}

	eventSource, err := client.BuildEvents(fmt.Sprintf("%d", buildId))
	if err != nil {
		return err
	}

	renderOptions := eventstream.RenderOptions{
		ShowTimestamp:            command.Timestamp,
		IgnoreEventParsingErrors: command.IgnoreEventParsingErrors,
	}

	exitCode := eventstream.Render(os.Stdout, eventSource, renderOptions)

	eventSource.Close()

	os.Exit(exitCode)

	return nil
}
