// Package client is for communication with the EdgeX Foundry platform.
package client

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/momper14/edgex-restapp/client/models"
	"github.com/spf13/viper"
)

var (
	// ClientCoreData is the client for core data.
	ClientCoreData *resty.Client
	// ClientCoreMetaData is the client for core metadata.
	ClientCoreMetaData *resty.Client
	// ClientCoreCommand is the client for core command.
	ClientCoreCommand *resty.Client
)

// inits the clients.
func init() {
	viper.SetDefault("client.coredata.protocol", "http")
	viper.SetDefault("client.coredata.host", "localhost")
	viper.SetDefault("client.coredata.port", "48080")

	viper.SetDefault("client.coremetadata.protocol", "http")
	viper.SetDefault("client.coremetadata.host", "localhost")
	viper.SetDefault("client.coremetadata.port", "48081")

	viper.SetDefault("client.corecommand.protocol", "http")
	viper.SetDefault("client.corecommand.host", "localhost")
	viper.SetDefault("client.corecommand.port", "48082")

	ClientCoreData = resty.New()
	ClientCoreData.SetHostURL(fmt.Sprintf("%s://%s:%s",
		viper.GetString("client.coredata.protocol"),
		viper.GetString("client.coredata.host"),
		viper.GetString("client.coredata.port"),
	))

	ClientCoreMetaData = resty.New()
	ClientCoreMetaData.SetHostURL(fmt.Sprintf("%s://%s:%s",
		viper.GetString("client.coremetadata.protocol"),
		viper.GetString("client.coremetadata.host"),
		viper.GetString("client.coremetadata.port"),
	))

	ClientCoreCommand = resty.New()
	ClientCoreCommand.SetHostURL(fmt.Sprintf("%s://%s:%s",
		viper.GetString("client.corecommand.protocol"),
		viper.GetString("client.corecommand.host"),
		viper.GetString("client.corecommand.port"),
	))
}

// processRequestWithValidation executes the give request with the given path, checks if it was successful and returns the body.
func processRequestWithValidation(fun func(url string) (*resty.Response, error), path string) ([]byte, error) {
	resp, err := fun(path)
	if err != nil {
		return resp.Body(), err
	}

	if resp.IsError() {
		respErr := &ResponseError{Code: resp.StatusCode()}

		if err := json.Unmarshal(resp.Body(), &respErr.Message); err != nil {
			respErr.Message = string(resp.Body())
		}

		return resp.Body(), respErr
	}

	return resp.Body(), nil
}

// processGetRequestWithValidation shorthand for get.
func processGetRequestWithValidation(request *resty.Request, path string) ([]byte, error) {
	return processRequestWithValidation(request.Get, path)
}

// processPutRequestWithValidation shorthand for put.
func processPutRequestWithValidation(request *resty.Request, path string) ([]byte, error) {
	return processRequestWithValidation(request.Put, path)
}

// GetLastReadingForDeviceAndResource returns the last reading of the devices resource.
func GetLastReadingForDeviceAndResource(device, resource string) (*models.Reading, error) {
	var (
		readings []*models.Reading
		path     = "/api/v1/reading/name/{resource}/device/{device}/1"
	)

	_, err := processGetRequestWithValidation(
		ClientCoreData.
			R().
			SetPathParam("resource", resource).
			SetPathParam("device", device).
			SetHeader("Accept", "application/json").
			SetResult(&readings),
		path,
	)

	if err != nil {
		return nil, err
	}

	if len(readings) == 0 {
		return nil, &ResponseError{Code: 404, Message: "no reading found"}
	}

	return readings[0], nil
}

// GetValueDescriptorsForDevice returns a list of all valuedescriptors asociated with the device.
func GetValueDescriptorsForDevice(device string) ([]*models.ValueDescriptor, error) {
	var (
		descriptors []*models.ValueDescriptor
		path        = "/api/v1/valuedescriptor/devicename/{device}"
	)

	_, err := processGetRequestWithValidation(
		ClientCoreData.
			R().
			SetPathParam("device", device).
			SetHeader("Accept", "application/json").
			SetResult(&descriptors),
		path,
	)

	return descriptors, err
}

// GetDevices returns all devices.
func GetDevices() ([]*models.Device, error) {
	var (
		devices []*models.Device
		path    = "/api/v1/device"
	)

	_, err := processGetRequestWithValidation(
		ClientCoreMetaData.
			R().
			SetHeader("Accept", "application/json").
			SetResult(&devices),
		path,
	)

	return devices, err
}

// GetDevice returns the device for the given name.
func GetDevice(device string) (*models.Device, error) {
	var (
		d    models.Device
		path = fmt.Sprintf("/api/v1/device/name/%s", device)
	)

	_, err := processGetRequestWithValidation(
		ClientCoreMetaData.
			R().
			SetHeader("Accept", "application/json").
			SetResult(&d),
		path,
	)

	return &d, err
}

// GetDeviceProfileForDevice returns the deviceprofile for the given device.
func GetDeviceProfileForDevice(device string) (*models.DeviceProfile, error) {
	dev, err := GetDevice(device)
	if err != nil {
		return nil, err
	}

	return dev.Profile, nil
}

// processCommandResponse tries to create a map from the response. Returns a string if not possible.
func processCommandResponse(body []byte, err error) (interface{}, error) {
	if err != nil {
		return string(body), err
	}

	dec := make(map[string]interface{})

	err = json.Unmarshal(body, &dec)
	if err == nil {
		return dec, nil
	}

	return string(body), nil
}

// GetDeviceCommandForDeviceAndCommand executes a GET-command on the given device and returns the response.
func GetDeviceCommandForDeviceAndCommand(device, command string) (models.CommandResponse, error) {
	var path = "/api/v1/device/name/{device}/command/{command}"

	return processCommandResponse(
		processGetRequestWithValidation(
			ClientCoreCommand.
				R().
				SetPathParam("device", device).
				SetPathParam("command", command).
				SetHeader("Accept", "application/json"),
			path,
		),
	)
}

// PutDeviceCommandForDeviceAndCommand executes a PUT-command on the given device and returns it response.
func PutDeviceCommandForDeviceAndCommand(device, command string, payload models.CommandPayload) (models.CommandResponse, error) {
	var path = "/api/v1/device/name/{device}/command/{command}"

	return processCommandResponse(
		processPutRequestWithValidation(
			ClientCoreCommand.
				R().
				SetPathParam("device", device).
				SetPathParam("command", command).
				SetHeader("Content-Type", "application/json").
				SetHeader("Accept", "application/json").
				SetBody(payload),
			path,
		),
	)
}
