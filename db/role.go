package db

import (
	"github.com/Momper14/skv"
	db "github.com/momper14/edgex-restapp/db/models"
	"github.com/momper14/edgex-restapp/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// RoleDB is the KV-Store for roles.
	RoleDB *skv.KVStore
)

// init inits the roleDB
// creates guest and admin role if database didn't exist
func init() {
	var (
		err  error
		init bool
	)

	viper.SetDefault("db.roledb", "persist/role.db")
	viper.SetDefault("role.default", "guest")
	viper.SetDefault("admin.role", "admin")

	file := viper.GetString("db.roledb")

	init = !util.FileExists(file)

	if init {
		util.CreateParents(file)
	}

	if RoleDB, err = skv.Open(file); err != nil {
		logrus.Fatalln(err)
	}

	if init {
		adminrole := viper.GetString("admin.role")
		defaultrole := viper.GetString("role.default")

		if err = RoleDB.Put(adminrole, db.Role(adminrole)); err != nil {
			logrus.Fatalln(err)
		}

		if err = RoleDB.Put(defaultrole, db.Role(defaultrole)); err != nil {
			logrus.Fatalln(err)
		}
	}

}

// GetRole returns the role with the name.
// possible errors: ErrNotFound.
func GetRole(name string) (db.Role, error) {
	var role db.Role

	err := RoleDB.Get(name, &role)

	if err == skv.ErrNotFound {
		return role, ErrNotFound
	}

	return role, err
}

// GetRoles returns alls roles.
func GetRoles() ([]db.Role, error) {
	names, err := RoleDB.GetKeys()
	if err != nil {
		return nil, err
	}

	roles := make([]db.Role, len(names))

	for i, name := range names {
		err := RoleDB.Get(name, &roles[i])
		if err != nil {
			return roles, err
		}
	}

	return []db.Role(roles), err
}

// RoleExists checks if the role exists.
func RoleExists(name string) (bool, error) {
	_, err := GetRole(name)

	if err == nil {
		return true, nil
	}

	if err == ErrNotFound {
		return false, nil
	}

	return false, err
}

// CreateRole creates the role.
// possible errors: ErrAlreadyExists.
func CreateRole(role db.Role) error {
	exists, err := RoleExists(string(role))
	if err != nil {
		return err
	}

	if exists {
		return ErrAlreadyExists
	}

	return RoleDB.Put(string(role), role)
}

// DeleteRole deletes the role.
// possible errors: ErrNotFound, ErrRoleHasUsers, ErrRoleHasPolicies.
func DeleteRole(name string) error {
	users, err := GetUsersByRole(name)
	if err != nil {
		if err == ErrRoleNotFound {
			return ErrNotFound
		}
	}

	if len(users) != 0 {
		return ErrRoleHasUsers
	}

	policies, err := GetPolicies()
	if err != nil {
		return err
	}

	for _, p := range policies {
		if *p.Role == name {
			return ErrRoleHasPolicies
		}
	}

	return RoleDB.Delete(name)
}
