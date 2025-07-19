package handlers

import (
	"encoding/xml"
	"net/http"

	"go-api/internal/models"
	"go-api/internal/services"

	"github.com/gin-gonic/gin"
)

// SEFAZHandler gerencia as requisições HTTP relacionadas ao SEFAZ
type SEFAZHandler struct {
	sefazService *services.SEFAZService
	nfeService   *services.NFEService
}

// NewSEFAZHandler cria uma nova instância do handler
func NewSEFAZHandler(sefazService *services.SEFAZService, nfeService *services.NFEService) *SEFAZHandler {
	return &SEFAZHandler{
		sefazService: sefazService,
		nfeService:   nfeService,
	}
}

// EnviarNFeParaSEFAZ envia uma NFe para o SEFAZ
func (h *SEFAZHandler) EnviarNFeParaSEFAZ(c *gin.Context) {
	var req models.EnvioSEFAZRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	// Buscar a NFe
	nfe, err := h.nfeService.GetNFe(req.NFeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "NFe não encontrada",
		})
		return
	}

	// Converter para modelo interno
	nfeInterna := &models.NFe{
		ID:           nfe.ID,
		Numero:       nfe.Numero,
		Serie:        nfe.Serie,
		DataEmissao:  nfe.DataEmissao,
		DataSaida:    nfe.DataSaida,
		Remetente:    nfe.Remetente,
		Destinatario: nfe.Destinatario,
		Itens:        nfe.Itens,
		ValorTotal:   nfe.ValorTotal,
		ValorICMS:    nfe.ValorICMS,
		ValorIPI:     nfe.ValorIPI,
		ValorPIS:     nfe.ValorPIS,
		ValorCOFINS:  nfe.ValorCOFINS,
		Status:       nfe.Status,
		CriadoEm:     nfe.CriadoEm,
		AtualizadoEm: nfe.AtualizadoEm,
	}

	// Verificar se a NFe está autorizada
	if nfeInterna.Status != "autorizada" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "NFe deve estar autorizada para ser enviada ao SEFAZ",
		})
		return
	}

	// Gerar XML da NFe
	xmlNFe, err := models.GerarXMLNFe(nfeInterna)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao gerar XML da NFe",
		})
		return
	}

	// Validar XML
	if err := h.sefazService.ValidarXMLNFe(xmlNFe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "XML da NFe inválido: " + err.Error(),
		})
		return
	}

	// Enviar para o SEFAZ
	resposta, err := h.sefazService.EnviarNFeParaSEFAZ(nfeInterna)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao enviar NFe para SEFAZ: " + err.Error(),
		})
		return
	}

	// Gerar XML de resposta
	xmlBytes, err := xml.MarshalIndent(xmlNFe, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao gerar XML de resposta",
		})
		return
	}

	response := models.EnvioSEFAZResponse{
		Status:          resposta.Status,
		Protocolo:       resposta.Protocolo,
		DataAutorizacao: resposta.DataAutorizacao,
		ChaveAcesso:     resposta.ChaveAcesso,
		Mensagem:        resposta.Mensagem,
		XMLGerado:       string(xmlBytes),
		XMLRetorno:      resposta.XMLRetorno,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "NFe enviada para SEFAZ com sucesso",
		"data":    response,
	})
}

// ConsultarStatusSEFAZ consulta o status do serviço do SEFAZ
func (h *SEFAZHandler) ConsultarStatusSEFAZ(c *gin.Context) {
	status, err := h.sefazService.ConsultarStatusSEFAZ()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao consultar status do SEFAZ: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": status,
	})
}

// GerarXMLNFe gera apenas o XML da NFe sem enviar ao SEFAZ
func (h *SEFAZHandler) GerarXMLNFe(c *gin.Context) {
	nfeID := c.Param("id")
	if nfeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da NFe é obrigatório",
		})
		return
	}

	// Buscar a NFe
	nfe, err := h.nfeService.GetNFe(nfeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "NFe não encontrada",
		})
		return
	}

	// Converter para modelo interno
	nfeInterna := &models.NFe{
		ID:           nfe.ID,
		Numero:       nfe.Numero,
		Serie:        nfe.Serie,
		DataEmissao:  nfe.DataEmissao,
		DataSaida:    nfe.DataSaida,
		Remetente:    nfe.Remetente,
		Destinatario: nfe.Destinatario,
		Itens:        nfe.Itens,
		ValorTotal:   nfe.ValorTotal,
		ValorICMS:    nfe.ValorICMS,
		ValorIPI:     nfe.ValorIPI,
		ValorPIS:     nfe.ValorPIS,
		ValorCOFINS:  nfe.ValorCOFINS,
		Status:       nfe.Status,
		CriadoEm:     nfe.CriadoEm,
		AtualizadoEm: nfe.AtualizadoEm,
	}

	// Gerar XML da NFe
	xmlNFe, err := models.GerarXMLNFe(nfeInterna)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao gerar XML da NFe",
		})
		return
	}

	// Validar XML
	if err := h.sefazService.ValidarXMLNFe(xmlNFe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "XML da NFe inválido: " + err.Error(),
		})
		return
	}

	// Gerar XML formatado
	xmlBytes, err := xml.MarshalIndent(xmlNFe, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao formatar XML",
		})
		return
	}

	c.Header("Content-Type", "application/xml")
	c.Header("Content-Disposition", "attachment; filename=nfe_"+nfeID+".xml")
	c.Data(http.StatusOK, "application/xml", xmlBytes)
}
