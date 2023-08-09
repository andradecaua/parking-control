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
	Apelido    string  `json:"apelido"`
}

// CriarVaga serve para podermos criar uma vaga no estacionamento
func CriarVaga(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	res.WriteHeader(http.StatusOK)
	if middleware.VerifyExpectedMethod(res, "post", req.Method) {
		if middleware.ContentVerify(req.Header.Get("Content-Type"), res) {
			var vaga Vaga
			var token string = req.Header.Get("Authorization")
			if middleware.VerifyAdmin(token, res) {
				body, err := io.ReadAll(req.Body)
				if err != nil {
					res.WriteHeader(500)
					res.Header().Set("Content-Type", "application/json")
					res.Write([]byte(`{"message": "Ouve um erro ao tentar ler a requisição"}`))
				} else {
					errMarshal := json.Unmarshal(body, &vaga)
					if vaga.Price != 0 && vaga.Price > 0.00 {
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
					} else {
						res.WriteHeader(400)
						res.Header().Set("Content-Type", "application/json")
						res.Write([]byte(`{"message": "Por gentileza enviar o preço de forma válida "}`))
					}
				}
			}

		}
	}
}
