package handlers

import (
	"net/http"

	"go-api/internal/models"
	"go-api/internal/services"

	"github.com/gin-gonic/gin"
)

// RemetenteHandler gerencia as requisições HTTP relacionadas aos remetentes
type RemetenteHandler struct {
	service *services.RemetenteService
}

// NewRemetenteHandler cria uma nova instância do handler
func NewRemetenteHandler(service *services.RemetenteService) *RemetenteHandler {
	return &RemetenteHandler{
		service: service,
	}
}

// CreateRemetente cria um novo remetente
func (h *RemetenteHandler) CreateRemetente(c *gin.Context) {
	var req models.CreateRemetenteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	remetente, err := h.service.CreateRemetente(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Remetente criado com sucesso",
		"data":    remetente,
	})
}

// GetRemetente busca um remetente pelo ID
func (h *RemetenteHandler) GetRemetente(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do remetente é obrigatório",
		})
		return
	}

	remetente, err := h.service.GetRemetente(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": remetente,
	})
}

// GetAllRemetentes retorna todos os remetentes
func (h *RemetenteHandler) GetAllRemetentes(c *gin.Context) {
	remetentes := h.service.GetAllRemetentes()

	c.JSON(http.StatusOK, gin.H{
		"data":  remetentes,
		"total": len(remetentes),
	})
}

// UpdateRemetente atualiza um remetente existente
func (h *RemetenteHandler) UpdateRemetente(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do remetente é obrigatório",
		})
		return
	}

	var req models.UpdateRemetenteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	remetente, err := h.service.UpdateRemetente(id, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Remetente atualizado com sucesso",
		"data":    remetente,
	})
}

// DeleteRemetente remove um remetente
func (h *RemetenteHandler) DeleteRemetente(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do remetente é obrigatório",
		})
		return
	}

	err := h.service.DeleteRemetente(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Remetente removido com sucesso",
	})
}
