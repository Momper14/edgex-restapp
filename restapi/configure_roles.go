package restapi

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/momper14/edgex-restapp/converter"
	"github.com/momper14/edgex-restapp/db"
	"github.com/momper14/edgex-restapp/models"
	"github.com/momper14/edgex-restapp/restapi/operations"
	"github.com/sirupsen/logrus"
)

func configureRoles(api *operations.EdgexRestappAPI) {

	api.DeleteV1AdminRolesRoleHandler = operations.DeleteV1AdminRolesRoleHandlerFunc(
		func(params operations.DeleteV1AdminRolesRoleParams, principal *models.User) middleware.Responder {
			err := db.DeleteRole(params.Role)
			if err != nil {
				switch err {
				case db.ErrNotFound:
					logrus.Infof("Role '%s' does not exist\n", params.Role)
					return operations.NewDeleteV1AdminRolesRoleNotFound().WithPayload(fmt.Sprintf("Role '%s' does not exist", params.Role))
				case db.ErrRoleHasUsers:
					logrus.Infof("Role '%s' has associated users\n", params.Role)
					return operations.NewDeleteV1AdminRolesRoleConflict().WithPayload(fmt.Sprintf("Role '%s' has associated users", params.Role))
				case db.ErrRoleHasPolicies:
					logrus.Infof("Role '%s' has associated policies\n", params.Role)
					return operations.NewDeleteV1AdminRolesRoleConflict().WithPayload(fmt.Sprintf("Role '%s' has associated policies", params.Role))
				default:
					logrus.Errorln(err)
					return operations.NewDeleteV1AdminRolesRoleInternalServerError().WithPayload(err.Error())
				}
			}

			return operations.NewDeleteV1AdminRolesRoleOK().WithPayload("Deleted")

		},
	)

	api.GetV1AdminRolesHandler = operations.GetV1AdminRolesHandlerFunc(
		func(params operations.GetV1AdminRolesParams, principal *models.User) middleware.Responder {
			roles, err := db.GetRoles()
			if err != nil {
				logrus.Errorln(err)
				return operations.NewGetV1AdminRolesInternalServerError().WithPayload(err.Error())
			}

			offset := int(*params.Offset)
			limit := int(*params.Limit)

			if offset >= len(roles) {
				roles = roles[0:0]
			} else if offset+limit >= len(roles) {
				roles = roles[offset:]
			} else {
				roles = roles[offset:limit]
			}

			return operations.NewGetV1AdminRolesOK().WithPayload(converter.RolesFrom(roles))

		},
	)

	api.GetV1AdminRolesRoleUsersHandler = operations.GetV1AdminRolesRoleUsersHandlerFunc(
		func(params operations.GetV1AdminRolesRoleUsersParams, principal *models.User) middleware.Responder {
			users, err := db.GetUsersByRole(params.Role)
			if err != nil {
				if err == db.ErrRoleNotFound {
					logrus.Infof("Role '%s' does not exist\n", params.Role)
					return operations.NewGetV1AdminRolesRoleUsersNotFound().WithPayload(fmt.Sprintf("Role '%s' does not exist", params.Role))
				}
				logrus.Errorln(err)
				return operations.NewGetV1AdminRolesRoleUsersInternalServerError().WithPayload(err.Error())
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

			return operations.NewGetV1AdminRolesRoleUsersOK().WithPayload(converter.UsersFrom(users))
		},
	)

	api.PostV1AdminRolesHandler = operations.PostV1AdminRolesHandlerFunc(
		func(params operations.PostV1AdminRolesParams, principal *models.User) middleware.Responder {
			err := db.CreateRole(converter.DbRoleFrom(params.Body))
			if err != nil {
				if err == db.ErrAlreadyExists {
					logrus.Infof("Role '%s' already exists\n", params.Body)
					return operations.NewPostV1AdminRolesConflict().WithPayload(fmt.Sprintf("Role '%s' already exists", params.Body))
				}
				logrus.Errorln(err)
				return operations.NewPostV1AdminRolesInternalServerError().WithPayload(err.Error())
			}

			role, err := db.GetRole(string(params.Body))
			if err != nil {
				logrus.Errorln(err)
				return operations.NewPostV1AdminRolesInternalServerError().WithPayload(err.Error())
			}

			return operations.NewPostV1AdminRolesCreated().WithPayload(converter.RoleFrom(role))
		},
	)
}
