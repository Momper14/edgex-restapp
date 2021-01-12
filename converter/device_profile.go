package converter

import (
	south "github.com/momper14/edgex-restapp/client/models"
	north "github.com/momper14/edgex-restapp/models"
)

// DeviceProfileFrom creates a DeviceProfile from a client/DeviceProfile.
func DeviceProfileFrom(from *south.DeviceProfile) *north.DeviceProfile {
	if from == nil {
		return nil
	}

	return &north.DeviceProfile{
		Labels:       from.Labels,
		Description:  from.Description,
		Created:      from.Created,
		Modified:     from.Modified,
		Manufacturer: from.Manufacturer,
		Model:        from.Model,
		Resources:    DeviceResourcesFrom(from.DeviceResources),
		Commands:     CommandsFrom(from.CoreCommands),
	}
}
