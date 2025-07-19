package main

import (
	"log"

	"go-api/internal/handlers"
	"go-api/internal/routes"
	"go-api/internal/services"
)

func main() {
	// Inicializar serviços
	destinatarioService := services.NewDestinatarioService()
	remetenteService := services.NewRemetenteService()
	nfeService := services.NewNFEService(destinatarioService, remetenteService)
	sefazService := services.NewSEFAZService()

	// Inicializar handlers
	destinatarioHandler := handlers.NewDestinatarioHandler(destinatarioService)
	remetenteHandler := handlers.NewRemetenteHandler(remetenteService)
	nfeHandler := handlers.NewNFeHandler(nfeService)
	sefazHandler := handlers.NewSEFAZHandler(sefazService, nfeService)

	// Configurar rotas
	router := routes.SetupRoutes(destinatarioHandler, remetenteHandler, nfeHandler, sefazHandler)

	// Iniciar servidor
	log.Println("🚀 Servidor iniciado na porta 8000")
	log.Println("📋 Endpoints disponíveis:")
	log.Println("   GET  /ping - Health check")
	log.Println("   POST /api/destinatarios - Criar destinatário")
	log.Println("   GET  /api/destinatarios - Listar destinatários")
	log.Println("   GET  /api/destinatarios/:id - Buscar destinatário")
	log.Println("   PUT  /api/destinatarios/:id - Atualizar destinatário")
	log.Println("   DELETE /api/destinatarios/:id - Remover destinatário")
	log.Println("   POST /api/remetentes - Criar remetente")
	log.Println("   GET  /api/remetentes - Listar remetentes")
	log.Println("   GET  /api/remetentes/:id - Buscar remetente")
	log.Println("   PUT  /api/remetentes/:id - Atualizar remetente")
	log.Println("   DELETE /api/remetentes/:id - Remover remetente")
	log.Println("   POST /api/nfes - Criar NFe")
	log.Println("   GET  /api/nfes - Listar NFes")
	log.Println("   GET  /api/nfes/:id - Buscar NFe")
	log.Println("   PUT  /api/nfes/:id - Atualizar NFe")
	log.Println("   DELETE /api/nfes/:id - Remover NFe")
	log.Println("   POST /api/nfes/:id/autorizar - Autorizar NFe")
	log.Println("   POST /api/nfes/:id/cancelar - Cancelar NFe")
	log.Println("   POST /api/sefaz/enviar - Enviar NFe para SEFAZ")
	log.Println("   GET  /api/sefaz/status - Consultar status SEFAZ")
	log.Println("   GET  /api/sefaz/xml/:id - Gerar XML da NFe")

	if err := router.Run(":8000"); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
