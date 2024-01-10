package main

import (
	"net/http"
	"testing"
	"strings"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)
	
    // здесь нужно добавить необходимые проверки
	//
}