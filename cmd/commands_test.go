package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmdAndCmdHandlerVars(t *testing.T) {
	assert.Equal(t, len(commands), len(commandHandlers))
	for i := 0; i < len(commands); i++ {
		cmdName := commands[i].Name
		if _, ok := commandHandlers[cmdName]; !ok {
			err := fmt.Errorf(
				"command %v needs to be in both commands and command handlers", cmdName,
			)
			assert.Fail(t, err.Error())
		}
	}
}