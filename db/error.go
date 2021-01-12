package db

import (
	"errors"
	"fmt"

	"github.com/Momper14/skv"
)

//  ErrNotFound means that the requested entity was not found.
var ErrNotFound = skv.ErrNotFound

// ErrAlreadyExists means that the new entity already existed.
var ErrAlreadyExists = errors.New("new entity already existed")

// ErrRoleNotFound means that the associated role was not found.
var ErrRoleNotFound = errors.New("associated role was not found")

// ErrRoleHasUsers means that there are users associated with this role.
var ErrRoleHasUsers = errors.New("there are users associated with this role")

// ErrRoleHasPolicies means that there are policies associated with this role.
var ErrRoleHasPolicies = errors.New("there are policies associated with this role")

// recovery recovers from panic and returns the content as error.
func recovery(err *error) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("%s", r)
	}
}
