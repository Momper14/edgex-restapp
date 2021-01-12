package restapi

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/momper14/edgex-restapp/converter"
	"github.com/momper14/edgex-restapp/db"
	"github.com/momper14/edgex-restapp/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// authenticate authenticates the user with the password.
func authenticate(user string, pass string) (*models.User, error) {
	u, err := db.GetUser(user)
	apiUser := converter.UserFrom(u)
	if err != nil {
		if err == db.ErrNotFound {
			logrus.Infof("User '%s' tried to authenticate but user does not exist\n", user)
			return apiUser, errors.New(http.StatusUnauthorized, "Wrong username or password")
		}
		logrus.Errorln(err)
		return apiUser, errors.New(http.StatusInternalServerError, err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(pass)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			logrus.Infof("User '%s' tried to authenticate but password was wrong\n", user)
			return apiUser, errors.New(http.StatusUnauthorized, "Wrong username or password")
		}
		logrus.Errorln(err)
		return apiUser, errors.New(http.StatusInternalServerError, err.Error())
	}

	logrus.Debugf("successfull authenticated user %s\n", *u.Name)
	return apiUser, nil
}
