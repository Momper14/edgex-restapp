package db

import (
	"github.com/Momper14/skv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	db "github.com/momper14/edgex-restapp/db/models"
	"github.com/momper14/edgex-restapp/util"
)

var (
	// UserDB is the KV-Store for users.
	UserDB *skv.KVStore
)

// init inits the userDB
// creates admin user if database didn't exist
func init() {
	var (
		err  error
		init bool
	)

	viper.SetDefault("db.userdb", "persist/user.db")
	viper.SetDefault("admin.username", "admin")
	viper.SetDefault("admin.password", "password")
	viper.SetDefault("admin.role", "admin")

	file := viper.GetString("db.userdb")

	init = !util.FileExists(file)

	if init {
		util.CreateParents(file)
	}

	if UserDB, err = skv.Open(file); err != nil {
		logrus.Fatalln(err)
	}

	if init {
		name := viper.GetString("admin.username")
		password := viper.GetString("admin.password")
		role := viper.GetString("admin.role")

		hash, err := hashAndSalt(&password)
		if err != nil {
			logrus.Fatalln(err)
		}

		err = UserDB.Put(name, &db.User{
			Name:     &name,
			Password: hash,
			Role:     &role,
		})
		if err != nil {
			logrus.Fatalln(err)
		}
	}
}

// GetUser returns the associated user.
// possible errors: ErrNotFound.
func GetUser(name string) (*db.User, error) {
	var user *db.User

	err := UserDB.Get(name, &user)

	if err == skv.ErrNotFound {
		return user, ErrNotFound
	}

	return user, err
}

// GetUsers returns all users.
func GetUsers() ([]*db.User, error) {
	var users []*db.User

	names, err := UserDB.GetKeys()
	if err != nil {
		return nil, err
	}

	users = make([]*db.User, len(names))

	for i, name := range names {
		err := UserDB.Get(name, &users[i])
		if err != nil {
			return users, err
		}
	}

	return users, nil
}

// UserExists checks if the user exists.
func UserExists(name string) (bool, error) {
	_, err := GetUser(name)

	if err == nil {
		return true, nil
	}

	if err == ErrNotFound {
		return false, nil
	}

	return false, err
}

// validateRole validates if the role is valide.
// possible errors: ErrRoleNotFound.
func validateRole(role string) error {
	exists, err := RoleExists(role)
	if err != nil {
		return err
	}

	if !exists {
		return ErrRoleNotFound
	}

	return nil
}

// CreateUser creates the user and hash and salts the password.
// possible errors: ErrAlreadyExists, ErrRoleNotFound.
func CreateUser(user *db.User) error {
	exists, err := UserExists(*user.Name)
	if err != nil {
		return err
	}

	if exists {
		return ErrAlreadyExists
	}

	if err := validateRole(*user.Role); err != nil {
		return err
	}

	if user.Password, err = hashAndSalt(user.Password); err != nil {
		return err
	}

	return UserDB.Put(*user.Name, user)
}

// UpdateUser updates the user.
// hash and salts the password if updated.
// possible errors: ErrNotFound, ErrRoleNotFound.
func UpdateUser(name string, update *db.User) error {
	var user *db.User

	if err := UserDB.Get(name, &user); err != nil {
		if err == skv.ErrNotFound {
			return ErrNotFound
		}
		return err
	}

	if update.Role != nil {
		if err := validateRole(*update.Role); err != nil {
			return err
		}

		user.Role = update.Role
	}

	if update.Password != nil {
		var err error
		if user.Password, err = hashAndSalt(update.Password); err != nil {
			return err
		}
	}

	return UserDB.Put(name, user)
}

// DeleteUser deletes the user.
// possible errors: ErrNotFound.
func DeleteUser(name string) error {
	err := UserDB.Delete(name)

	if err == skv.ErrNotFound {
		return ErrNotFound
	}

	return err
}

// GetUsersByRole returns all users associated with the role
// possible errors: ErrRoleNotFound
func GetUsersByRole(role string) ([]*db.User, error) {

	if err := validateRole(role); err != nil {
		return nil, err
	}

	names, err := UserDB.GetKeys()
	if err != nil {
		return nil, err
	}

	var users []*db.User

	for _, name := range names {
		user := new(db.User)
		err := UserDB.Get(name, &user)
		if err != nil {
			return users, err
		}
		if *user.Role == role {
			users = append(users, user)
		}
	}

	return users, nil
}
