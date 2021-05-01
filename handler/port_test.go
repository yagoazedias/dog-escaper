package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yagoazedias/dog-escaper/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PortRepositoryMock struct {
	GetLastStatusResponse struct{
		status bool
		err error
	}
	UpdateLastStatusResponse struct{
		port *model.Port
		err error
	}
}

func (r PortRepositoryMock) GetLastStatus() (bool, error) {
	return r.GetLastStatusResponse.status, r.GetLastStatusResponse.err
}

func (r PortRepositoryMock) UpdateLastStatus(bool) (*model.Port, error) {
	return r.UpdateLastStatusResponse.port, r.UpdateLastStatusResponse.err
}

func TestHandlerGetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Mocking response from handler
	PortHandler := NewPortHandler(PortRepositoryMock{
		GetLastStatusResponse: struct {
			status bool
			err    error
		}{
			status: false,
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

	expected := `{"status":false}`

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}