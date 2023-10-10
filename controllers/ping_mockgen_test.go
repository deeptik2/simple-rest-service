package controllers

import (
	"fmt"
	mock_services "github.com/deeptik2/simple-rest-service/mocks"
	"github.com/deeptik2/simple-rest-service/services"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingControllersMockGen(t *testing.T) {
	fakeResponseWriter := httptest.NewRecorder()
	fakeGinContext, _ := gin.CreateTestContext(fakeResponseWriter)

	// prepare mockgen
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockPingController := mock_services.NewMockpingServiceInterface(mockCtrl)
	mockPingController.EXPECT().PingService().Return("pong", fmt.Errorf(http.StatusText(http.StatusOK)))
	services.PingServiceVar = mockPingController

	// now write test
	PingControllers(fakeGinContext)

	// 1st test
	// Given the status is OK, when I hit the endpoint, then I expect 200OK
	if fakeResponseWriter.Code != http.StatusOK {
		t.Errorf("response code should be 200")
	}

	// 2nd test
	// Given the status is OK, when I hit the endpoint, then I expect to see pong
	if fakeResponseWriter.Body.String() != "pong" {
		t.Errorf("response string should be pong")
	}
}
