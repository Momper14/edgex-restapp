package restapi_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Momper14/skv"
	"github.com/momper14/edgex-restapp/db"
	dbmodels "github.com/momper14/edgex-restapp/db/models"
	"github.com/momper14/edgex-restapp/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestGetUsersHandler200(t *testing.T) {
	t.Cleanup(cleanupUsers)

	var (
		testUserName     = "test"
		testUserRole     = guestRole
		testUserPassword = "password"

		testUser = dbmodels.User{
			Name:     &testUserName,
			Role:     &testUserRole,
			Password: &testUserPassword,
		}

		expect = []models.User{
			{
				Name: adminName,
				Role: adminRole,
			},
			{
				Name: testUserName,
				Role: testUserRole,
			},
		}
		actual []models.User
	)

	require.NoError(t, db.UserDB.Put(testUserName, testUser))

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		SetResult(&actual).
		Get(pathUsers)
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.EqualValues(t, expect, actual)
}

func TestGetUsersHandler401(t *testing.T) {
	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Accept", "application/json").
		Get(pathUsers)
	testUnauthorized(t, resp, err)
}

func TestGetUsersHandler403(t *testing.T) {
	resp, err := rclient.R().
		SetHeader("Accept", "application/json").
		Get(pathUsers)
	testForbidden(t, resp, err)
}

func TestPostUsersHandler201(t *testing.T) {
	t.Cleanup(cleanupUsers)

	var (
		testUserName     = "test"
		testUserRole     = guestRole
		testUserPassword = "password"

		testUser = models.UserCreate{
			Name:     &testUserName,
			Role:     &testUserRole,
			Password: &testUserPassword,
		}

		expectResponse = models.User{
			Name: testUserName,
			Role: testUserRole,
		}
		expectDB = dbmodels.User{
			Name: &testUserName,
			Role: &testUserRole,
		}
		actualResponse models.User
		actualDB       dbmodels.User
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(testUser).
		SetResult(&actualResponse).
		Post(pathUsers)
	require.NoError(t, err)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, expectResponse, actualResponse)

	assert.NoError(t, db.UserDB.Get(testUserName, &actualDB))
	assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(*actualDB.Password), []byte(testUserPassword)))

	expectDB.Password = actualDB.Password
	assert.Equal(t, expectDB, actualDB)
}

func TestPostUsersHandler401(t *testing.T) {
	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		Post(pathUsers)
	testUnauthorized(t, resp, err)
}

func TestPostUsersHandler403(t *testing.T) {
	resp, err := rclient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		Post(pathUsers)
	testForbidden(t, resp, err)
}

func TestPostUsersHandler409WrongRole(t *testing.T) {
	t.Cleanup(cleanupUsers)

	var (
		testUserName     = "test"
		testUserRole     = "missingRole"
		testUserPassword = "password"

		testUser = models.UserCreate{
			Name:     &testUserName,
			Role:     &testUserRole,
			Password: &testUserPassword,
		}

		expect = fmt.Sprintf("Role '%s' does not exist", testUserRole)
		actual string
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(testUser).
		Post(pathUsers)
	require.NoError(t, err)

	assert.Equal(t, 409, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expect, actual)

	assert.ErrorIs(t, db.UserDB.Get(testUserName, &dbmodels.User{}), skv.ErrNotFound)
}

func TestPostUsersHandler409AlreadyExists(t *testing.T) {
	t.Cleanup(cleanupUsers)

	var (
		testUserName     = adminName
		testUserRole     = "missingRole"
		testUserPassword = "password"

		testUser = models.UserCreate{
			Name:     &testUserName,
			Role:     &testUserRole,
			Password: &testUserPassword,
		}

		expect = fmt.Sprintf("User '%s' already exists", testUserName)
		actual string
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(testUser).
		Post(pathUsers)
	require.NoError(t, err)

	assert.Equal(t, 409, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expect, actual)
}

func TestGetUsersUserHandler200(t *testing.T) {
	t.Cleanup(cleanupUsers)

	var (
		testUserName     = "test"
		testUserRole     = guestRole
		testUserPassword = "password"

		testUser = dbmodels.User{
			Name:     &testUserName,
			Role:     &testUserRole,
			Password: &testUserPassword,
		}

		expect = models.User{
			Name: testUserName,
			Role: testUserRole,
		}

		actual models.User
	)

	require.NoError(t, db.UserDB.Put(testUserName, testUser))

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		SetResult(&actual).
		Get(fmt.Sprintf("%s/%s", pathUsers, testUserName))
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.EqualValues(t, expect, actual)
}

func TestGetUsersUserHandler401(t *testing.T) {
	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Accept", "application/json").
		Get(pathUsers)
	testUnauthorized(t, resp, err)
}

func TestGetUsersUserHandler403(t *testing.T) {
	resp, err := rclient.R().
		SetHeader("Accept", "application/json").
		Get(pathUsers)
	testForbidden(t, resp, err)
}

func TestGetUsersUserHandler404(t *testing.T) {
	var (
		testUserName = "missingUser"

		expect = fmt.Sprintf("User '%s' does not exist", testUserName)
		actual string
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		SetResult(&actual).
		Get(fmt.Sprintf("%s/%s", pathUsers, testUserName))
	require.NoError(t, err)

	assert.Equal(t, 404, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expect, actual)
}

func TestPatchUsersUserHandler200(t *testing.T) {
	t.Cleanup(cleanupUsers)

	var (
		testUserName     = "test"
		testUserRole     = guestRole
		testUserPassword = "password"

		testUser = dbmodels.User{
			Name:     &testUserName,
			Role:     &testUserRole,
			Password: &testUserPassword,
		}

		newPassword = "newPassword"
		newRole     = adminRole

		update = models.UserUpdate{
			Password: &newPassword,
			Role:     &newRole,
		}

		expectResponse = "Updated"
		expectDB       = dbmodels.User{
			Name: &testUserName,
			Role: &newRole,
		}
		actualResponse string
		actualDB       dbmodels.User
	)

	require.NoError(t, db.UserDB.Put(testUserName, testUser))

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(update).
		SetResult(&actualResponse).
		Patch(fmt.Sprintf("%s/%s", pathUsers, testUserName))
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, expectResponse, actualResponse)

	assert.NoError(t, db.UserDB.Get(testUserName, &actualDB))
	assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(*actualDB.Password), []byte(newPassword)))

	expectDB.Password = actualDB.Password
	assert.Equal(t, expectDB, actualDB)
}

