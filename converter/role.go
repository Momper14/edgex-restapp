package converter

import (
	db "github.com/momper14/edgex-restapp/db/models"
	"github.com/momper14/edgex-restapp/models"
)

// RoleFrom creates a Role from a db/Role.
func RoleFrom(from db.Role) models.Role {
	return models.Role(from)
}

// DbRoleFrom creates a db/Role from a Role.
func DbRoleFrom(from models.Role) db.Role {
	return db.Role(from)
}

// RolesFrom creates a Role array from a db/Role array.
func RolesFrom(from []db.Role) []models.Role {
	if from == nil {
		return nil
	}

	var result = make([]models.Role, len(from))

	for i, u := range from {
		result[i] = RoleFrom(u)
	}

	return result
}
