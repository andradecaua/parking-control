package main

import (
	"log"
	"net/http"
	"parking-control/controller"
	"parking-control/services"
)

func main() {
	services.DbStart()
	http.HandleFunc("/ver-vagas", controller.VerVagas)
	http.HandleFunc("/criar-vaga", controller.CriarVaga)
	//http.HandleFunc("/editar-vaga", controller.EditarVaga)
	http.HandleFunc("/deletar-vaga", controller.DeletarVaga)
	log.Fatal(http.ListenAndServe(":80", nil))
}
