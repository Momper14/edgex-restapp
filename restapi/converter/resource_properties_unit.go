package converter

import (
	south "github.com/momper14/edgex-restapp/client/models"
	north "github.com/momper14/edgex-restapp/models"
)

// ResourcePropertiesUnitFrom creates a ResourcePropertiesUnit from a client/Unit.
func ResourcePropertiesUnitFrom(from *south.Unit) *north.ResourcePropertiesUnit {
	if from == nil {
		return nil
	}

	return &north.ResourcePropertiesUnit{
		Type:         from.Type,
		ReadWrite:    from.ReadWrite,
		DefaultValue: from.DefaultValue,
	}
}
