package handlers

import (
	"net/http"

	"go-api/internal/models"
	"go-api/internal/services"

	"github.com/gin-gonic/gin"
)

// NFeHandler gerencia as requisições HTTP relacionadas às NFes
type NFeHandler struct {
	service *services.NFEService
}

// NewNFeHandler cria uma nova instância do handler
func NewNFeHandler(service *services.NFEService) *NFeHandler {
	return &NFeHandler{
		service: service,
	}
}

// CreateNFe cria uma nova NFe
func (h *NFeHandler) CreateNFe(c *gin.Context) {
	var req models.CreateNFERequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	nfe, err := h.service.CreateNFe(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "NFe criada com sucesso",
		"data":    nfe,
	})
}

// GetNFe busca uma NFe pelo ID
func (h *NFeHandler) GetNFe(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da NFe é obrigatório",
		})
		return
	}

	nfe, err := h.service.GetNFe(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": nfe,
	})
}

// GetAllNFes retorna todas as NFes
func (h *NFeHandler) GetAllNFes(c *gin.Context) {
	nfes := h.service.GetAllNFes()

	c.JSON(http.StatusOK, gin.H{
		"data":  nfes,
		"total": len(nfes),
	})
}

// UpdateNFe atualiza uma NFe existente
func (h *NFeHandler) UpdateNFe(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da NFe é obrigatório",
		})
		return
	}

	var req models.CreateNFERequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	nfe, err := h.service.UpdateNFe(id, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "NFe atualizada com sucesso",
		"data":    nfe,
	})
}

// DeleteNFe remove uma NFe
func (h *NFeHandler) DeleteNFe(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da NFe é obrigatório",
		})
		return
	}

	err := h.service.DeleteNFe(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "NFe removida com sucesso",
	})
}

// AutorizarNFe autoriza uma NFe
func (h *NFeHandler) AutorizarNFe(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da NFe é obrigatório",
		})
		return
	}

	nfe, err := h.service.AutorizarNFe(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "NFe autorizada com sucesso",
		"data":    nfe,
	})
}

// CancelarNFe cancela uma NFe
func (h *NFeHandler) CancelarNFe(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da NFe é obrigatório",
		})
		return
	}

	nfe, err := h.service.CancelarNFe(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "NFe cancelada com sucesso",
		"data":    nfe,
	})
}
