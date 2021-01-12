package converter

import (
	south "github.com/momper14/edgex-restapp/client/models"
	north "github.com/momper14/edgex-restapp/models"
)

// ResourcePropertiesValueFrom creates a ResourcePropertiesValue from a client/Value.
func ResourcePropertiesValueFrom(from *south.Value) *north.ResourcePropertiesValue {
	if from == nil {
		return nil
	}

	return &north.ResourcePropertiesValue{
		Type:          from.Type,
		ReadWrite:     from.ReadWrite,
		Minimum:       from.Minimum,
		Maximum:       from.Maximum,
		DefaultValue:  from.DefaultValue,
		Mask:          from.Mask,
		Shift:         from.Shift,
		Scale:         from.Scale,
		Offset:        from.Offset,
		Base:          from.Base,
		Assertion:     from.Assertion,
		FloatEncoding: from.FloatEncoding,
		Mediatype:     from.MediaType,
	}
}
