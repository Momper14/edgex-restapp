package converter

import (
	south "github.com/momper14/edgex-restapp/client/models"
	north "github.com/momper14/edgex-restapp/models"
)

// CommandGetResponseFrom creates a CommandGetResponse from a client/CoreCommandGetResponse.
func CommandGetResponseFrom(from *south.CoreCommandGetResponse) *north.CommandGetResponse {
	if from == nil {
		return nil
	}

	return &north.CommandGetResponse{
		Code:           from.Code,
		Description:    from.Description,
		ExpectedValues: from.ExpectedValues,
	}
}

// CommandGetResponsesFrom creates a CommandGetResponse array from a client/CoreCommandGetResponse array.
func CommandGetResponsesFrom(from []*south.CoreCommandGetResponse) []*north.CommandGetResponse {
	if from == nil {
		return nil
	}

	var result = make([]*north.CommandGetResponse, len(from))

	for i, r := range from {
		result[i] = CommandGetResponseFrom(r)
	}

	return result
}
