package command

import (
	"bytes"
	"fmt"

	"github.com/mitchellh/cli"
)

// VersionCommand is the Command implementation that prints the version of HSC.
type VersionCommand struct {
	Revision                   string
	Version                    string
	VersionPrerelease          string
	VersionCompatibilityBroken string
	UI                         cli.Ui
}

// Help normally prints the help for this sub-command.  But it's not necessary here.
func (c *VersionCommand) Help() string {
	return ""
}

// Run executes the actual sub-command. Specifically, prints the version of HSC.
func (c *VersionCommand) Run(_ []string) int {
	var versionString bytes.Buffer
	fmt.Fprintf(&versionString, "HSC v%s", c.Version)
	if c.VersionPrerelease != "" {
		fmt.Fprintf(&versionString, ".%s", c.VersionPrerelease)
		if c.Revision != "" {
			fmt.Fprintf(&versionString, " (%s)", c.Revision)
		}
	}

	c.UI.Output(versionString.String())

	if c.VersionCompatibilityBroken != "" {
		c.UI.Output(fmt.Sprintf("Note: breaks compatibility with v%s", c.VersionCompatibilityBroken))
	}

	return 0
}

// Synopsis prints a brief description of the sub-command.
func (c *VersionCommand) Synopsis() string {
	return "Prints the HSC version."
}
