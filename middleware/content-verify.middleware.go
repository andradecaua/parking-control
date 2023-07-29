package middleware

import "net/http"

func ContentVerify(contentType string, res http.ResponseWriter) bool {
	if contentType != "application/json" {
		res.WriteHeader(403)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(`{"message": "Por gentileza enviar o conteudo em JSON"}`))
		return false
	}
	return true
}
