package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestVersionCommand(t *testing.T) {
	var _ cli.Command = &VersionCommand{}
}
