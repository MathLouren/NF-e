package routes

import (
	"go-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes(destinatarioHandler *handlers.DestinatarioHandler, remetenteHandler *handlers.RemetenteHandler, nfeHandler *handlers.NFeHandler, sefazHandler *handlers.SEFAZHandler) *gin.Engine {
	router := gin.Default()

	// Middleware para CORS
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Rota de health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Grupo de rotas para destinatários
	destinatarios := router.Group("/api/destinatarios")
	{
		destinatarios.POST("/", destinatarioHandler.CreateDestinatario)
		destinatarios.GET("/", destinatarioHandler.GetAllDestinatarios)
		destinatarios.GET("/:id", destinatarioHandler.GetDestinatario)
		destinatarios.PUT("/:id", destinatarioHandler.UpdateDestinatario)
		destinatarios.DELETE("/:id", destinatarioHandler.DeleteDestinatario)
	}

	// Grupo de rotas para remetentes
	remetentes := router.Group("/api/remetentes")
	{
		remetentes.POST("/", remetenteHandler.CreateRemetente)
		remetentes.GET("/", remetenteHandler.GetAllRemetentes)
		remetentes.GET("/:id", remetenteHandler.GetRemetente)
		remetentes.PUT("/:id", remetenteHandler.UpdateRemetente)
		remetentes.DELETE("/:id", remetenteHandler.DeleteRemetente)
	}

	// Grupo de rotas para NFes
	nfes := router.Group("/api/nfes")
	{
		nfes.POST("/", nfeHandler.CreateNFe)
		nfes.GET("/", nfeHandler.GetAllNFes)
		nfes.GET("/:id", nfeHandler.GetNFe)
		nfes.PUT("/:id", nfeHandler.UpdateNFe)
		nfes.DELETE("/:id", nfeHandler.DeleteNFe)
		nfes.POST("/:id/autorizar", nfeHandler.AutorizarNFe)
		nfes.POST("/:id/cancelar", nfeHandler.CancelarNFe)
	}

	// Grupo de rotas para SEFAZ
	sefaz := router.Group("/api/sefaz")
	{
		sefaz.POST("/enviar", sefazHandler.EnviarNFeParaSEFAZ)
		sefaz.GET("/status", sefazHandler.ConsultarStatusSEFAZ)
		sefaz.GET("/xml/:id", sefazHandler.GerarXMLNFe)
	}

	return router
}
