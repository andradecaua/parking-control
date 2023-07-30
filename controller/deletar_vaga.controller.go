package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"parking-control/middleware"
	"parking-control/services"
)

type DeleteVaga struct {
	Id    uint   `json:"id_vaga"`
	Token string `json:"token"`
}

// DeletarVaga serve para podermos deletar uma vaga do estacionamento
func DeletarVaga(res http.ResponseWriter, req *http.Request) {

	var deleteVaga DeleteVaga
	var vaga services.Vaga
	var admin services.Admins

	if middleware.VerifyExpectedMethod(res, "delete", req.Method) {
		if middleware.ContentVerify(req.Header.Get("Content-Type"), res) {
			body, err := io.ReadAll(req.Body)
			if err != nil {
				res.WriteHeader(500)
				res.Header().Set("Content-Type", "application/json")
				res.Write([]byte(`{"message": "Ouve um erro ao ler a resposta"}`))
			} else {
				errJSON := json.Unmarshal(body, &deleteVaga)
				if errJSON != nil {
					res.WriteHeader(500)
					res.Header().Set("Content-Type", "application/json")
					res.Write([]byte(`{"message": "Ouve um erro ao converter a resposta em JSON"}`))
				} else {
					services.Db.Where(fmt.Sprintf("token like '%s' and valid = true", deleteVaga.Token)).First(&admin)
					if middleware.VerifyAdmin(deleteVaga.Token, res) {
						services.Db.Table("vagas").First(&vaga, deleteVaga.Id)
						if deleteVaga.Id == vaga.ID {
							services.Db.Delete(&vaga, deleteVaga.Id)
							res.WriteHeader(201)
							res.Header().Set("Content-Type", "application/json")
							res.Write([]byte(`{"message": "Vaga deletado com sucesso!"}`))
						} else {
							res.WriteHeader(400)
							res.Header().Set("Content-Type", "application/json")
							res.Write([]byte(`{"message": "Essa vaga não existe!"}`))
						}
					}
				}
			}
		}
	}

}
