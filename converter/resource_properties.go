package converter

import (
	south "github.com/momper14/edgex-restapp/client/models"
	north "github.com/momper14/edgex-restapp/models"
)

// ResourcePropertiesFrom creates a ResourceProperties from a client/Propertie.
func ResourcePropertiesFrom(from *south.Propertie) *north.ResourceProperties {
	if from == nil {
		return nil
	}

	return &north.ResourceProperties{
		Unit:  ResourcePropertiesUnitFrom(from.Units),
		Value: ResourcePropertiesValueFrom(from.Value),
	}
}
