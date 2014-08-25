package command

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/pinterb/hsc/utils"
)

// InitCommand is the Command implementation that initializes the installation of HSC.
type InitCommand struct {
	Utils *utils.Utils
	UI    cli.Ui
}

// Help print help message for the init sub-command.
func (c *InitCommand) Help() string {
	helpText := `
Usage: hsc init [options]

  Initializes the installation of HSC.
  
Options:

  -user=repouser      Your GitHub username.  
  -org=repoowner      The owner of the GitHub repository.  Only required when different than repository username (e.g. you're working with an organization's GitHub repo).  
  -token=repotoken    Either the value of your GitHub API token or "GITHUB_TOKEN" -- which is the env. variable containing your token.  
  -dir=pathinfo       The local directory that will be home to the working copies of your projects (i.e. cloned repositories).  
 
`
	return strings.TrimSpace(helpText)
}

// Run executes the actual sub-command. Specifically, prints the version of HSC.
func (c *InitCommand) Run(args []string) int {
	var projectPath, repoUser, repoOwner, repoApiToken string

	cmdFlags := flag.NewFlagSet("init", flag.ContinueOnError)
	cmdFlags.StringVar(&repoUser, "user", "", "user")
	cmdFlags.StringVar(&repoOwner, "org", "", "org")
	cmdFlags.StringVar(&projectPath, "dir", "", "dir")
	cmdFlags.StringVar(&repoApiToken, "token", "", "token")
	cmdFlags.Usage = func() { c.UI.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	args = cmdFlags.Args()
	if len(args) > 4 {
		c.UI.Error(
			"The init command expects at most four arguments to initialize the HSC installation.")
		cmdFlags.Usage()
		return 1
	}

	if projectPath == "" {
		c.UI.Error(
			"Dir must be entered.")
		cmdFlags.Usage()
		return 1
	} else if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		c.Ui.Error(fmt.Sprintf("Your local project directory '%s' does not exist.  Create the directory and re-try.", projectPath))
		cmdFlags.Usage()
		return 1
	}

	Utils.config

	return 0
}

// Synopsis prints a brief description of the sub-command.
func (c *InitCommand) Synopsis() string {
	return "Initializes the installation of HSC."
}
