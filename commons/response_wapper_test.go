package commons

import (
	"github.com/SawitProRecruitment/UserService/errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponse(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := Response(c, http.StatusOK, map[string]string{"message": "success"})
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message":"success"}`, rec.Body.String())
}

func TestErrorResponse(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := ErrorResponse(c, http.StatusBadRequest, errors.ErrorInvalidRequest.Error())
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, `{"code":400,"errors":{"list_error":["invalid request"]},"message":""}`, rec.Body.String())
}

func TestErrorResponses(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := ErrorResponses(c, http.StatusBadRequest, errors.ErrorInvalidRequest.Error(), []string{"error 1", "error 2"})
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, `{"code":400,"errors":{"list_error":["error 1","error 2"]},"message":"invalid request"}`, rec.Body.String())
}

func TestSuccessResponse(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := SuccessResponse(c, http.StatusOK, "success")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"code":200, "data":"success", "message":"success"}`, rec.Body.String())
}
