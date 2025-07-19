package handlers

import (
	"net/http"

	"go-api/internal/models"
	"go-api/internal/services"

	"github.com/gin-gonic/gin"
)

// DestinatarioHandler gerencia as requisições HTTP relacionadas aos destinatários
type DestinatarioHandler struct {
	service *services.DestinatarioService
}

// NewDestinatarioHandler cria uma nova instância do handler
func NewDestinatarioHandler(service *services.DestinatarioService) *DestinatarioHandler {
	return &DestinatarioHandler{
		service: service,
	}
}

// CreateDestinatario cria um novo destinatário
func (h *DestinatarioHandler) CreateDestinatario(c *gin.Context) {
	var req models.CreateDestinatarioRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	destinatario, err := h.service.CreateDestinatario(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Destinatário criado com sucesso",
		"data":    destinatario,
	})
}

// GetDestinatario busca um destinatário pelo ID
func (h *DestinatarioHandler) GetDestinatario(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do destinatário é obrigatório",
		})
		return
	}

	destinatario, err := h.service.GetDestinatario(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": destinatario,
	})
}

// GetAllDestinatarios retorna todos os destinatários
func (h *DestinatarioHandler) GetAllDestinatarios(c *gin.Context) {
	destinatarios := h.service.GetAllDestinatarios()

	c.JSON(http.StatusOK, gin.H{
		"data":  destinatarios,
		"total": len(destinatarios),
	})
}

// UpdateDestinatario atualiza um destinatário existente
func (h *DestinatarioHandler) UpdateDestinatario(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do destinatário é obrigatório",
		})
		return
	}

	var req models.UpdateDestinatarioRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	destinatario, err := h.service.UpdateDestinatario(id, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Destinatário atualizado com sucesso",
		"data":    destinatario,
	})
}

// DeleteDestinatario remove um destinatário
func (h *DestinatarioHandler) DeleteDestinatario(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do destinatário é obrigatório",
		})
		return
	}

	err := h.service.DeleteDestinatario(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Destinatário removido com sucesso",
	})
}
