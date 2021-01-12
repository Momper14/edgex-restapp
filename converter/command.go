package converter

import (
	south "github.com/momper14/edgex-restapp/client/models"
	north "github.com/momper14/edgex-restapp/models"
)

// CommandFrom creates a Command from a client/CoreCommand.
func CommandFrom(from *south.CoreCommand) *north.Command {
	if from == nil {
		return nil
	}

	return &north.Command{
		Name: from.Name,
		Get:  CommandGetFrom(from.Get),
		Put:  CommandPutFrom(from.Put),
	}
}

// CommandsFrom creates a Command array from a client/CoreCommand array.
func CommandsFrom(from []*south.CoreCommand) []*north.Command {
	if from == nil {
		return nil
	}

	var result = make([]*north.Command, len(from))

	for i, c := range from {
		result[i] = CommandFrom(c)
	}

	return result
}
