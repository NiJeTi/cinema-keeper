package discordUtils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nijeti/cinema-keeper/internal/pkg/discordUtils"
)

func TestParseCustomID(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		id         string
		command    string
		subCommand string
		err        error
	}{
		"invalid_id": {
			id:  "invalid id",
			err: discordUtils.ErrInvalidCustomID,
		},
		"command_only": {
			id:      "command_args",
			command: "command",
		},
		"command_and_subcommand": {
			id:         "command_sub_args",
			command:    "command",
			subCommand: "sub",
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				command, subCommand, err := discordUtils.ParseCustomID(tt.id)
				assert.Equal(t, tt.command, command)
				assert.Equal(t, tt.subCommand, subCommand)
				assert.ErrorIs(t, err, tt.err)
			},
		)
	}
}
