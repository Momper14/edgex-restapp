package converter

import (
	south "github.com/momper14/edgex-restapp/client/models"
	north "github.com/momper14/edgex-restapp/models"
)

// DeviceResourceFrom creates a Resource from a client/DeviceResource.
func DeviceResourceFrom(from *south.DeviceResource) *north.Resource {
	if from == nil {
		return nil
	}

	return &north.Resource{
		Name:        from.Name,
		Description: from.Description,
		Properties:  ResourcePropertiesFrom(from.Properties),
	}
}

// DeviceResourcesFrom creates a Resource array from a client/DeviceResource array.
func DeviceResourcesFrom(from []*south.DeviceResource) []*north.Resource {
	if from == nil {
		return nil
	}

	var result = make([]*north.Resource, len(from))

	for i, r := range from {
		result[i] = DeviceResourceFrom(r)
	}

	return result
}
