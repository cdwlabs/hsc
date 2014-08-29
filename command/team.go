package command

import (
	"flag"

	"strings"

	"github.com/mitchellh/cli"
	"github.com/pinterb/hsc/config"
	"github.com/pinterb/hsc/utils"
)

// TeamCommand is the Command implementation that initializes the installation of HSC.
type TeamCommand struct {
	Utils *utils.Utils
	UI    cli.Ui
}

// Help print help message for the init sub-command.
func (c *TeamCommand) Help() string {
	helpText := `

Usage: hsc new [options]

  Hack some code in a team setting.

Options:

  -list                          List teams within your organization.

  -members=team name             List members of a team within your organization.

  -repos=team name               List repositories of a team within your organization.

  -create=team name              Create a team within your organization.

  -join=team name                Join a team within your organization.

  -leave=team name               Leave a team within your organization.
 
`
	return strings.TrimSpace(helpText)
}

// Run executes the actual sub-command. Specifically, prints the version of HSC.
func (c *TeamCommand) Run(args []string) int {
	var projectPath, repoUser, repoOwner, repoApiToken string
	var force bool

	cmdFlags := flag.NewFlagSet("init", flag.ContinueOnError)
	cmdFlags.StringVar(&repoUser, "user", "", "user")
	cmdFlags.StringVar(&repoOwner, "org", "", "org")
	cmdFlags.StringVar(&projectPath, "dir", "", "dir")
	cmdFlags.StringVar(&repoApiToken, "token", "", "token")
	cmdFlags.BoolVar(&force, "force", false, "force")
	cmdFlags.Usage = func() { c.UI.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	args = cmdFlags.Args()
	if len(args) > 5 {
		c.UI.Error(
			"The init command expects at most five arguments to initialize the HSC installation.")
		cmdFlags.Usage()
		return 1
	}

	if !force {
		oldConfig, err := config.LoadConfig()
		if err != nil {
			c.UI.Error(err.Error())
			return 1
		}

		if oldConfig != nil {
			c.UI.Error("HSC has already been initialized.  Use the -force flag if you'd like to re-initialize.\n")
			cmdFlags.Usage()
			return 1
		}
	}

	newConfig := &config.Config{
		User:  repoUser,
		Org:   repoOwner,
		Token: repoApiToken,
		Dir:   projectPath,
	}

	if err := newConfig.Validate(); err != nil {
		c.UI.Error(err.Error() + "\n")
		cmdFlags.Usage()
		return 1
	}

	if err := newConfig.Write(); err != nil {
		c.UI.Error(err.Error() + "\n")
		cmdFlags.Usage()
		return 1
	}

	return 0
}

// Synopsis prints a brief description of the sub-command.
func (c *TeamCommand) Synopsis() string {
	return "Hack some code in a team setting."
}
