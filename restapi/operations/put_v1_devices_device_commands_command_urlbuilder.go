// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// PutV1DevicesDeviceCommandsCommandURL generates an URL for the put v1 devices device commands command operation
type PutV1DevicesDeviceCommandsCommandURL struct {
	Command string
	Device  string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PutV1DevicesDeviceCommandsCommandURL) WithBasePath(bp string) *PutV1DevicesDeviceCommandsCommandURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PutV1DevicesDeviceCommandsCommandURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PutV1DevicesDeviceCommandsCommandURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v1/devices/{device}/commands/{command}"

	command := o.Command
	if command != "" {
		_path = strings.Replace(_path, "{command}", command, -1)
	} else {
		return nil, errors.New("command is required on PutV1DevicesDeviceCommandsCommandURL")
	}

	device := o.Device
	if device != "" {
		_path = strings.Replace(_path, "{device}", device, -1)
	} else {
		return nil, errors.New("device is required on PutV1DevicesDeviceCommandsCommandURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PutV1DevicesDeviceCommandsCommandURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PutV1DevicesDeviceCommandsCommandURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PutV1DevicesDeviceCommandsCommandURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PutV1DevicesDeviceCommandsCommandURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PutV1DevicesDeviceCommandsCommandURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *PutV1DevicesDeviceCommandsCommandURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
