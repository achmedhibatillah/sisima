package handler_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/dvvnFrtn/sisima/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestGetStudent(t *testing.T) {
	app := testutil.SetupTestApp()

	req := httptest.NewRequest("GET", "/student", nil)

	res, err := app.Test(req)

	assert.NoError(t, err)

	assert.Equal(t, 200, res.StatusCode)

	var body map[string]interface{}
	err = json.NewEncoder(res.Body).Decode(&body)
	assert.NoError(t, err)

	assert.Contains(t, body, "data")

}
