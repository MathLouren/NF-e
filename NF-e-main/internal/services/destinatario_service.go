package services

import (
	"errors"
	"time"

	"go-api/internal/models"
)

// DestinatarioService gerencia as operações relacionadas aos destinatários
type DestinatarioService struct {
	destinatarios map[string]*models.Destinatario
}

// NewDestinatarioService cria uma nova instância do serviço
func NewDestinatarioService() *DestinatarioService {
	return &DestinatarioService{
		destinatarios: make(map[string]*models.Destinatario),
	}
}

// CreateDestinatario cria um novo destinatário
func (s *DestinatarioService) CreateDestinatario(req *models.CreateDestinatarioRequest) (*models.DestinatarioResponse, error) {
	// Validar CPF/CNPJ
	if err := s.validarCPFCNPJ(req.CPFCNPJ, req.TipoPessoa); err != nil {
		return nil, err
	}

	// Verificar se já existe um destinatário com o mesmo CPF/CNPJ
	for _, d := range s.destinatarios {
		if d.CPFCNPJ == req.CPFCNPJ {
			return nil, errors.New("já existe um destinatário com este CPF/CNPJ")
		}
	}

	now := time.Now()
	destinatario := &models.Destinatario{
		ID:                 generateID(),
		Nome:               req.Nome,
		CPFCNPJ:            req.CPFCNPJ,
		TipoPessoa:         req.TipoPessoa,
		Email:              req.Email,
		Telefone:           req.Telefone,
		Endereco:           req.Endereco,
		InscricaoEstadual:  req.InscricaoEstadual,
		InscricaoMunicipal: req.InscricaoMunicipal,
		CriadoEm:           now,
		AtualizadoEm:       now,
	}

	s.destinatarios[destinatario.ID] = destinatario

	return s.toResponse(destinatario), nil
}

// GetDestinatario busca um destinatário pelo ID
func (s *DestinatarioService) GetDestinatario(id string) (*models.DestinatarioResponse, error) {
	destinatario, exists := s.destinatarios[id]
	if !exists {
		return nil, errors.New("destinatário não encontrado")
	}

	return s.toResponse(destinatario), nil
}

// GetAllDestinatarios retorna todos os destinatários
func (s *DestinatarioService) GetAllDestinatarios() []*models.DestinatarioResponse {
	var responses []*models.DestinatarioResponse
	for _, destinatario := range s.destinatarios {
		responses = append(responses, s.toResponse(destinatario))
	}
	return responses
}

// UpdateDestinatario atualiza um destinatário existente
func (s *DestinatarioService) UpdateDestinatario(id string, req *models.UpdateDestinatarioRequest) (*models.DestinatarioResponse, error) {
	destinatario, exists := s.destinatarios[id]
	if !exists {
		return nil, errors.New("destinatário não encontrado")
	}

	// Atualizar campos se fornecidos
	if req.Nome != "" {
		destinatario.Nome = req.Nome
	}
	if req.CPFCNPJ != "" {
		if err := s.validarCPFCNPJ(req.CPFCNPJ, req.TipoPessoa); err != nil {
			return nil, err
		}
		destinatario.CPFCNPJ = req.CPFCNPJ
	}
	if req.TipoPessoa != "" {
		destinatario.TipoPessoa = req.TipoPessoa
	}
	if req.Email != "" {
		destinatario.Email = req.Email
	}
	if req.Telefone != "" {
		destinatario.Telefone = req.Telefone
	}
	if req.Endereco.Logradouro != "" {
		destinatario.Endereco = req.Endereco
	}
	if req.InscricaoEstadual != "" {
		destinatario.InscricaoEstadual = req.InscricaoEstadual
	}
	if req.InscricaoMunicipal != "" {
		destinatario.InscricaoMunicipal = req.InscricaoMunicipal
	}

	destinatario.AtualizadoEm = time.Now()

	return s.toResponse(destinatario), nil
}

// DeleteDestinatario remove um destinatário
func (s *DestinatarioService) DeleteDestinatario(id string) error {
	if _, exists := s.destinatarios[id]; !exists {
		return errors.New("destinatário não encontrado")
	}

	delete(s.destinatarios, id)
	return nil
}

// validarCPFCNPJ valida o formato do CPF ou CNPJ
func (s *DestinatarioService) validarCPFCNPJ(cpfcnpj, tipoPessoa string) error {
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

// toResponse converte um destinatário para a resposta da API
func (s *DestinatarioService) toResponse(destinatario *models.Destinatario) *models.DestinatarioResponse {
	return &models.DestinatarioResponse{
		ID:                 destinatario.ID,
		Nome:               destinatario.Nome,
		CPFCNPJ:            destinatario.CPFCNPJ,
		TipoPessoa:         destinatario.TipoPessoa,
		Email:              destinatario.Email,
		Telefone:           destinatario.Telefone,
		Endereco:           destinatario.Endereco,
		InscricaoEstadual:  destinatario.InscricaoEstadual,
		InscricaoMunicipal: destinatario.InscricaoMunicipal,
		CriadoEm:           destinatario.CriadoEm,
		AtualizadoEm:       destinatario.AtualizadoEm,
	}
}

// generateID gera um ID único (implementação simples para demonstração)
func generateID() string {
	return time.Now().Format("20060102150405")
}
