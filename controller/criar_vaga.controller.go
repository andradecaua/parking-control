package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"parking-control/middleware"
	"parking-control/services"
)

type Vaga struct {
	Disponivel bool    `json:"disponivel"`
	Price      float64 `json:"price"`
}

// CriarVaga serve para podermos criar uma vaga no estacionamento
func CriarVaga(res http.ResponseWriter, req *http.Request) {

	if middleware.VerifyExpectedMethod(res, "post", req.Method) {
		if middleware.ContentVerify(req.Header.Get("Content-Type"), res) {
			var vaga Vaga
			body, err := io.ReadAll(req.Body)

			req.Body.Close()
			if err != nil {
				res.WriteHeader(500)
				res.Header().Set("Content-Type", "application/json")
				res.Write([]byte(`{"message": "Ouve um erro ao tentar ler a requisição"}`))
			} else {
				errMarshal := json.Unmarshal(body, &vaga)
				if errMarshal != nil {
					res.WriteHeader(500)
					res.Header().Set("Content-Type", "application/json")
					res.Write([]byte(`{"message": "Ouve um erro ao transformar o arquivo em JSON"}`))
				} else {
					services.Db.Table("vagas").Create(vaga)
					res.WriteHeader(201)
					res.Header().Set("Content-Type", "application/json")
					res.Write([]byte(`{"message": "Vaga criada com sucesso!"}`))
				}
			}
		}
	}
}
