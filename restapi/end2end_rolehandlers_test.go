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
)

func TestGetRolesHandler200(t *testing.T) {
	t.Cleanup(cleanupRoles)

	var (
		testRole = "test"

		expect = []string{adminRole, guestRole, testRole}
		actual []string
	)

	require.NoError(t, db.RoleDB.Put(testRole, testRole))

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		SetResult(&actual).
		Get(pathRoles)
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.EqualValues(t, expect, actual)
}

func TestGetRolesHandler401(t *testing.T) {
	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Accept", "application/json").
		Get(pathRoles)
	testUnauthorized(t, resp, err)
}

func TestGetRolesHandler403(t *testing.T) {
	resp, err := rclient.R().
		SetHeader("Accept", "application/json").
		Get(pathRoles)
	testForbidden(t, resp, err)
}

func TestPostRolesHandler201(t *testing.T) {
	t.Cleanup(cleanupRoles)

	var (
		testRole       = "test"
		expectResponse = "test"
		expectDB       = testRole
		actualDB       string
		actualResponse string
	)

	encoded, _ := json.Marshal(testRole)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(encoded).
		SetResult(&actualResponse).
		Post(pathRoles)
	require.NoError(t, err)

	assert.Equal(t, 201, resp.StatusCode())

	assert.Equal(t, expectResponse, actualResponse)

	assert.NoError(t, db.RoleDB.Get(testRole, &actualDB))
	assert.Equal(t, expectDB, actualDB)
}

func TestPostRolesHandler401(t *testing.T) {
	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		Post(pathRoles)
	testUnauthorized(t, resp, err)
}

func TestPostRolesHandler403(t *testing.T) {
	resp, err := rclient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		Post(pathRoles)
	testForbidden(t, resp, err)
}

func TestPostRolesHandler409AlreadyExists(t *testing.T) {
	t.Cleanup(cleanupRoles)

	var (
		testRole   = guestRole
		encoded, _ = json.Marshal(testRole)
		expect     = fmt.Sprintf("Role '%s' already exists", testRole)
		actual     string
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(encoded).
		Post(pathRoles)
	require.NoError(t, err)

	assert.Equal(t, 409, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expect, actual)
}

func TestDeleteRolesRoleHandler200(t *testing.T) {
	t.Cleanup(cleanupRoles)

	var (
		testRole = "test"
		expect   = "Deleted"
		actual   string
	)

	require.NoError(t, db.RoleDB.Put(testRole, testRole))

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		SetResult(&actual).
		Delete(fmt.Sprintf("%s/%s", pathRoles, testRole))
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, expect, actual)

	var actualDB string
	assert.ErrorIs(t, db.RoleDB.Get(testRole, &actualDB), skv.ErrNotFound)
}

func TestDeleteRolesRoleHandler401(t *testing.T) {
	t.Cleanup(cleanupRoles)

	var (
		testRole = "test"
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Accept", "application/json").
		Delete(fmt.Sprintf("%s/%s", pathRoles, testRole))
	testUnauthorized(t, resp, err)
}

func TestDeleteRolesRoleHandler403(t *testing.T) {
	t.Cleanup(cleanupRoles)

	var (
		testRole = "test"
	)

	resp, err := rclient.R().
		SetHeader("Accept", "application/json").
		Delete(fmt.Sprintf("%s/%s", pathRoles, testRole))
	testForbidden(t, resp, err)
}

func TestDeleteRolesRoleHandler404(t *testing.T) {
	t.Cleanup(cleanupRoles)

	var (
		testRole = "test"

		expected = fmt.Sprintf("Role '%s' does not exist", testRole)
		actual   string
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		Delete(fmt.Sprintf("%s/%s", pathRoles, testRole))
	require.NoError(t, err)

	assert.Equal(t, 404, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expected, actual)
}

func TestDeleteRolesRoleHandler409HasUsers(t *testing.T) {
	t.Cleanup(cleanupRoles)

	var (
		testRole = adminRole

		expectedResponse = fmt.Sprintf("Role '%s' has associated users", testRole)
		expectedDB       = testRole
		actualResponse   string
		actualDB         string
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		Delete(fmt.Sprintf("%s/%s", pathRoles, testRole))
	require.NoError(t, err)

	assert.Equal(t, 409, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actualResponse))
	assert.Equal(t, expectedResponse, actualResponse)

	assert.NoError(t, db.RoleDB.Get(testRole, &actualDB))
	assert.Equal(t, expectedDB, actualDB)
}

func TestDeleteRolesRoleHandler409HasPolicies(t *testing.T) {
	t.Cleanup(cleanupRoles)
	t.Cleanup(cleanupPolicies)

	var (
		testRole = guestRole

		expectedResponse = fmt.Sprintf("Role '%s' has associated policies", testRole)
		expectedDB       = testRole
		actualResponse   string
		actualDB         string
	)

	ok, err := enforcer.AddPolicy(guestRole, "/api/v1/devices", "GET")
	require.NoError(t, err)
	require.True(t, ok)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		Delete(fmt.Sprintf("%s/%s", pathRoles, testRole))
	require.NoError(t, err)

	assert.Equal(t, 409, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actualResponse))
	assert.Equal(t, expectedResponse, actualResponse)

	assert.NoError(t, db.RoleDB.Get(testRole, &actualDB))
	assert.Equal(t, expectedDB, actualDB)
}

func TestGetRolesRoleUsersHandler200(t *testing.T) {
	t.Cleanup(cleanupRoles)
	t.Cleanup(cleanupUsers)

	var (
		testRole = adminRole

		testUserName = "test"
		testUser     = dbmodels.User{Name: &testUserName, Role: &adminRole}

		expect = []models.User{
			{
				Name: adminName,
				Role: adminRole,
			},
			{
				Name: testUserName,
				Role: adminRole,
			},
		}
		actual []models.User
	)

	require.NoError(t, db.UserDB.Put(testUserName, testUser))

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		SetResult(&actual).
		Get(fmt.Sprintf("%s/%s/%s", pathRoles, testRole, "users"))
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.EqualValues(t, expect, actual)
}

func TestGetRolesRoleUsersHandler401(t *testing.T) {
	t.Cleanup(cleanupRoles)

	var (
		testRole = adminRole
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s/%s", pathRoles, testRole, "users"))
	testUnauthorized(t, resp, err)
}

func TestGetRolesRoleUsersHandler403(t *testing.T) {
	t.Cleanup(cleanupRoles)

	var (
		testRole = adminRole
	)

	resp, err := rclient.R().
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/%s/%s", pathRoles, testRole, "users"))
	testForbidden(t, resp, err)
}
