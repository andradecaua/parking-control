package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"parking-control/middleware"
	"parking-control/services"
)

var vagas []services.Vaga

// VerVagas é uma rota get para poder conseguir ver as vagas que temos disponivél no estacionamento
func VerVagas(res http.ResponseWriter, req *http.Request) {
	var token string = req.Header.Get("Authorization")
	if middleware.VerifyExpectedMethod(res, "get", req.Method) {
		if middleware.VerifyAdmin(token, res) {
			services.Db.Table("vagas").Find(&vagas)
			dados, err := json.Marshal(vagas)
			if err != nil {
				fmt.Println(err.Error())
			}
			res.WriteHeader(200)
			res.Header().Set("Content-Type", "application/json")
			res.Write(dados)
		}
	}
}
