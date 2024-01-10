package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)


// Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое
func TestRequestIsCorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expCode := http.StatusOK

	assert.Equal(t, expCode, responseRecorder.Code)

	assert.NotEmpty(t, responseRecorder.Body)
}

// Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expCount := 4
	expCode := http.StatusOK
	expAnswer := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"

	assert.Equal(t, expCode, responseRecorder.Code)

	serverResponse := responseRecorder.Body.String()
	serverResponseCount := strings.Split(serverResponse, ",")
	assert.Equal(t, expAnswer, serverResponse)
	assert.Len(t, serverResponseCount, expCount)
}
