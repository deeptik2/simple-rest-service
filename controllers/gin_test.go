package controllers

import (
	"fmt"
	"github.com/deeptik2/simple-rest-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockPingServiceStruct struct {
}

type mockPingServiceStructWithError struct{}

func (service mockPingServiceStruct) PingService() (string, error) {
	fmt.Println("this is a fake")
	return "ponging", nil
}

func (service mockPingServiceStructWithError) PingService() (string, error) {
	fmt.Println("this is a fake func for error")
	err := fmt.Errorf(http.StatusText(http.StatusInternalServerError))
	return "", err
}

func TestPingControllers1(t *testing.T) {
	fakeResponseWriter := httptest.NewRecorder()
	fakeGinContext, _ := gin.CreateTestContext(fakeResponseWriter)

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

func TestPingControllersWithMock(t *testing.T) {
	fakeResponseWriter := httptest.NewRecorder()
	fakeGinContext, _ := gin.CreateTestContext(fakeResponseWriter)

	services.PingServiceVar = mockPingServiceStruct{}

	// now write test
	PingControllers(fakeGinContext)

	// 1st test
	// Given the status is OK, when I hit the endpoint, then I expect 200OK
	if fakeResponseWriter.Code != http.StatusOK {
		t.Errorf("response code should be 200")
	}

	// 2nd test
	// Given the status is OK, when I hit the endpoint, then I expect to see pong
	if fakeResponseWriter.Body.String() != "ponging" {
		t.Errorf("response string should be ponging")
	}
}

func TestPingControllersWithMockError(t *testing.T) {
	fakeResponseWriter := httptest.NewRecorder()
	fakeGinContext, _ := gin.CreateTestContext(fakeResponseWriter)

	services.PingServiceVar = mockPingServiceStructWithError{}

	// now write test
	PingControllers(fakeGinContext)

	// 1st test
	// Given the status is OK, when I hit the endpoint, then I expect 200OK
	if fakeResponseWriter.Code == http.StatusOK {
		t.Errorf("response code should not be 200")
	}

	// 2nd test
	// Given the status is OK, when I hit the endpoint, then I expect to see pong
	if fakeResponseWriter.Body.String() == "ponging" {
		t.Errorf("response string should not be ponging")
	}
}
