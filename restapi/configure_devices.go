package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/momper14/edgex-restapp/client"
	"github.com/momper14/edgex-restapp/converter"
	"github.com/momper14/edgex-restapp/models"
	"github.com/momper14/edgex-restapp/restapi/operations"
	"github.com/sirupsen/logrus"
)

func configureDevices(api *operations.EdgexRestappAPI) {

	api.GetV1DevicesHandler = operations.GetV1DevicesHandlerFunc(
		func(params operations.GetV1DevicesParams, principal *models.User) middleware.Responder {
			devices, err := client.GetDevices()
			if err != nil {
				resperr, ok := err.(*client.ResponseError)
				if ok {
					switch resperr.Code {
					case 500:
						logrus.Warnln(resperr.Message)
						return operations.NewGetV1DevicesInternalServerError().WithPayload(resperr.Message)
					}
				}
				logrus.Errorln(err)
				return operations.NewGetV1DevicesInternalServerError().WithPayload(err.Error())
			}

			offset := int(*params.Offset)
			limit := int(*params.Limit)

			if offset >= len(devices) {
				devices = devices[0:0]
			} else if offset+limit >= len(devices) {
				devices = devices[offset:]
			} else {
				devices = devices[offset:limit]
			}

			return operations.NewGetV1DevicesOK().WithPayload(converter.DevicesFrom(devices))
		},
	)

	api.GetV1DevicesDeviceHandler = operations.GetV1DevicesDeviceHandlerFunc(
		func(params operations.GetV1DevicesDeviceParams, principal *models.User) middleware.Responder {
			device, err := client.GetDevice(params.Device)
			if err != nil {
				resperr, ok := err.(*client.ResponseError)
				if ok {
					switch resperr.Code {
					case 404:
						logrus.Infoln(resperr.Message)
						return operations.NewGetV1DevicesDeviceNotFound().WithPayload(resperr.Message)
					case 500:
						logrus.Warnln(resperr.Message)
						return operations.NewGetV1DevicesDeviceInternalServerError().WithPayload(resperr.Message)
					}
				}
				logrus.Errorln(err)
				return operations.NewGetV1DevicesDeviceInternalServerError().WithPayload(err.Error())
			}
			return operations.NewGetV1DevicesDeviceOK().WithPayload(converter.DeviceFrom(device))
		},
	)

	api.GetV1DevicesDeviceCommandsCommandHandler = operations.GetV1DevicesDeviceCommandsCommandHandlerFunc(
		func(params operations.GetV1DevicesDeviceCommandsCommandParams, principal *models.User) middleware.Responder {
			resp, err := client.GetDeviceCommandForDeviceAndCommand(params.Device, params.Command)
			if err != nil {
				resperr, ok := err.(*client.ResponseError)
				if ok {
					switch resperr.Code {
					case 404:
						logrus.Infoln(resperr.Message)
						return operations.NewGetV1DevicesDeviceCommandsCommandNotFound().WithPayload(resperr.Message)
					case 500:
						logrus.Warnln(resperr.Message)
						return operations.NewGetV1DevicesDeviceCommandsCommandInternalServerError().WithPayload(resperr.Message)
					}
				}
				logrus.Errorln(err)
				return operations.NewGetV1DevicesDeviceCommandsCommandInternalServerError().WithPayload(err.Error())

			}
			return operations.NewGetV1DevicesDeviceCommandsCommandOK().WithPayload(converter.CommandResponseFrom(resp))
		},
	)

	api.GetV1DevicesDeviceProfileHandler = operations.GetV1DevicesDeviceProfileHandlerFunc(
		func(params operations.GetV1DevicesDeviceProfileParams, principal *models.User) middleware.Responder {
			profile, err := client.GetDeviceProfileForDevice(params.Device)
			if err != nil {
				resperr, ok := err.(*client.ResponseError)
				if ok {
					switch resperr.Code {
					case 404:
						logrus.Infoln(resperr.Message)
						return operations.NewGetV1DevicesDeviceProfileNotFound().WithPayload(resperr.Message)
					case 500:
						logrus.Warnln(resperr.Message)
						return operations.NewGetV1DevicesDeviceProfileInternalServerError().WithPayload(resperr.Message)
					}
				}
				logrus.Errorln(err)
				return operations.NewGetV1DevicesDeviceProfileInternalServerError().WithPayload(err.Error())
			}
			return operations.NewGetV1DevicesDeviceProfileOK().WithPayload(converter.DeviceProfileFrom(profile))
		},
	)

	api.GetV1DevicesDeviceResourcesResourceHandler = operations.GetV1DevicesDeviceResourcesResourceHandlerFunc(
		func(params operations.GetV1DevicesDeviceResourcesResourceParams, principal *models.User) middleware.Responder {
			reading, err := client.GetLastReadingForDeviceAndResource(params.Device, params.Resource)
			if err != nil {
				resperr, ok := err.(*client.ResponseError)
				if ok {
					switch resperr.Code {
					case 404:
						logrus.Infoln(resperr.Message)
						return operations.NewGetV1DevicesDeviceResourcesResourceNotFound().WithPayload(resperr.Message)
					case 500:
						logrus.Warnln(resperr.Message)
						return operations.NewGetV1DevicesDeviceResourcesResourceInternalServerError().WithPayload(resperr.Message)
					}
				}
				logrus.Errorln(err)
				return operations.NewGetV1DevicesDeviceResourcesResourceInternalServerError().WithPayload(err.Error())
			}
			return operations.NewGetV1DevicesDeviceResourcesResourceOK().WithPayload(converter.ReadingFrom(reading))
		},
	)

	api.GetV1DevicesDeviceValuedescriptorsHandler = operations.GetV1DevicesDeviceValuedescriptorsHandlerFunc(
		func(params operations.GetV1DevicesDeviceValuedescriptorsParams, principal *models.User) middleware.Responder {
			descriptors, err := client.GetValueDescriptorsForDevice(params.Device)
			if err != nil {
				resperr, ok := err.(*client.ResponseError)
				if ok {
					switch resperr.Code {
					case 404:
						logrus.Infoln(resperr.Message)
						return operations.NewGetV1DevicesDeviceValuedescriptorsNotFound().WithPayload(resperr.Message)
					case 500:
						logrus.Warnln(resperr.Message)
						return operations.NewGetV1DevicesDeviceValuedescriptorsInternalServerError().WithPayload(resperr.Message)
					}
				}
				logrus.Errorln(err)
				return operations.NewGetV1DevicesDeviceValuedescriptorsInternalServerError().WithPayload(err.Error())
			}
			return operations.NewGetV1DevicesDeviceValuedescriptorsOK().WithPayload(converter.ValueDescriptorsFrom(descriptors))
		},
	)

	api.PutV1DevicesDeviceCommandsCommandHandler = operations.PutV1DevicesDeviceCommandsCommandHandlerFunc(
		func(params operations.PutV1DevicesDeviceCommandsCommandParams, principal *models.User) middleware.Responder {
			resp, err := client.PutDeviceCommandForDeviceAndCommand(params.Device, params.Command, converter.CommandPayloadTo(params.Body))
			if err != nil {
				resperr, ok := err.(*client.ResponseError)
				if ok {
					switch resperr.Code {
					case 404:
						logrus.Infoln(resperr)
						return operations.NewPutV1DevicesDeviceCommandsCommandNotFound().WithPayload(resperr.Message)
					case 500:
						logrus.Warnln(err)
						return operations.NewPutV1DevicesDeviceCommandsCommandInternalServerError().WithPayload(resperr.Message)
					}
				}
				logrus.Errorln(err)
				return operations.NewPutV1DevicesDeviceCommandsCommandInternalServerError().WithPayload(err.Error())
			}
			return operations.NewPutV1DevicesDeviceCommandsCommandOK().WithPayload(converter.CommandResponseFrom(resp))
		},
	)
}
