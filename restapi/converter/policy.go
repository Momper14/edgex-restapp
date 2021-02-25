package converter

import (
	db "github.com/momper14/edgex-restapp/db/models"
	"github.com/momper14/edgex-restapp/models"
)

// PolicyFrom creates a Policy from a db/Policy.
func PolicyFrom(from *db.Policy) *models.Policy {
	return &models.Policy{
		Role:     from.Role,
		Resource: from.Resource,
		Method:   from.Method,
	}
}

// DbPolicyFrom creates a db/Policy from a Policy.
func DbPolicyFrom(from *models.Policy) *db.Policy {
	return &db.Policy{
		Role:     from.Role,
		Resource: from.Resource,
		Method:   from.Method,
	}
}

// PolicysFrom creates a Policy array from a db/Policy array.
func PolicysFrom(from []*db.Policy) []*models.Policy {
	if from == nil {
		return nil
	}

	var result = make([]*models.Policy, len(from))

	for i, u := range from {
		result[i] = PolicyFrom(u)
	}

	return result
}
