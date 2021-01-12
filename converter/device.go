package converter

import (
	south "github.com/momper14/edgex-restapp/client/models"
	north "github.com/momper14/edgex-restapp/models"
)

// DeviceFrom creates a Device from a client/Device.
func DeviceFrom(from *south.Device) *north.Device {
	if from == nil {
		return nil
	}

	return &north.Device{
		OperatingState: from.OperatingState,
		Created:        from.Created,
		Description:    from.Description,
		Labels:         from.Labels,
		LastConnected:  from.LastConnected,
		LastReported:   from.LastReported,
		Modified:       from.Modified,
		Name:           from.Name,
		Profile:        DeviceProfileFrom(from.Profile),
		Location:       from.Location,
	}
}

// DevicesFrom creates a Device array from a client/Device array.
func DevicesFrom(from []*south.Device) []*north.Device {
	if from == nil {
		return nil
	}

	var result = make([]*north.Device, len(from))

	for i, d := range from {
		result[i] = DeviceFrom(d)
	}

	return result
}
