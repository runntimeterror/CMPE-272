package server

/**
 * File created by Rajat Banerjee
 */
 
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTweetsEndpoint(t *testing.T) {
	err := godotenv.Load("../.env")
	assert.NoError(t, err)

	request, err := http.NewRequest("GET", "/tweet", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	Init().router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestCreateAndDeleteTweetEndpoint(t *testing.T) {
	godotenv.Load("../.env")
	s := Init()
	body, err := json.Marshal(map[string]interface{}{
		"status": fmt.Sprintf("test tweet from unit test %v", time.Now().UnixNano()),
	})
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/tweet", bytes.NewBuffer(body))
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	s.router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")

	var b map[string]interface{}

	json.NewDecoder(response.Body).Decode(&b)
	assert.NotNil(t, b)

	assert.NotNil(t, b["id"])

	id := b["id"]

	request, err = http.NewRequest("DELETE", fmt.Sprintf("/tweet/%v", id), nil)
	assert.NoError(t, err)
	response = httptest.NewRecorder()
	s.router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")

}
