package converter

import (
	south "github.com/momper14/edgex-restapp/client/models"
	north "github.com/momper14/edgex-restapp/models"
)

// CommandPutFrom creates a CommandPut from a client/CoreCommandPut.
func CommandPutFrom(from *south.CoreCommandPut) *north.CommandPut {
	if from == nil {
		return nil
	}

	return &north.CommandPut{
		ParameterNames: from.ParameterNames,
		Responses:      CommandPutResponsesFrom(from.Responses),
	}
}
