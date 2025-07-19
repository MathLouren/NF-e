package services

import (
	"errors"
	"time"

	"go-api/internal/models"
)

// NFEService gerencia as operações relacionadas às NFes
type NFEService struct {
	nfeService       *DestinatarioService
	remetenteService *RemetenteService
	nfes             map[string]*models.NFe
}

// NewNFEService cria uma nova instância do serviço
func NewNFEService(destinatarioService *DestinatarioService, remetenteService *RemetenteService) *NFEService {
	return &NFEService{
		nfeService:       destinatarioService,
		remetenteService: remetenteService,
		nfes:             make(map[string]*models.NFe),
	}
}

// CreateNFe cria uma nova NFe
func (s *NFEService) CreateNFe(req *models.CreateNFERequest) (*models.NFeResponse, error) {
	// Validar se o remetente existe (se tiver ID)
	if req.Remetente.ID != "" {
		_, err := s.remetenteService.GetRemetente(req.Remetente.ID)
		if err != nil {
			return nil, errors.New("remetente não encontrado")
		}
	}

	// Validar se o destinatário existe (se tiver ID)
	if req.Destinatario.ID != "" {
		_, err := s.nfeService.GetDestinatario(req.Destinatario.ID)
		if err != nil {
			return nil, errors.New("destinatário não encontrado")
		}
	}

	// Calcular valores dos itens
	for i := range req.Itens {
		req.Itens[i].ValorTotal = req.Itens[i].Quantidade * req.Itens[i].ValorUnitario
	}

	// Calcular totais da NFe
	valorTotal, valorICMS, valorIPI, valorPIS, valorCOFINS := s.calcularTotais(req.Itens)

	now := time.Now()
	nfe := &models.NFe{
		ID:           generateID(),
		Numero:       req.Numero,
		Serie:        req.Serie,
		DataEmissao:  req.DataEmissao,
		DataSaida:    req.DataSaida,
		Remetente:    req.Remetente,
		Destinatario: req.Destinatario,
		Itens:        req.Itens,
		ValorTotal:   valorTotal,
		ValorICMS:    valorICMS,
		ValorIPI:     valorIPI,
		ValorPIS:     valorPIS,
		ValorCOFINS:  valorCOFINS,
		Status:       "rascunho",
		CriadoEm:     now,
		AtualizadoEm: now,
	}

	s.nfes[nfe.ID] = nfe

	return s.toResponse(nfe), nil
}

// GetNFe busca uma NFe pelo ID
func (s *NFEService) GetNFe(id string) (*models.NFeResponse, error) {
	nfe, exists := s.nfes[id]
	if !exists {
		return nil, errors.New("NFe não encontrada")
	}

	return s.toResponse(nfe), nil
}

// GetAllNFes retorna todas as NFes
func (s *NFEService) GetAllNFes() []*models.NFeResponse {
	var responses []*models.NFeResponse
	for _, nfe := range s.nfes {
		responses = append(responses, s.toResponse(nfe))
	}
	return responses
}

// UpdateNFe atualiza uma NFe existente
func (s *NFEService) UpdateNFe(id string, req *models.CreateNFERequest) (*models.NFeResponse, error) {
	nfe, exists := s.nfes[id]
	if !exists {
		return nil, errors.New("NFe não encontrada")
	}

	// Verificar se a NFe pode ser alterada
	if nfe.Status != "rascunho" {
		return nil, errors.New("NFe não pode ser alterada pois não está em rascunho")
	}

	// Atualizar campos
	nfe.Numero = req.Numero
	nfe.Serie = req.Serie
	nfe.DataEmissao = req.DataEmissao
	nfe.DataSaida = req.DataSaida
	nfe.Remetente = req.Remetente
	nfe.Destinatario = req.Destinatario
	nfe.Itens = req.Itens

	// Recalcular valores
	for i := range nfe.Itens {
		nfe.Itens[i].ValorTotal = nfe.Itens[i].Quantidade * nfe.Itens[i].ValorUnitario
	}

	valorTotal, valorICMS, valorIPI, valorPIS, valorCOFINS := s.calcularTotais(nfe.Itens)
	nfe.ValorTotal = valorTotal
	nfe.ValorICMS = valorICMS
	nfe.ValorIPI = valorIPI
	nfe.ValorPIS = valorPIS
	nfe.ValorCOFINS = valorCOFINS

	nfe.AtualizadoEm = time.Now()

	return s.toResponse(nfe), nil
}

// DeleteNFe remove uma NFe
func (s *NFEService) DeleteNFe(id string) error {
	nfe, exists := s.nfes[id]
	if !exists {
		return errors.New("NFe não encontrada")
	}

	if nfe.Status != "rascunho" {
		return errors.New("NFe não pode ser excluída pois não está em rascunho")
	}

	delete(s.nfes, id)
	return nil
}

// AutorizarNFe autoriza uma NFe
func (s *NFEService) AutorizarNFe(id string) (*models.NFeResponse, error) {
	nfe, exists := s.nfes[id]
	if !exists {
		return nil, errors.New("NFe não encontrada")
	}

	if nfe.Status != "rascunho" {
		return nil, errors.New("NFe não pode ser autorizada pois não está em rascunho")
	}

	nfe.Status = "autorizada"
	nfe.AtualizadoEm = time.Now()

	return s.toResponse(nfe), nil
}

// CancelarNFe cancela uma NFe
func (s *NFEService) CancelarNFe(id string) (*models.NFeResponse, error) {
	nfe, exists := s.nfes[id]
	if !exists {
		return nil, errors.New("NFe não encontrada")
	}

	if nfe.Status == "cancelada" {
		return nil, errors.New("NFe já está cancelada")
	}

	nfe.Status = "cancelada"
	nfe.AtualizadoEm = time.Now()

	return s.toResponse(nfe), nil
}

// calcularTotais calcula os valores totais da NFe
func (s *NFEService) calcularTotais(itens []models.ItemNFe) (valorTotal, valorICMS, valorIPI, valorPIS, valorCOFINS float64) {
	for _, item := range itens {
		valorTotal += item.ValorTotal
		valorICMS += (item.ValorTotal * item.AliquotaICMS) / 100
		valorIPI += (item.ValorTotal * item.AliquotaIPI) / 100
		valorPIS += (item.ValorTotal * item.AliquotaPIS) / 100
		valorCOFINS += (item.ValorTotal * item.AliquotaCOFINS) / 100
	}
	return
}

// toResponse converte uma NFe para a resposta da API
func (s *NFEService) toResponse(nfe *models.NFe) *models.NFeResponse {
	return &models.NFeResponse{
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
}
