package apiserve

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

/* Test for reading files; temporary file creation */
func TestReadFromModule(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.WriteString("42")
	assert.NoError(t, err)
	tmpfile.Close()

	handler := ApiHandler{FileName: tmpfile.Name()}
	randomNumber, err := handler.ReadFromModule()
	assert.NoError(t, err)
	assert.Equal(t, 42, randomNumber)
}

/* Tests GIN starting and reading from temporary file */
func TestGetRandomNumber(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.WriteString("42")
	assert.NoError(t, err)
	tmpfile.Close()

	handler := ApiHandler{FileName: tmpfile.Name()}
	router := gin.Default()
	router.GET("/randnumber", handler.getRandomNumber)

	req, _ := http.NewRequest("GET", "/randnumber", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"randomNumber":42`)
}

/* Tests server creation and response code */
func TestHandleRequests(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.WriteString("42")
	assert.NoError(t, err)
	tmpfile.Close()

	handler := ApiHandler{FileName: tmpfile.Name(), Port: "8080"}
	router := gin.Default()
	router.GET("/randnumber", handler.getRandomNumber)

	ts := httptest.NewServer(router)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/randnumber")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
