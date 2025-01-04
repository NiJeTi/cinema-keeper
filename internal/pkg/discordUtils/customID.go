package discordUtils

import (
	"errors"
	"regexp"
)

const (
	customIDCommandIndex    = 1
	customIDSubCommandIndex = 3
)

var ErrInvalidCustomID = errors.New("invalid custom id")

var customIDRegex = regexp.MustCompile(`^([a-z]+)(_([a-z]+))?_(.*)$`)

func ParseCustomID(
	id string,
) (command string, subCommand string, err error) {
	matches := customIDRegex.FindStringSubmatch(id)
	if matches == nil {
		err = ErrInvalidCustomID
		return
	}

	command = matches[customIDCommandIndex]
	subCommand = matches[customIDSubCommandIndex]

	return command, subCommand, nil
}
