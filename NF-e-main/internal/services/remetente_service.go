package services

import (
	"errors"
	"time"

	"go-api/internal/models"
)

// RemetenteService gerencia as operações relacionadas aos remetentes
type RemetenteService struct {
	remetentes map[string]*models.Remetente
}

// NewRemetenteService cria uma nova instância do serviço
func NewRemetenteService() *RemetenteService {
	return &RemetenteService{
		remetentes: make(map[string]*models.Remetente),
	}
}

// CreateRemetente cria um novo remetente
func (s *RemetenteService) CreateRemetente(req *models.CreateRemetenteRequest) (*models.RemetenteResponse, error) {
	// Validar CPF/CNPJ
	if err := s.validarCPFCNPJ(req.CPFCNPJ, req.TipoPessoa); err != nil {
		return nil, err
	}

	// Validar regime tributário
	if err := s.validarRegimeTributario(req.RegimeTributario); err != nil {
		return nil, err
	}

	// Verificar se já existe um remetente com o mesmo CPF/CNPJ
	for _, r := range s.remetentes {
		if r.CPFCNPJ == req.CPFCNPJ {
			return nil, errors.New("já existe um remetente com este CPF/CNPJ")
		}
	}

	now := time.Now()
	remetente := &models.Remetente{
		ID:                 generateID(),
		Nome:               req.Nome,
		CPFCNPJ:            req.CPFCNPJ,
		TipoPessoa:         req.TipoPessoa,
		Email:              req.Email,
		Telefone:           req.Telefone,
		Endereco:           req.Endereco,
		InscricaoEstadual:  req.InscricaoEstadual,
		InscricaoMunicipal: req.InscricaoMunicipal,
		RegimeTributario:   req.RegimeTributario,
		RazaoSocial:        req.RazaoSocial,
		NomeFantasia:       req.NomeFantasia,
		Website:            req.Website,
		Banco:              req.Banco,
		Agencia:            req.Agencia,
		Conta:              req.Conta,
		CriadoEm:           now,
		AtualizadoEm:       now,
	}

	s.remetentes[remetente.ID] = remetente

	return s.toResponse(remetente), nil
}

// GetRemetente busca um remetente pelo ID
func (s *RemetenteService) GetRemetente(id string) (*models.RemetenteResponse, error) {
	remetente, exists := s.remetentes[id]
	if !exists {
		return nil, errors.New("remetente não encontrado")
	}

	return s.toResponse(remetente), nil
}

// GetAllRemetentes retorna todos os remetentes
func (s *RemetenteService) GetAllRemetentes() []*models.RemetenteResponse {
	var responses []*models.RemetenteResponse
	for _, remetente := range s.remetentes {
		responses = append(responses, s.toResponse(remetente))
	}
	return responses
}

// UpdateRemetente atualiza um remetente existente
func (s *RemetenteService) UpdateRemetente(id string, req *models.UpdateRemetenteRequest) (*models.RemetenteResponse, error) {
	remetente, exists := s.remetentes[id]
	if !exists {
		return nil, errors.New("remetente não encontrado")
	}

	// Atualizar campos se fornecidos
	if req.Nome != "" {
		remetente.Nome = req.Nome
	}
	if req.CPFCNPJ != "" {
		if err := s.validarCPFCNPJ(req.CPFCNPJ, req.TipoPessoa); err != nil {
			return nil, err
		}
		remetente.CPFCNPJ = req.CPFCNPJ
	}
	if req.TipoPessoa != "" {
		remetente.TipoPessoa = req.TipoPessoa
	}
	if req.Email != "" {
		remetente.Email = req.Email
	}
	if req.Telefone != "" {
		remetente.Telefone = req.Telefone
	}
	if req.Endereco.Logradouro != "" {
		remetente.Endereco = req.Endereco
	}
	if req.InscricaoEstadual != "" {
		remetente.InscricaoEstadual = req.InscricaoEstadual
	}
	if req.InscricaoMunicipal != "" {
		remetente.InscricaoMunicipal = req.InscricaoMunicipal
	}
	if req.RegimeTributario != "" {
		if err := s.validarRegimeTributario(req.RegimeTributario); err != nil {
			return nil, err
		}
		remetente.RegimeTributario = req.RegimeTributario
	}
	if req.RazaoSocial != "" {
		remetente.RazaoSocial = req.RazaoSocial
	}
	if req.NomeFantasia != "" {
		remetente.NomeFantasia = req.NomeFantasia
	}
	if req.Website != "" {
		remetente.Website = req.Website
	}
	if req.Banco != "" {
		remetente.Banco = req.Banco
	}
	if req.Agencia != "" {
		remetente.Agencia = req.Agencia
	}
	if req.Conta != "" {
		remetente.Conta = req.Conta
	}

	remetente.AtualizadoEm = time.Now()

	return s.toResponse(remetente), nil
}

// DeleteRemetente remove um remetente
func (s *RemetenteService) DeleteRemetente(id string) error {
	if _, exists := s.remetentes[id]; !exists {
		return errors.New("remetente não encontrado")
	}

	delete(s.remetentes, id)
	return nil
}

// validarCPFCNPJ valida o formato do CPF ou CNPJ
func (s *RemetenteService) validarCPFCNPJ(cpfcnpj, tipoPessoa string) error {
	if tipoPessoa == "F" {
		if len(cpfcnpj) != 11 {
			return errors.New("CPF deve ter 11 dígitos")
		}
	} else if tipoPessoa == "J" {
		if len(cpfcnpj) != 14 {
			return errors.New("CNPJ deve ter 14 dígitos")
		}
	} else {
		return errors.New("tipo de pessoa deve ser 'F' (Física) ou 'J' (Jurídica)")
	}

	return nil
}

// validarRegimeTributario valida o regime tributário
func (s *RemetenteService) validarRegimeTributario(regime string) error {
	regimesValidos := map[string]string{
		"1": "Simples Nacional",
		"2": "Simples Nacional - excesso",
		"3": "Regime Normal",
	}

	if _, exists := regimesValidos[regime]; !exists {
		return errors.New("regime tributário deve ser '1' (Simples Nacional), '2' (Simples Nacional - excesso) ou '3' (Regime Normal)")
	}

	return nil
}

// toResponse converte um remetente para a resposta da API
func (s *RemetenteService) toResponse(remetente *models.Remetente) *models.RemetenteResponse {
	return &models.RemetenteResponse{
		ID:                 remetente.ID,
		Nome:               remetente.Nome,
		CPFCNPJ:            remetente.CPFCNPJ,
		TipoPessoa:         remetente.TipoPessoa,
		Email:              remetente.Email,
		Telefone:           remetente.Telefone,
		Endereco:           remetente.Endereco,
		InscricaoEstadual:  remetente.InscricaoEstadual,
		InscricaoMunicipal: remetente.InscricaoMunicipal,
		RegimeTributario:   remetente.RegimeTributario,
		RazaoSocial:        remetente.RazaoSocial,
		NomeFantasia:       remetente.NomeFantasia,
		Website:            remetente.Website,
		Banco:              remetente.Banco,
		Agencia:            remetente.Agencia,
		Conta:              remetente.Conta,
		CriadoEm:           remetente.CriadoEm,
		AtualizadoEm:       remetente.AtualizadoEm,
	}
}
