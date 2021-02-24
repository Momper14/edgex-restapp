package restapi

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/momper14/edgex-restapp/db"
	"github.com/momper14/edgex-restapp/models"
	"github.com/momper14/edgex-restapp/restapi/converter"
	"github.com/momper14/edgex-restapp/restapi/operations"
	"github.com/sirupsen/logrus"
)

func configureUsers(api *operations.EdgexRestappAPI) {
	api.DeleteV1AdminUsersUserHandler = operations.DeleteV1AdminUsersUserHandlerFunc(
		func(params operations.DeleteV1AdminUsersUserParams, principal *models.User) middleware.Responder {
			err := db.DeleteUser(params.User)
			if err != nil {
				if err == db.ErrNotFound {
					logrus.Infof("User '%s' does not exist\n", params.User)
					return operations.NewDeleteV1AdminUsersUserNotFound().WithPayload(fmt.Sprintf("User '%s' does not exist", params.User))
				}
				logrus.Errorln(err)
				return operations.NewDeleteV1AdminUsersUserInternalServerError().WithPayload(err.Error())
			}

			return operations.NewDeleteV1AdminUsersUserOK().WithPayload("Deleted")
		},
	)

	api.GetV1AdminUsersHandler = operations.GetV1AdminUsersHandlerFunc(
		func(params operations.GetV1AdminUsersParams, principal *models.User) middleware.Responder {
			users, err := db.GetUsers()
			if err != nil {
				logrus.Errorln(err)
				return operations.NewGetV1AdminUsersInternalServerError().WithPayload(err.Error())
			}

			offset := int(*params.Offset)
			limit := int(*params.Limit)

			if offset >= len(users) {
				users = users[0:0]
			} else if offset+limit >= len(users) {
				users = users[offset:]
			} else {
				users = users[offset:limit]
			}

			return operations.NewGetV1AdminUsersOK().WithPayload(converter.UsersFrom(users))
		},
	)

	api.GetV1AdminUsersUserHandler = operations.GetV1AdminUsersUserHandlerFunc(
		func(params operations.GetV1AdminUsersUserParams, principal *models.User) middleware.Responder {
			user, err := db.GetUser(params.User)
			if err != nil {
				if err == db.ErrNotFound {
					logrus.Infof("User '%s' does not exist\n", params.User)
					return operations.NewGetV1AdminUsersUserNotFound().WithPayload(fmt.Sprintf("User '%s' does not exist", params.User))
				}
				logrus.Errorln(err)
				return operations.NewGetV1AdminUsersUserInternalServerError().WithPayload(err.Error())
			}

			return operations.NewGetV1AdminUsersUserOK().WithPayload(converter.UserFrom(user))
		},
	)

	api.PostV1AdminUsersHandler = operations.PostV1AdminUsersHandlerFunc(
		func(params operations.PostV1AdminUsersParams, principal *models.User) middleware.Responder {
			err := db.CreateUser(converter.DbUserFromCreate(params.Body))
			if err != nil {
				switch err {
				case db.ErrAlreadyExists:
					logrus.Infof("User '%s' already exists\n", *params.Body.Name)
					return operations.NewPostV1AdminUsersConflict().WithPayload(fmt.Sprintf("User '%s' already exists", *params.Body.Name))
				case db.ErrRoleNotFound:
					logrus.Infof("Role '%s' does not exist\n", *params.Body.Role)
					return operations.NewPostV1AdminUsersConflict().WithPayload(fmt.Sprintf("Role '%s' does not exist", *params.Body.Role))
				default:
					logrus.Errorln(err)
					return operations.NewPostV1AdminUsersInternalServerError().WithPayload(err.Error())
				}
			}

			user, err := db.GetUser(*params.Body.Name)
			if err != nil {
				logrus.Errorln(err)
				return operations.NewPostV1AdminUsersInternalServerError().WithPayload(err.Error())
			}

			return operations.NewPostV1AdminUsersCreated().WithPayload(converter.UserFrom(user))
		},
	)

	api.PatchV1AdminUsersUserHandler = operations.PatchV1AdminUsersUserHandlerFunc(
		func(params operations.PatchV1AdminUsersUserParams, principal *models.User) middleware.Responder {
			if params.Body.Password == nil && params.Body.Role == nil {
				return operations.NewPatchV1AdminUsersUserBadRequest().WithPayload("Nothing to update")
			}

			err := db.UpdateUser(params.User, converter.DbUserFromUpdate(params.Body))
			if err != nil {
				switch err {
				case db.ErrNotFound:
					logrus.Infof("User '%s' does not exist\n", params.User)
					return operations.NewGetV1AdminUsersUserNotFound().WithPayload(fmt.Sprintf("User '%s' does not exist", params.User))
				case db.ErrRoleNotFound:
					logrus.Infof("Role '%s' not found\n", *params.Body.Role)
					return operations.NewPostV1AdminUsersConflict().WithPayload(fmt.Sprintf("Role '%s' not found", *params.Body.Role))
				default:
					logrus.Errorln(err)
					return operations.NewPatchV1AdminUsersUserInternalServerError().WithPayload(err.Error())
				}
			}

			return operations.NewPatchV1AdminUsersUserOK().WithPayload("Updated")
		},
	)
}
