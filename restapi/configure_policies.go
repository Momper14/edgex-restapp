package restapi

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/momper14/edgex-restapp/db"
	"github.com/momper14/edgex-restapp/models"
	"github.com/momper14/edgex-restapp/restapi/converter"
	"github.com/momper14/edgex-restapp/restapi/operations"
	"github.com/sirupsen/logrus"
)

func configurePolicies(api *operations.EdgexRestappAPI) {

	api.DeleteV1AdminPoliciesHandler = operations.DeleteV1AdminPoliciesHandlerFunc(
		func(params operations.DeleteV1AdminPoliciesParams, principal *models.User) middleware.Responder {
			err := db.DeletePolicy(converter.DbPolicyFrom(params.Body))
			if err != nil {
				if err == db.ErrNotFound {
					encoded, _ := json.Marshal(params.Body)
					logrus.Infof("Policy '%s' does not exist\n", encoded)
					return operations.NewDeleteV1AdminPoliciesNotFound().WithPayload(fmt.Sprintf("Policy '%s' does not exist", encoded))
				}
				logrus.Errorln(err)
				return operations.NewDeleteV1AdminPoliciesInternalServerError().WithPayload(err.Error())
			}

			return operations.NewDeleteV1AdminPoliciesOK().WithPayload("Deleted")
		},
	)

	api.GetV1AdminPoliciesHandler = operations.GetV1AdminPoliciesHandlerFunc(
		func(params operations.GetV1AdminPoliciesParams, principal *models.User) middleware.Responder {
			policies, err := db.GetPolicies()
			if err != nil {
				logrus.Errorln(err)
				return operations.NewGetV1AdminPoliciesInternalServerError().WithPayload(err.Error())
			}

			offset := int(*params.Offset)
			limit := int(*params.Limit)

			if offset >= len(policies) {
				policies = policies[0:0]
			} else if offset+limit >= len(policies) {
				policies = policies[offset:]
			} else {
				policies = policies[offset:limit]
			}

			return operations.NewGetV1AdminPoliciesOK().WithPayload(converter.PolicysFrom(policies))
		},
	)

	api.PostV1AdminPoliciesHandler = operations.PostV1AdminPoliciesHandlerFunc(
		func(params operations.PostV1AdminPoliciesParams, principal *models.User) middleware.Responder {
			err := db.AddPolicy(converter.DbPolicyFrom(params.Body))
			if err != nil {
				switch err {
				case db.ErrRoleNotFound:
					logrus.Infof("Role '%s' does not exist\n", *params.Body.Role)
					return operations.NewPostV1AdminPoliciesConflict().WithPayload(fmt.Sprintf("Role '%s' does not exist", *params.Body.Role))
				case db.ErrAlreadyExists:
					encoded, _ := json.Marshal(params.Body)
					logrus.Infof("Policy '%s' already exists\n", encoded)
					return operations.NewPostV1AdminPoliciesConflict().WithPayload(fmt.Sprintf("Policy '%s' already exists", encoded))
				default:
					logrus.Errorln(err)
					return operations.NewPostV1AdminPoliciesInternalServerError().WithPayload(err.Error())
				}
			}
			return operations.NewPostV1AdminPoliciesCreated().WithPayload(params.Body)
		},
	)
}