func TestPatchUsersUserHandler400(t *testing.T) {
	var (
		testUserName = "test"

		update = models.UserUpdate{}

		expect = "Nothing to update"
		actual string
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(update).
		Patch(fmt.Sprintf("%s/%s", pathUsers, testUserName))
	require.NoError(t, err)

	assert.Equal(t, 400, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expect, actual)
}

func TestPatchUsersUserHandler401(t *testing.T) {
	var (
		testUserName = "test"
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		Patch(fmt.Sprintf("%s/%s", pathUsers, testUserName))
	testUnauthorized(t, resp, err)
}

func TestPatchUsersUserHandler403(t *testing.T) {
	var (
		testUserName = "test"
	)

	resp, err := rclient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		Patch(fmt.Sprintf("%s/%s", pathUsers, testUserName))
	testForbidden(t, resp, err)
}

func TestPatchUsersUserHandler404(t *testing.T) {
	var (
		testUserName = "missingUser"

		newPassword = "newPassword"
		newRole     = adminRole

		update = models.UserUpdate{
			Password: &newPassword,
			Role:     &newRole,
		}

		expect = fmt.Sprintf("User '%s' does not exist", testUserName)
		actual string
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(update).
		Patch(fmt.Sprintf("%s/%s", pathUsers, testUserName))
	require.NoError(t, err)

	assert.Equal(t, 404, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expect, actual)
}

func TestDeleteUsersUserHandler200(t *testing.T) {
	t.Cleanup(cleanupUsers)

	var (
		testUserName     = "test"
		testUserRole     = guestRole
		testUserPassword = "password"

		testUser = dbmodels.User{
			Name:     &testUserName,
			Role:     &testUserRole,
			Password: &testUserPassword,
		}

		expect = "Deleted"
		actual string
	)

	require.NoError(t, db.UserDB.Put(testUserName, testUser))

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		SetResult(&actual).
		Delete(fmt.Sprintf("%s/%s", pathUsers, testUserName))
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, expect, actual)

	assert.ErrorIs(t, db.RoleDB.Get(testUserName, &dbmodels.User{}), skv.ErrNotFound)
}

func TestDeleteUsersUserHandler401(t *testing.T) {
	t.Cleanup(cleanupRoles)

	var (
		testRole = "test"
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Accept", "application/json").
		Delete(fmt.Sprintf("%s/%s", pathUsers, testRole))
	testUnauthorized(t, resp, err)
}

func TestDeleteUsersUserHandler403(t *testing.T) {
	t.Cleanup(cleanupRoles)

	var (
		testRole = "test"
	)

	resp, err := rclient.R().
		SetHeader("Accept", "application/json").
		Delete(fmt.Sprintf("%s/%s", pathUsers, testRole))
	testForbidden(t, resp, err)
}

func TestDeleteUsersUserHandler404(t *testing.T) {
	t.Cleanup(cleanupRoles)

	var (
		testUserName = "test"

		expected = fmt.Sprintf("User '%s' does not exist", testUserName)
		actual   string
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		Delete(fmt.Sprintf("%s/%s", pathUsers, testUserName))
	require.NoError(t, err)

	assert.Equal(t, 404, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expected, actual)
}
