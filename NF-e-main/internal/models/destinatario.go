package models

import "time"

// Destinatario representa os dados do destinatário de uma NFe
type Destinatario struct {
	ID         string `json:"id" bson:"_id,omitempty"`
	Nome       string `json:"nome" binding:"required"`
	CPFCNPJ    string `json:"cpf_cnpj" binding:"required"`
	TipoPessoa string `json:"tipo_pessoa" binding:"required"` // "F" para Física, "J" para Jurídica
	Email      string `json:"email" binding:"required,email"`
	Telefone   string `json:"telefone"`

	// Endereço
	Endereco Endereco `json:"endereco" binding:"required"`

	// Dados adicionais
	InscricaoEstadual  string `json:"inscricao_estadual"`
	InscricaoMunicipal string `json:"inscricao_municipal"`

	// Metadados
	CriadoEm     time.Time `json:"criado_em"`
	AtualizadoEm time.Time `json:"atualizado_em"`
}

// Endereco representa o endereço do destinatário
type Endereco struct {
	Logradouro  string `json:"logradouro" binding:"required"`
	Numero      string `json:"numero" binding:"required"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro" binding:"required"`
	Cidade      string `json:"cidade" binding:"required"`
	Estado      string `json:"estado" binding:"required"`
	CEP         string `json:"cep" binding:"required"`
	Pais        string `json:"pais" binding:"required"`
}

// CreateDestinatarioRequest representa a requisição para criar um destinatário
type CreateDestinatarioRequest struct {
	Nome               string   `json:"nome" binding:"required"`
	CPFCNPJ            string   `json:"cpf_cnpj" binding:"required"`
	TipoPessoa         string   `json:"tipo_pessoa" binding:"required"`
	Email              string   `json:"email" binding:"required,email"`
	Telefone           string   `json:"telefone"`
	Endereco           Endereco `json:"endereco" binding:"required"`
	InscricaoEstadual  string   `json:"inscricao_estadual"`
	InscricaoMunicipal string   `json:"inscricao_municipal"`
}

// UpdateDestinatarioRequest representa a requisição para atualizar um destinatário
type UpdateDestinatarioRequest struct {
	Nome               string   `json:"nome"`
	CPFCNPJ            string   `json:"cpf_cnpj"`
	TipoPessoa         string   `json:"tipo_pessoa"`
	Email              string   `json:"email" binding:"omitempty,email"`
	Telefone           string   `json:"telefone"`
	Endereco           Endereco `json:"endereco"`
	InscricaoEstadual  string   `json:"inscricao_estadual"`
	InscricaoMunicipal string   `json:"inscricao_municipal"`
}

// DestinatarioResponse representa a resposta da API
type DestinatarioResponse struct {
	ID                 string    `json:"id"`
	Nome               string    `json:"nome"`
	CPFCNPJ            string    `json:"cpf_cnpj"`
	TipoPessoa         string    `json:"tipo_pessoa"`
	Email              string    `json:"email"`
	Telefone           string    `json:"telefone"`
	Endereco           Endereco  `json:"endereco"`
	InscricaoEstadual  string    `json:"inscricao_estadual"`
	InscricaoMunicipal string    `json:"inscricao_municipal"`
	CriadoEm           time.Time `json:"criado_em"`
	AtualizadoEm       time.Time `json:"atualizado_em"`
}
