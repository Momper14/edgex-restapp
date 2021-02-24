package converter

import (
	south "github.com/momper14/edgex-restapp/client/models"
	north "github.com/momper14/edgex-restapp/models"
)

// ReadingFrom creates a Reading from a client/Reading.
func ReadingFrom(from *south.Reading) *north.Reading {
	if from == nil {
		return nil
	}

	return &north.Reading{
		Name:      from.Name,
		Created:   from.Created,
		Modified:  from.Modified,
		Value:     from.Value,
		ValueType: from.ValueType,
		Pushed:    from.Pushed,
	}
}
