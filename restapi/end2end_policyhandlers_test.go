package restapi_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/momper14/edgex-restapp/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPoliciesHandler200(t *testing.T) {
	t.Cleanup(cleanupPolicies)

	var (
		testRole     = guestRole
		testResource = "/api/v1/devices"
		testMethod   = "GET"

		expect = []models.Policy{
			{
				Role:     &adminRole,
				Resource: &policyWildcard,
				Method:   &policyWildcard,
			},
			{
				Role:     &testRole,
				Resource: &testResource,
				Method:   &testMethod,
			},
		}
		actual []models.Policy
	)

	ok, err := enforcer.AddPolicy(testRole, testResource, testMethod)
	require.NoError(t, err)
	require.True(t, ok)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Accept", "application/json").
		SetResult(&actual).
		Get(pathPolicies)
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.EqualValues(t, expect, actual)
}

func TestGetPoliciesHandler401(t *testing.T) {
	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Accept", "application/json").
		Get(pathPolicies)
	testUnauthorized(t, resp, err)
}

func TestGetPoliciesHandler403(t *testing.T) {
	resp, err := rclient.R().
		SetHeader("Accept", "application/json").
		Get(pathPolicies)
	testForbidden(t, resp, err)
}

func TestPostPoliciesHandler201(t *testing.T) {
	t.Cleanup(cleanupPolicies)

	var (
		testRole     = guestRole
		testResource = "/api/v1/devices"
		testMethod   = "GET"

		expect = models.Policy{
			Role:     &testRole,
			Resource: &testResource,
			Method:   &testMethod,
		}
		actual models.Policy
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(expect).
		SetResult(&actual).
		Post(pathPolicies)
	require.NoError(t, err)

	assert.Equal(t, 201, resp.StatusCode())
	assert.EqualValues(t, expect, actual)
	assert.True(t, enforcer.HasPolicy(testRole, testResource, testMethod))
}

func TestPostPoliciesHandler401(t *testing.T) {
	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		Post(pathPolicies)
	testUnauthorized(t, resp, err)
}

func TestPostPoliciesHandler403(t *testing.T) {
	resp, err := rclient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		Post(pathPolicies)
	testForbidden(t, resp, err)
}

func TestPostPoliciesHandler409WrongRole(t *testing.T) {
	t.Cleanup(cleanupPolicies)

	var (
		testRole     = "notExist"
		testResource = "/api/v1/devices"
		testMethod   = "GET"

		body = models.Policy{
			Role:     &testRole,
			Resource: &testResource,
			Method:   &testMethod,
		}

		expect = fmt.Sprintf("Role '%s' does not exist", testRole)
		actual string
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(body).
		Post(pathPolicies)
	require.NoError(t, err)

	assert.Equal(t, 409, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expect, actual)

	assert.False(t, enforcer.HasPolicy(testRole, testResource, testMethod))
}

func TestPostPoliciesHandler409AlreadyExists(t *testing.T) {
	t.Cleanup(cleanupPolicies)

	var (
		testRole     = adminRole
		testResource = policyWildcard
		testMethod   = policyWildcard

		body = models.Policy{
			Role:     &testRole,
			Resource: &testResource,
			Method:   &testMethod,
		}
		encoded, _ = json.Marshal(body)

		expect = fmt.Sprintf("Policy '%s' already exists", encoded)
		actual string
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(body).
		Post(pathPolicies)
	require.NoError(t, err)

	assert.Equal(t, 409, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expect, actual)
}

func TestDeletePoliciesHandler200(t *testing.T) {
	t.Cleanup(cleanupPolicies)

	var (
		testRole     = guestRole
		testResource = "/api/v1/devices"
		testMethod   = "GET"

		body = models.Policy{
			Role:     &testRole,
			Resource: &testResource,
			Method:   &testMethod,
		}

		expect = "Deleted"
		actual string
	)

	ok, err := enforcer.AddPolicy(testRole, testResource, testMethod)
	require.NoError(t, err)
	require.True(t, ok)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(body).
		SetResult(&actual).
		Delete(pathPolicies)
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, expect, actual)

	assert.False(t, enforcer.HasPolicy(testRole, testResource, testMethod))
}

func TestDeletePoliciesHandler401(t *testing.T) {
	resp, err := rclient.R().
		SetBasicAuth(adminName, "wrongpassword").
		SetHeader("Accept", "application/json").
		Delete(pathPolicies)
	testUnauthorized(t, resp, err)
}

func TestDeletePoliciesHandler403(t *testing.T) {
	resp, err := rclient.R().
		SetHeader("Accept", "application/json").
		Delete(pathPolicies)
	testForbidden(t, resp, err)
}

func TestDeletePoliciesHandler404(t *testing.T) {
	var (
		testRole     = guestRole
		testResource = "/api/v1/devices"
		testMethod   = "GET"

		body = models.Policy{
			Role:     &testRole,
			Resource: &testResource,
			Method:   &testMethod,
		}

		encoded, _ = json.Marshal(body)

		expected = fmt.Sprintf("Policy '%s' does not exist", encoded)
		actual   string
	)

	resp, err := rclient.R().
		SetBasicAuth(adminName, adminPassword).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(body).
		Delete(pathPolicies)
	require.NoError(t, err)

	assert.Equal(t, 404, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expected, actual)

	assert.False(t, enforcer.HasPolicy(testRole, testResource, testMethod))
}
