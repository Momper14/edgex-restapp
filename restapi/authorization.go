package restapi

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/momper14/edgex-restapp/db"
	"github.com/momper14/edgex-restapp/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// authorizerFunc returns an AuthorizerFunc for go-swagger
func authorizerFunc() runtime.AuthorizerFunc {
	return runtime.AuthorizerFunc(
		func(r *http.Request, input interface{}) error {
			var principal *models.User

			if input == nil {
				logrus.Debugln("No user given. Assuming default role")
				principal = &models.User{Role: viper.GetString("role.default")}
			} else {
				principal = input.(*models.User)
			}

			auth, err := db.Enforcer.Enforce(principal.Role, r.URL.Path, r.Method)
			if err != nil {
				logrus.Errorln(err)
				return errors.New(http.StatusInternalServerError, err.Error())
			}

			if !auth {
				logrus.Infof("User '%s' of group '%s' tried to access '%s %s' but is unauthorized\n", principal.Name, principal.Role, r.Method, r.URL.Path)
				return errors.New(http.StatusForbidden, "Forbidden")
			}

			logrus.Debugf("User '%s' of group '%s' authorized for '%s %s'\n", principal.Name, principal.Role, r.Method, r.URL.Path)
			return nil
		},
	)
}
