package converter

import (
	south "github.com/momper14/edgex-restapp/client/models"
	north "github.com/momper14/edgex-restapp/models"
)

// CommandResponseFrom creates a CommandResponse from a client/CommandResponse.
func CommandResponseFrom(from south.CommandResponse) north.CommandResponse {
	if from == nil {
		return nil
	}

	return (north.CommandResponse)(from)
}
