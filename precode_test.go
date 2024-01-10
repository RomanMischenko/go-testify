package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func createRequest(url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

// Город, который передаётся в параметре city, не поддерживается. 
// Сервис возвращает код ответа 400 и ошибку wrong count value в теле ответа.
func TestIncorrectCityInRequest(t *testing.T) {
	responseRecorder := createRequest("/cafe?count=2&city=tokio")

	expCode := http.StatusBadRequest
	expAnswer := "wrong city value"

	assert.Equal(t, expCode, responseRecorder.Code)
	serverResponse := responseRecorder.Body.String()
	assert.Equal(t, expAnswer, serverResponse)
}

// Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое
func TestRequestIsCorrect(t *testing.T) {
	responseRecorder := createRequest("/cafe?count=2&city=moscow")

	expCode := http.StatusOK

	assert.Equal(t, expCode, responseRecorder.Code)

	assert.NotEmpty(t, responseRecorder.Body)
}

// Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	responseRecorder := createRequest("/cafe?count=10&city=moscow")

	expCount := 4
	expCode := http.StatusOK
	expAnswer := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"

	assert.Equal(t, expCode, responseRecorder.Code)

	serverResponse := responseRecorder.Body.String()
	serverResponseCount := strings.Split(serverResponse, ",")
	assert.Equal(t, expAnswer, serverResponse)
	assert.Len(t, serverResponseCount, expCount)
}
