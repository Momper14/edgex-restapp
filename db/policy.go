package db

import (
	"os"

	"github.com/casbin/casbin/v2"
	db "github.com/momper14/edgex-restapp/db/models"
	"github.com/momper14/edgex-restapp/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// model contains the content of model.conf for casbin
// used to create a new file if its missing
const model = `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch3(r.obj, p.obj) && regexMatch(r.act, p.act)
`

// Enforcer is the casbin enforcer used for authorization.
var Enforcer *casbin.Enforcer

// create a new enforcer and initial config if missing
// config is a policy for admin to access everything
func init() {
	var (
		err                   error
		initModel, initPolicy bool
	)

	viper.SetDefault("enforcer.model", "persist/model.conf")
	viper.SetDefault("enforcer.policy", "persist/policy.csv")
	viper.SetDefault("admin.role", "admin")

	fileModel := viper.GetString("enforcer.model")
	filePolicy := viper.GetString("enforcer.policy")

	initModel = !util.FileExists(fileModel)
	initPolicy = !util.FileExists(filePolicy)

	if initModel {
		util.CreateParents(fileModel)
		f, err := os.Create(fileModel)
		if err != nil {
			logrus.Fatalln(err)
		}
		//#nosec
		defer f.Close()

		if _, err = f.WriteString(model); err != nil {
			logrus.Fatalln(err)
		}

		//#nosec
		//nolint:errcheck
		f.Sync()
	}

	if initPolicy {
		util.CreateParents(filePolicy)

		f, err := os.Create(filePolicy)
		if err != nil {
			logrus.Fatalln(err)
		}

		//#nosec
		f.Close()
	}

	Enforcer, err = casbin.NewEnforcer(fileModel, filePolicy)

	if err != nil {
		logrus.Fatalln(err)
	}

	if initPolicy {
		if _, err := Enforcer.AddPolicy(viper.GetString("admin.role"), ".*", ".*"); err != nil {
			logrus.Fatalln(err)
		}
		if err := Enforcer.SavePolicy(); err != nil {
			logrus.Fatalln()
		}
	}
}

// AddPolicy adds the policy.
// possible errors: ErrRoleNotFound, ErrAlreadyExists.
func AddPolicy(p *db.Policy) (err error) {
	defer func() { recovery(&err) }()

	roleExists, err := RoleExists(*p.Role)
	if err != nil {
		return err
	}

	if !roleExists {
		return ErrRoleNotFound
	}

	added, err := Enforcer.AddPolicy(*p.Role, *p.Resource, *p.Method)
	if err != nil {
		return err
	}

	if !added {
		return ErrAlreadyExists
	}

	return Enforcer.SavePolicy()
}

// DeletePolicy deletes the policy.
// possible errors: ErrNotFound.
func DeletePolicy(p *db.Policy) (err error) {
	defer func() { recovery(&err) }()

	deleted, err := Enforcer.RemovePolicy(*p.Role, *p.Resource, *p.Method)
	if err != nil {
		return err
	}

	if !deleted {
		return ErrNotFound
	}

	return Enforcer.SavePolicy()

}

// GetPolicies returns all policies.
func GetPolicies() (result []*db.Policy, err error) {
	defer func() { recovery(&err) }()

	policies := Enforcer.GetPolicy()

	result = make([]*db.Policy, len(policies))

	for i, p := range policies {
		result[i] = &db.Policy{
			Role:     &p[0],
			Resource: &p[1],
			Method:   &p[2],
		}
	}

	return result, nil
}
