package converter

import (
	south "github.com/momper14/edgex-restapp/client/models"
	north "github.com/momper14/edgex-restapp/models"
)

// ValueDescriptorFrom creates a ValueDescriptor from a client/ValueDescriptor.
func ValueDescriptorFrom(from *south.ValueDescriptor) *north.ValueDescriptor {
	if from == nil {
		return nil
	}

	return &north.ValueDescriptor{
		Name:         from.Name,
		Labels:       from.Labels,
		Description:  from.Description,
		Created:      from.Created,
		Modified:     from.Modified,
		Type:         from.Type,
		UomLabel:     from.UomLabel,
		Min:          from.Min,
		Max:          from.Max,
		DefaultValue: from.DefaultValue,
		Formatting:   from.Formatting,
	}
}

// ValueDescriptorsFrom creates a ValueDescriptor array from a client/ValueDescriptor array.
func ValueDescriptorsFrom(from []*south.ValueDescriptor) []*north.ValueDescriptor {
	if from == nil {
		return nil
	}

	var result = make([]*north.ValueDescriptor, len(from))

	for i, d := range from {
		result[i] = ValueDescriptorFrom(d)
	}

	return result
}
