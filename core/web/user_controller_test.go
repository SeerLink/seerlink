package web_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/SeerLink/seerlink/core/services/eth"

	"github.com/SeerLink/seerlink/core/auth"
	"github.com/SeerLink/seerlink/core/internal/cltest"
	"github.com/SeerLink/seerlink/core/store/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserController_UpdatePassword(t *testing.T) {
	t.Parallel()

	rpcClient, gethClient, _, assertMocksCalled := cltest.NewEthMocksWithStartupAssertions(t)
	defer assertMocksCalled()
	app, cleanup := cltest.NewApplicationWithKey(t,
		eth.NewClientWith(rpcClient, gethClient),
	)
	defer cleanup()
	require.NoError(t, app.Start())
	client := app.NewHTTPClient()

	// Invalid request
	resp, cleanup := client.Patch("/v2/user/password", bytes.NewBufferString(""))
	defer cleanup()
	errors := cltest.ParseJSONAPIErrors(t, resp.Body)
	require.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
	assert.Len(t, errors.Errors, 1)

	// Old password is wrong
	resp, cleanup = client.Patch(
		"/v2/user/password",
		bytes.NewBufferString(`{"oldPassword": "wrong password"}`))
	defer cleanup()
	errors = cltest.ParseJSONAPIErrors(t, resp.Body)
	require.Equal(t, http.StatusConflict, resp.StatusCode)
	assert.Len(t, errors.Errors, 1)
	assert.Equal(t, "old password does not match", errors.Errors[0].Detail)

	// Success
	resp, cleanup = client.Patch(
		"/v2/user/password",
		bytes.NewBufferString(fmt.Sprintf(`{"newPassword": "%v", "oldPassword": "%v"}`, cltest.Password, cltest.Password)))
	defer cleanup()
	errors = cltest.ParseJSONAPIErrors(t, resp.Body)
	assert.Len(t, errors.Errors, 0)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUserController_NewAPIToken(t *testing.T) {
	t.Parallel()

	rpcClient, gethClient, _, assertMocksCalled := cltest.NewEthMocksWithStartupAssertions(t)
	defer assertMocksCalled()
	app, cleanup := cltest.NewApplicationWithKey(t,
		eth.NewClientWith(rpcClient, gethClient),
	)
	defer cleanup()
	require.NoError(t, app.Start())

	client := app.NewHTTPClient()
	req, err := json.Marshal(models.ChangeAuthTokenRequest{
		Password: cltest.Password,
	})
	require.NoError(t, err)
	resp, cleanup := client.Post("/v2/user/token", bytes.NewBuffer(req))
	defer cleanup()

	require.Equal(t, http.StatusCreated, resp.StatusCode)
	var authToken auth.Token
	err = cltest.ParseJSONAPIResponse(t, resp, &authToken)
	require.NoError(t, err)
	assert.NotEmpty(t, authToken.AccessKey)
	assert.NotEmpty(t, authToken.Secret)
}

func TestUserController_NewAPIToken_unauthorized(t *testing.T) {
	t.Parallel()

	rpcClient, gethClient, _, assertMocksCalled := cltest.NewEthMocksWithStartupAssertions(t)
	defer assertMocksCalled()
	app, cleanup := cltest.NewApplicationWithKey(t,
		eth.NewClientWith(rpcClient, gethClient),
	)
	defer cleanup()
	require.NoError(t, app.Start())

	client := app.NewHTTPClient()
	req, err := json.Marshal(models.ChangeAuthTokenRequest{
		Password: "wrong-password",
	})
	require.NoError(t, err)
	resp, cleanup := client.Post("/v2/user/token", bytes.NewBuffer(req))
	defer cleanup()
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestUserController_DeleteAPIKey(t *testing.T) {
	t.Parallel()

	rpcClient, gethClient, _, assertMocksCalled := cltest.NewEthMocksWithStartupAssertions(t)
	defer assertMocksCalled()
	app, cleanup := cltest.NewApplicationWithKey(t,
		eth.NewClientWith(rpcClient, gethClient),
	)
	defer cleanup()
	require.NoError(t, app.Start())

	client := app.NewHTTPClient()
	req, err := json.Marshal(models.ChangeAuthTokenRequest{
		Password: cltest.Password,
	})
	require.NoError(t, err)
	resp, cleanup := client.Post("/v2/user/token/delete", bytes.NewBuffer(req))
	defer cleanup()

	require.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestUserController_DeleteAPIKey_unauthorized(t *testing.T) {
	t.Parallel()

	rpcClient, gethClient, _, assertMocksCalled := cltest.NewEthMocksWithStartupAssertions(t)
	defer assertMocksCalled()
	app, cleanup := cltest.NewApplicationWithKey(t,
		eth.NewClientWith(rpcClient, gethClient),
	)
	defer cleanup()
	require.NoError(t, app.Start())

	client := app.NewHTTPClient()
	req, err := json.Marshal(models.ChangeAuthTokenRequest{
		Password: "wrong-password",
	})
	require.NoError(t, err)
	resp, cleanup := client.Post("/v2/user/token/delete", bytes.NewBuffer(req))
	defer cleanup()

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
