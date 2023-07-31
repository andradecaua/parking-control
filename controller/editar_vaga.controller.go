package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"parking-control/middleware"
	"parking-control/services"
)

type ConfigEditaVaga struct {
	TypeEdit    string `json:"type_edit"`
	Placa       string `json:"placa"`
	Responsavel string `json:"responsavel"`
	Id          uint   `json:"id_vaga"`
}

//^(?:[A-Z]{3}-?\d{4}|[A-Z]{3}\d[A-Z]\d{2})$

// EditarVaga serve para podermos editar uma vaga do estacionamento
func EditarVaga(res http.ResponseWriter, req *http.Request) {
	if middleware.VerifyExpectedMethod(res, "put", req.Method) {
		if middleware.ContentVerify(req.Header.Get("Content-Type"), res) {
			var token string = req.Header.Get("Authorization")
			if middleware.VerifyAdmin(token, res) {
				body, err := io.ReadAll(req.Body)
				var configEdit ConfigEditaVaga
				if err != nil {
					res.WriteHeader(500)
					res.Header().Set("Content-Type", "application/json")
					res.Write([]byte(`{"message": "Ouve um erro ao ler a resposta"}`))
				} else {
					errJSON := json.Unmarshal(body, &configEdit)
					if configEdit.Id >= 1 {
						if errJSON != nil {
							res.WriteHeader(500)
							res.Header().Set("Content-Type", "application/json")
							res.Write([]byte(`{"message": "Ouve um erro interno, por gentileza contatar o suporte"}`))
						} else {
							switch configEdit.TypeEdit {
							case "ocupar":
								ocuparVaga(configEdit.Id, res, configEdit.Placa, configEdit.Responsavel)
							case "desocupar":
								desocuparVaga(configEdit.Id, res, configEdit.Placa)
							default:
								res.WriteHeader(400)
								res.Header().Set("Content-Type", "application/json")
								res.Write([]byte(`{"message": "Opção de tipo inválida para editar"}`))
							}
						}
					} else {
						res.WriteHeader(400)
						res.Header().Set("Content-Type", "application/json")
						res.Write([]byte(`{"message": "Por gentileza enviar um ID válido para a vaga"}`))
					}
				}
			}
		}
	}

}

func ocuparVaga(id uint, res http.ResponseWriter, placa, responsavel string) {
	var carro services.Carro
	var vaga services.Vaga
	carro.Placa = placa
	carro.Responsavel = responsavel
	services.Db.Table("vagas").First(&vaga, id)

	if middleware.PlacaVerify(placa, res) {
		if vaga.ID >= 1 {
			if vaga.Disponivel {

				services.Db.Table("carros").Where(fmt.Sprintf("placa like '%s'", carro.Placa)).FirstOrCreate(&carro)
				services.Db.Model(&vaga).Where(vaga.ID).Updates(map[string]interface{}{"placa": carro.Placa, "disponivel": false})
				services.Db.Table("vagas").First(&vaga, vaga.ID)
				if vaga.Placa == placa {
					res.WriteHeader(201)
					res.Header().Set("Content-Type", "application/json")
					res.Write([]byte(`{"message": "Vaga editada com sucesso!"}`))
				} else {
					res.WriteHeader(500)
					res.Header().Set("Content-Type", "application/json")
					res.Write([]byte(`{"message": "Ouve um erro interno, por gentileza contatar o suporte"}`))
				}
			} else {
				res.WriteHeader(400)
				res.Header().Set("Content-Type", "application/json")
				res.Write([]byte(`{"message": "Essa vaga já está ocupada"}`))
			}
		} else {
			res.WriteHeader(400)
			res.Header().Set("Content-Type", "application/json")
			res.Write([]byte(`{"message": "Essa vaga não existe"}`))
		}
	}

}

func desocuparVaga(id uint, res http.ResponseWriter, placa string) {
	var vaga services.Vaga
	services.Db.Table("vagas").First(&vaga, id)
	if middleware.PlacaVerify(placa, res) {
		if vaga.ID >= 1 {
			if !vaga.Disponivel {
				if vaga.Placa == placa {
					services.Db.Model(&vaga).Where(vaga.ID).Updates(map[string]interface{}{"placa": "", "disponivel": true})
					if vaga.Disponivel {
						res.WriteHeader(201)
						res.Header().Set("Content-Type", "application/json")
						res.Write([]byte(fmt.Sprintf(`{"message": "Vaga desocupada com sucesso pelo carro de placa %s", "price": %f}`, placa, vaga.Price)))
					} else {
						res.WriteHeader(400)
						res.Header().Set("Content-Type", "application/json")
						res.Write([]byte(`{"message": "Essa vaga não existe "}`))
					}
				} else {
					res.WriteHeader(400)
					res.Header().Set("Content-Type", "application/json")
					res.Write([]byte(`{"message": "O veiculo com esta placa não está ocupando está vaga"}`))
				}
			} else {
				res.WriteHeader(400)
				res.Header().Set("Content-Type", "application/json")
				res.Write([]byte(`{"message": "Essa vaga não está ocupada"}`))
			}
		} else {
			res.WriteHeader(400)
			res.Header().Set("Content-Type", "application/json")
			res.Write([]byte(`{"message": "Essa vaga não existe"}`))
		}
	}
}
