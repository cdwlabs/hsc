package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
	"github.com/pinterb/hsc/command"
	"github.com/pinterb/hsc/config"
	"github.com/pinterb/hsc/utils"
)

// Commands is the mapping of all the available HSC commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Recovering from loading config: ", error)
		}
	}()

	c, err := config.LoadConfig()
	if err != nil {
		panic(err.Error())
	}

	utils := utils.NewUtils(c)

	/*

		init: First time installing HSC
		new: Start a new project
		fork: Collaborate on an existing project
		issues: Manage project ideas, features, stories, and bugs
		version: Display version of your local HSC install
		check: Check for available updates for HSC
		download: Updates to latest HSC version
	*/

	Commands = map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &command.InitCommand{
				Utils: utils,
				UI:    ui,
			}, nil
		},

		"new": func() (cli.Command, error) {
			return &command.NewCommand{
				Utils: utils,
				UI:    ui,
			}, nil
		},

		"teams": func() (cli.Command, error) {
			return &command.TeamCommand{
				Utils: utils,
				UI:    ui,
			}, nil
		},
		/*
			"fork": func() (cli.Command, error) {
				return &command.ForkCommand{
					Utils: utils,
					UI:    ui,
				}, nil
			},

			"issues": func() (cli.Command, error) {
				return &command.IssuesCommand{
					Utils: utils,
					UI:    ui,
				}, nil
			},

			"check": func() (cli.Command, error) {
				return &command.CheckCommand{
					Utils: utils,
					UI:    ui,
				}, nil
			},

			"download": func() (cli.Command, error) {
				return &command.DownloadCommand{
					Utils: utils,
					UI:    ui,
				}, nil
			},
		*/
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Revision:                   GitCommit,
				Version:                    Version,
				VersionPrerelease:          VersionPrerelease,
				VersionCompatibilityBroken: BreaksCompatibilityWithVersion,
				UI: ui,
			}, nil
		},
	}
}
