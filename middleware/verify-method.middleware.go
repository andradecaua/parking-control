package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func VerifyExpectedMethod(res http.ResponseWriter, expected, recived string) bool {
	if expected != strings.ToLower(recived) {
		res.WriteHeader(403)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(fmt.Sprintf(`{"message": "Por gentileza enviar o method da request em %s"}`, expected)))
		return false
	}
	return true
}
