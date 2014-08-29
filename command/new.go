package command

import (
	"flag"

	"strings"

	"github.com/mitchellh/cli"
	"github.com/pinterb/hsc/config"
	"github.com/pinterb/hsc/utils"
)

// NewCommand is the Command implementation that initializes the installation of HSC.
type NewCommand struct {
	Utils *utils.Utils
	UI    cli.Ui
}

// Help print help message for the init sub-command.
func (c *NewCommand) Help() string {
	helpText := `

Usage: hsc new [options]

  Start a new project.

  You can start a new project from either a sample project or by
  forking an existing repository. Sample applications are available 
  in a number of languages.  
  
Options:

  -name=project name             Name of your new project. A short name with no 
                                 white space works best -- as this may also become
                                 the GitHub repository name.

  -desc=project description      A brief description of the project. 

  -type=application type         The type of application you're creating (e.g. web,
                                 command line, or microservice api).
                                 Choose from one of the following: api | cli | web

  -lang=programmng lanague       Choose from one of these programming langauges:
                                 perl | python | ruby | java | golang

  -github                        This starts your project by first creating a GitHub
                                 repository and then cloning from it.  Otherwise
                                 the project is just created locally.

  -team=GitHub team              If creating a GitHub repository and working as part 
                                 of a GitHub organization. This option allows you to
                                 create the repository as part of a team.  
                                 NOTE: You must have the necessary team permissions
                                 in GitHub to perform this task.  

  -fork=repo name                Fork an existing project. 
 
`
	return strings.TrimSpace(helpText)
}

// Run executes the actual sub-command. Specifically, prints the version of HSC.
func (c *NewCommand) Run(args []string) int {
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
func (c *NewCommand) Synopsis() string {
	return "Start a new project."
}
