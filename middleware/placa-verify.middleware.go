package middleware

import (
	"net/http"
	"regexp"
	"strings"
)

func PlacaVerify(placa string, res http.ResponseWriter) bool {
	regex := regexp.MustCompile(`^(?:[A-Z]{3}-?\d{4}|[A-Z]{3}\d[A-Z]\d{2})$`)
	if regex.Match([]byte(strings.ToUpper(placa))) {
		return true
	} else {
		res.WriteHeader(400)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(`{"message": "Placa em formato inválido"}`))
		return false
	}
}
