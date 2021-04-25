package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ContextMock struct {
	JSONCalled bool
}

func (c *ContextMock) JSON(code int, obj interface{}){
	c.JSONCalled = true
}

type PortRepositoryMock struct {
	GetLastStatusResponse struct{
		status string
		err error
	}
	UpdateLastStatusResponse error
}

func (r PortRepositoryMock) GetLastStatus() (string, error) {
	return r.GetLastStatusResponse.status, r.GetLastStatusResponse.err
}

func (r PortRepositoryMock) UpdateLastStatus(status string) error {
	return r.UpdateLastStatusResponse
}

func TestHandlerGetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Mocking response from handler
	PortHandler := NewPortHandler(PortRepositoryMock{
		GetLastStatusResponse: struct {
			status string
			err    error
		}{
			status: "test",
			err: nil,
		},
	})

	rr := httptest.NewRecorder()
	router.GET("/port", PortHandler.GetLastStatus)
	request, err := http.NewRequest(http.MethodGet, "/port", nil)
	router.ServeHTTP(rr, request)

	if err != nil {
		t.Fail()
	}

	expected := `{"status":"tesst"}`

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}