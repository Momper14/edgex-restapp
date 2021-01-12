package converter

import (
	db "github.com/momper14/edgex-restapp/db/models"
	"github.com/momper14/edgex-restapp/models"
)

// DbUserFromCreate creates a db/User from a UserCreate.
func DbUserFromCreate(from *models.UserCreate) *db.User {
	if from == nil {
		return nil
	}
	return &db.User{
		Name:     from.Name,
		Password: from.Password,
		Role:     from.Role,
	}
}

// DbUserFromUpdate creates a db/User from a UserUpdate.
func DbUserFromUpdate(from *models.UserUpdate) *db.User {
	if from == nil {
		return nil
	}
	return &db.User{
		Password: from.Password,
		Role:     from.Role,
	}
}

// DbUserFrom creates a db/User from a User.
func DbUserFrom(from *models.User) *db.User {
	if from == nil {
		return nil
	}
	return &db.User{
		Name: &from.Name,
		Role: &from.Role,
	}
}

// UserFrom creates a User from a db/User.
func UserFrom(from *db.User) *models.User {
	if from == nil {
		return nil
	}
	return &models.User{
		Name: *from.Name,
		Role: *from.Role,
	}
}

// UsersFrom creates a Role array from a db/Role array.
func UsersFrom(from []*db.User) []*models.User {
	if from == nil {
		return nil
	}

	var result = make([]*models.User, len(from))

	for i, u := range from {
		result[i] = UserFrom(u)
	}

	return result
}
