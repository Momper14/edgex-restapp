package restapi_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/momper14/edgex-restapp/client"
	clientmodels "github.com/momper14/edgex-restapp/client/models"
	"github.com/momper14/edgex-restapp/models"
	"github.com/momper14/edgex-restapp/restapi/converter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDevicesHandler200(t *testing.T) {
	var (
		mockResponse []*clientmodels.Device

		expect []*models.Device
		actual []*models.Device
	)

	httpmock.ActivateNonDefault(client.ClientCoreMetaData.GetClient())
	t.Cleanup(httpmock.DeactivateAndReset)

	require.NoError(t, readJSONFile("./json/devices.json", &mockResponse))

	expect = converter.DevicesFrom(mockResponse)

	httpmock.RegisterResponder(
		"GET",
		"/api/v1/device",
		httpmock.NewJsonResponderOrPanic(200, mockResponse),
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		SetResult(&actual).
		Get(pathDevices)
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.EqualValues(t, expect, actual)
}

func TestGetDevicesHandler401(t *testing.T) {
	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Accept", "application/json").
		Get(pathDevices)
	testUnauthorized(t, resp, err)
}

func TestGetDevicesHandler403(t *testing.T) {
	resp, err := rclient.R().
		SetHeader("Accept", "application/json").
		Get(pathDevices)
	testForbidden(t, resp, err)
}

func TestGetDevicesDeviceHandler200(t *testing.T) {
	var (
		mockResponse *clientmodels.Device

		expect *models.Device
		actual *models.Device

		deviceName = "Random-Integer-Generator01"
	)

	httpmock.ActivateNonDefault(client.ClientCoreMetaData.GetClient())
	t.Cleanup(httpmock.DeactivateAndReset)

	require.NoError(t, readJSONFile("./json/device.json", &mockResponse))

	expect = converter.DeviceFrom(mockResponse)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("/api/v1/device/name/%s", deviceName),
		httpmock.NewJsonResponderOrPanic(200, mockResponse),
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		SetResult(&actual).
		Get(fmt.Sprintf("%s/%s", pathDevices, deviceName))
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.EqualValues(t, expect, actual)
}

