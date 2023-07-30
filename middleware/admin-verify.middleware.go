package middleware

import (
	"fmt"
	"net/http"
	"parking-control/services"
)

func VerifyAdmin(token string, res http.ResponseWriter) bool {
	var user services.Admins

	if token != "" {
		services.Db.Table("admins").Where(fmt.Sprintf("token like '%s' and valid = true", token)).First(&user)
		if user.Token == token {
			if user.Valid {
				return true
			}
			res.WriteHeader(403)
			res.Header().Set("Content-Type", "application/json")
			res.Write([]byte(`{"message": "Usuário inválido"}`))
			return false
		} else {
			res.WriteHeader(403)
			res.Header().Set("Content-Type", "application/json")
			res.Write([]byte(`{"message": "Token ínvalido"}`))
			return false
		}
	}
	res.WriteHeader(403)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"message": "Por gentileza enviar um token para analíse"}`))
	return false
}
