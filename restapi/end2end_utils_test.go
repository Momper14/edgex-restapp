package restapi_test

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-openapi/loads"
	"github.com/go-resty/resty/v2"
	"github.com/momper14/edgex-restapp/db"
	dbmodels "github.com/momper14/edgex-restapp/db/models"
	"github.com/momper14/edgex-restapp/restapi"
	"github.com/momper14/edgex-restapp/restapi/operations"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	pathRoles    = "/api/v1/admin/roles"
	pathPolicies = "/api/v1/admin/policies"
	pathUsers    = "/api/v1/admin/users"
	pathDevices  = "/api/v1/devices"
)

var (
	policyWildcard = ".*"
	adminName      = viper.GetString("admin.username")
	adminPassword  = viper.GetString("admin.password")
	adminRole      = viper.GetString("admin.role")
	guestRole      = viper.GetString("role.default")
	testServer     *httptest.Server
	rclient        *resty.Client
	enforcer       = db.Enforcer
)

func getAPIHandler() (http.Handler, error) {
	var server *restapi.Server

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, err
	}

	flag.Parse()

	api := operations.NewEdgexRestappAPI(swaggerSpec)
	server = restapi.NewServer(api)
	server.ConfigureAPI()

	err = api.Validate()
	if err != nil {
		return nil, err
	}
	return server.GetHandler(), nil
}

func readJSONFile(file string, target interface{}) error {
	fileContent, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileContent, &target)
}

//nolint:errcheck
func cleanupPolicies() {
	enforcer.RemovePolicies(enforcer.GetPolicy())
	enforcer.AddPolicy([]string{adminRole, ".*", ".*"})
	enforcer.SavePolicy()
}

// nolint:errcheck
func cleanupRoles() {
	roles, _ := db.RoleDB.GetKeys()

	for _, r := range roles {
		db.RoleDB.Delete(r)
	}

	db.RoleDB.Put(adminRole, adminRole)
	db.RoleDB.Put(guestRole, guestRole)
}

// nolint:errcheck
func cleanupUsers() {
	var admin *dbmodels.User
	db.UserDB.Get(adminName, &admin)

	users, _ := db.UserDB.GetKeys()

	for _, u := range users {
		db.UserDB.Delete(u)
	}

	db.UserDB.Put(adminName, admin)
}

func testUnauthorized(t *testing.T, resp *resty.Response, err error) {
	var (
		expect = "Wrong username or password"
		actual string
	)

	require.NoError(t, err)

	assert.Equal(t, 401, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expect, actual)
}

func testForbidden(t *testing.T, resp *resty.Response, err error) {
	var (
		expect = "Forbidden"
		actual string
	)

	require.NoError(t, err)
	assert.Equal(t, 403, resp.StatusCode())

	require.NoError(t, json.Unmarshal(resp.Body(), &actual))
	assert.Equal(t, expect, actual)
}

func TestMain(m *testing.M) {
	handler, err := getAPIHandler()
	if err != nil {
		log.Fatal(err)
	}

	testServer = httptest.NewServer(handler)
	rclient = resty.New().SetHostURL(testServer.URL)
	cleanupPolicies()
	cleanupRoles()
	cleanupUsers()

	code := m.Run()

	testServer.Close()
	os.Exit(code)
}
