package converter

import (
	south "github.com/momper14/edgex-restapp/client/models"
	north "github.com/momper14/edgex-restapp/models"
)

// CommandPayloadTo creates a client/CommandPayload from a CommandPayload.
func CommandPayloadTo(from north.CommandPayload) south.CommandPayload {
	if from == nil {
		return nil
	}

	return (south.CommandPayload)(from)
}