func TestGetDevicesDeviceHandler401(t *testing.T) {
	var (
		deviceName = "Random-Integer-Generator01"
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s", pathDevices, deviceName))
	testUnauthorized(t, resp, err)
}

func TestGetDevicesDeviceHandler403(t *testing.T) {
	var (
		deviceName = "Random-Integer-Generator01"
	)

	resp, err := rclient.R().
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s", pathDevices, deviceName))
	testForbidden(t, resp, err)
}

func TestGetDevicesDeviceHandler404(t *testing.T) {
	var (
		expect = "Item not found"
		actual string

		deviceName = "missing"
	)

	httpmock.ActivateNonDefault(client.ClientCoreMetaData.GetClient())
	t.Cleanup(httpmock.DeactivateAndReset)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("/api/v1/device/name/%s", deviceName),
		httpmock.NewJsonResponderOrPanic(404, expect),
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s", pathDevices, deviceName))
	require.NoError(t, err)

	assert.Equal(t, 404, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expect, actual)
}

func TestGetDevicesDeviceProfileHandler200(t *testing.T) {
	var (
		mockResponse *clientmodels.Device

		expect *models.DeviceProfile
		actual *models.DeviceProfile

		deviceName = "Random-Integer-Generator01"
	)

	httpmock.ActivateNonDefault(client.ClientCoreMetaData.GetClient())
	t.Cleanup(httpmock.DeactivateAndReset)

	require.NoError(t, readJSONFile("./json/device.json", &mockResponse))

	expect = converter.DeviceProfileFrom(mockResponse.Profile)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("/api/v1/device/name/%s", deviceName),
		httpmock.NewJsonResponderOrPanic(200, mockResponse),
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		SetResult(&actual).
		Get(fmt.Sprintf("%s/%s/profile", pathDevices, deviceName))
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.EqualValues(t, expect, actual)
}

func TestGetDevicesDeviceProfileHandler401(t *testing.T) {
	var (
		deviceName = "Random-Integer-Generator01"
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s/profile", pathDevices, deviceName))
	testUnauthorized(t, resp, err)
}

func TestGetDevicesDeviceProfileHandler403(t *testing.T) {
	var (
		deviceName = "Random-Integer-Generator01"
	)

	resp, err := rclient.R().
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s/profile", pathDevices, deviceName))
	testForbidden(t, resp, err)
}

func TestGetDevicesDeviceProfileHandler404(t *testing.T) {
	var (
		deviceName = "missing"

		expect = "Item not found"
		actual string
	)

	httpmock.ActivateNonDefault(client.ClientCoreMetaData.GetClient())
	t.Cleanup(httpmock.DeactivateAndReset)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("/api/v1/device/name/%s", deviceName),
		httpmock.NewJsonResponderOrPanic(404, expect),
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s/profile", pathDevices, deviceName))
	require.NoError(t, err)

	assert.Equal(t, 404, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expect, actual)
}

func TestGetDevicesDeviceResourcesResourceHandler200(t *testing.T) {
	var (
		mockResponse []*clientmodels.Reading

		expect *models.Reading
		actual *models.Reading

		deviceName     = "Random-Integer-Generator01"
		deviceResource = "RandomValue_Int8"
	)

	httpmock.ActivateNonDefault(client.ClientCoreData.GetClient())
	t.Cleanup(httpmock.DeactivateAndReset)

	require.NoError(t, readJSONFile("./json/reading.json", &mockResponse))

	expect = converter.ReadingFrom(mockResponse[0])

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("/api/v1/reading/name/%s/device/%s/1", deviceResource, deviceName),
		httpmock.NewJsonResponderOrPanic(200, mockResponse),
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		SetResult(&actual).
		Get(fmt.Sprintf("%s/%s/resources/%s", pathDevices, deviceName, deviceResource))
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.EqualValues(t, expect, actual)
}

func TestGetDevicesDeviceResourcesResourceHandler401(t *testing.T) {
	var (
		deviceName     = "Random-Integer-Generator01"
		deviceResource = "RandomValue_Int8"
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s/resources/%s", pathDevices, deviceName, deviceResource))
	testUnauthorized(t, resp, err)
}

func TestGetDevicesDeviceResourcesResourceHandler403(t *testing.T) {
	var (
		deviceName     = "Random-Integer-Generator01"
		deviceResource = "RandomValue_Int8"
	)

	resp, err := rclient.R().
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s/resources/%s", pathDevices, deviceName, deviceResource))
	testForbidden(t, resp, err)
}

func TestGetDevicesDeviceResourcesResourceHandler404(t *testing.T) {
	var (
		deviceName     = "Random-Integer-Generator01"
		deviceResource = "missing"

		expect = "Item not found"
		actual string
	)

	httpmock.ActivateNonDefault(client.ClientCoreData.GetClient())
	t.Cleanup(httpmock.DeactivateAndReset)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("/api/v1/reading/name/%s/device/%s/1", deviceResource, deviceName),
		httpmock.NewJsonResponderOrPanic(404, expect),
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s/resources/%s", pathDevices, deviceName, deviceResource))
	require.NoError(t, err)

	assert.Equal(t, 404, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expect, actual)
}

func TestGetDevicesDeviceValuedescriptorsHandler200(t *testing.T) {
	var (
		mockResponse []*clientmodels.ValueDescriptor

		expect []*models.ValueDescriptor
		actual []*models.ValueDescriptor

		deviceName = "Random-Integer-Generator01"
	)

	httpmock.ActivateNonDefault(client.ClientCoreData.GetClient())
	t.Cleanup(httpmock.DeactivateAndReset)

	require.NoError(t, readJSONFile("./json/valuedescriptors.json", &mockResponse))

	expect = converter.ValueDescriptorsFrom(mockResponse)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("/api/v1/valuedescriptor/devicename/%s", deviceName),
		httpmock.NewJsonResponderOrPanic(200, mockResponse),
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		SetResult(&actual).
		Get(fmt.Sprintf("%s/%s/valuedescriptors", pathDevices, deviceName))
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.EqualValues(t, expect, actual)
}

func TestGetDevicesDeviceValuedescriptorsHandler401(t *testing.T) {
	var (
		deviceName = "Random-Integer-Generator01"
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s/valuedescriptors", pathDevices, deviceName))
	testUnauthorized(t, resp, err)
}

func TestGetDevicesDeviceValuedescriptorsHandler403(t *testing.T) {
	var (
		deviceName = "Random-Integer-Generator01"
	)

	resp, err := rclient.R().
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s/valuedescriptors", pathDevices, deviceName))
	testForbidden(t, resp, err)
}

func TestGetDevicesDeviceValuedescriptorsHandler404(t *testing.T) {
	var (
		deviceName = "missing"

		expect = "Item not found"
		actual string
	)

	httpmock.ActivateNonDefault(client.ClientCoreData.GetClient())
	t.Cleanup(httpmock.DeactivateAndReset)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("/api/v1/valuedescriptor/devicename/%s", deviceName),
		httpmock.NewJsonResponderOrPanic(404, expect),
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s/valuedescriptors", pathDevices, deviceName))
	require.NoError(t, err)

	assert.Equal(t, 404, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expect, actual)
}

func TestGetDevicesDeviceCommandsCommandHandler200(t *testing.T) {
	var (
		mockResponse clientmodels.CommandResponse = make(map[string]interface{})

		expect models.CommandResponse
		actual models.CommandResponse

		deviceName    = "Random-Integer-Generator01"
		deviceCommand = "GenerateRandomValue_Int8"
	)

	httpmock.ActivateNonDefault(client.ClientCoreCommand.GetClient())
	t.Cleanup(httpmock.DeactivateAndReset)

	require.NoError(t, readJSONFile("./json/get-command-response.json", &mockResponse))

	expect = converter.CommandResponseFrom(mockResponse)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("/api/v1/device/name/%s/command/%s", deviceName, deviceCommand),
		httpmock.NewJsonResponderOrPanic(200, mockResponse),
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		SetResult(&actual).
		Get(fmt.Sprintf("%s/%s/commands/%s", pathDevices, deviceName, deviceCommand))
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.EqualValues(t, expect, actual)
}

func TestGetDevicesDeviceCommandsCommandHandler401(t *testing.T) {
	var (
		deviceName    = "Random-Integer-Generator01"
		deviceCommand = "GenerateRandomValue_Int8"
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s/commands/%s", pathDevices, deviceName, deviceCommand))
	testUnauthorized(t, resp, err)
}

func TestGetDevicesDeviceCommandsCommandHandler403(t *testing.T) {
	var (
		deviceName    = "Random-Integer-Generator01"
		deviceCommand = "GenerateRandomValue_Int8"
	)

	resp, err := rclient.R().
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s/commands/%s", pathDevices, deviceName, deviceCommand))
	testForbidden(t, resp, err)
}

func TestGetDevicesDeviceCommandsCommandHandler404(t *testing.T) {
	var (
		deviceName    = "Random-Integer-Generator01"
		deviceCommand = "missing"

		expect = "command not found"
		actual string
	)

	httpmock.ActivateNonDefault(client.ClientCoreCommand.GetClient())
	t.Cleanup(httpmock.DeactivateAndReset)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("/api/v1/device/name/%s/command/%s", deviceName, deviceCommand),
		httpmock.NewJsonResponderOrPanic(404, expect),
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s/commands/%s", pathDevices, deviceName, deviceCommand))
	require.NoError(t, err)

	assert.Equal(t, 404, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expect, actual)
}

func TestPutDevicesDeviceCommandsCommandHandler200(t *testing.T) {
	var (
		mockResponse clientmodels.CommandResponse

		payload models.CommandPayload = map[string]interface{}{
			"Min_Int8": "5",
			"Max_Int8": "10",
		}

		expect models.CommandResponse
		actual models.CommandResponse

		deviceName    = "Random-Integer-Generator01"
		deviceCommand = "GenerateRandomValue_Int8"
	)

	httpmock.ActivateNonDefault(client.ClientCoreCommand.GetClient())
	t.Cleanup(httpmock.DeactivateAndReset)

	expect = converter.CommandResponseFrom(mockResponse)

	httpmock.RegisterResponder(
		"PUT",
		fmt.Sprintf("/api/v1/device/name/%s/command/%s", deviceName, deviceCommand),
		func(req *http.Request) (*http.Response, error) {
			var (
				content []byte
				actual  map[string]interface{}
			)

			content, err := ioutil.ReadAll(req.Body)
			require.NoError(t, err)

			require.NoError(t, json.Unmarshal(content, &actual))

			assert.EqualValues(t, payload, actual)

			return httpmock.NewJsonResponse(200, mockResponse)
		},
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		SetBody(payload).
		SetResult(&actual).
		Put(fmt.Sprintf("%s/%s/commands/%s", pathDevices, deviceName, deviceCommand))
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.EqualValues(t, expect, actual)
}

func TestPutDevicesDeviceCommandsCommandHandler401(t *testing.T) {
	var (
		deviceName    = "Random-Integer-Generator01"
		deviceCommand = "GenerateRandomValue_Int8"
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Accept", "application/json").
		Put(fmt.Sprintf("%s/%s/commands/%s", pathDevices, deviceName, deviceCommand))
	testUnauthorized(t, resp, err)
}

func TestPutDevicesDeviceCommandsCommandHandler403(t *testing.T) {
	var (
		deviceName    = "Random-Integer-Generator01"
		deviceCommand = "GenerateRandomValue_Int8"
	)

	resp, err := rclient.R().
		SetHeader("Accept", "application/json").
		Put(fmt.Sprintf("%s/%s/commands/%s", pathDevices, deviceName, deviceCommand))
	testForbidden(t, resp, err)
}

func TestPutDevicesDeviceCommandsCommandHandler404(t *testing.T) {
	var (
		deviceName    = "Random-Integer-Generator01"
		deviceCommand = "missing"

		expect = "command not found"
		actual string
	)

	httpmock.ActivateNonDefault(client.ClientCoreCommand.GetClient())
	t.Cleanup(httpmock.DeactivateAndReset)

	httpmock.RegisterResponder(
		"PUT",
		fmt.Sprintf("/api/v1/device/name/%s/command/%s", deviceName, deviceCommand),
		httpmock.NewJsonResponderOrPanic(404, expect),
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		Put(fmt.Sprintf("%s/%s/commands/%s", pathDevices, deviceName, deviceCommand))
	require.NoError(t, err)

	assert.Equal(t, 404, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expect, actual)
}
