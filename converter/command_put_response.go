package converter

import (
	south "github.com/momper14/edgex-restapp/client/models"
	north "github.com/momper14/edgex-restapp/models"
)

// CommandPutResponseFrom creates a CommandPutResponse from a client/CoreCommandPutResponse.
func CommandPutResponseFrom(from *south.CoreCommandPutResponse) *north.CommandPutResponse {
	if from == nil {
		return nil
	}

	return &north.CommandPutResponse{
		Code:        from.Code,
		Description: from.Description,
	}
}

// CommandPutResponsesFrom creates a CommandPutResponse array from a client/CoreCommandPutResponse array.
func CommandPutResponsesFrom(from []*south.CoreCommandPutResponse) []*north.CommandPutResponse {
	if from == nil {
		return nil
	}

	var result = make([]*north.CommandPutResponse, len(from))

	for i, r := range from {
		result[i] = CommandPutResponseFrom(r)
	}

	return result
}
