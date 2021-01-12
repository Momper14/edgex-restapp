package converter

import (
	south "github.com/momper14/edgex-restapp/client/models"
	north "github.com/momper14/edgex-restapp/models"
)

// CommandGetFrom creates a CommandGet from a client/CoreCommandGet.
func CommandGetFrom(from *south.CoreCommandGet) *north.CommandGet {
	if from == nil {
		return nil
	}

	return &north.CommandGet{
		Responses: CommandGetResponsesFrom(from.Responses),
	}
}
