package models

import "time"

// Remetente representa os dados do remetente/emissor de uma NFe
type Remetente struct {
	ID         string `json:"id" bson:"_id,omitempty"`
	Nome       string `json:"nome" binding:"required"`
	CPFCNPJ    string `json:"cpf_cnpj" binding:"required"`
	TipoPessoa string `json:"tipo_pessoa" binding:"required"` // "F" para Física, "J" para Jurídica
	Email      string `json:"email" binding:"required,email"`
	Telefone   string `json:"telefone"`

	// Endereço
	Endereco Endereco `json:"endereco" binding:"required"`

	// Dados fiscais
	InscricaoEstadual  string `json:"inscricao_estadual" binding:"required"`
	InscricaoMunicipal string `json:"inscricao_municipal"`
	RegimeTributario   string `json:"regime_tributario" binding:"required"` // "1" - Simples Nacional, "2" - Simples Nacional - excesso, "3" - Regime Normal

	// Dados da empresa
	RazaoSocial  string `json:"razao_social"`
	NomeFantasia string `json:"nome_fantasia"`
	Website      string `json:"website"`

	// Dados bancários
	Banco   string `json:"banco"`
	Agencia string `json:"agencia"`
	Conta   string `json:"conta"`

	// Metadados
	CriadoEm     time.Time `json:"criado_em"`
	AtualizadoEm time.Time `json:"atualizado_em"`
}

// CreateRemetenteRequest representa a requisição para criar um remetente
type CreateRemetenteRequest struct {
	Nome               string   `json:"nome" binding:"required"`
	CPFCNPJ            string   `json:"cpf_cnpj" binding:"required"`
	TipoPessoa         string   `json:"tipo_pessoa" binding:"required"`
	Email              string   `json:"email" binding:"required,email"`
	Telefone           string   `json:"telefone"`
	Endereco           Endereco `json:"endereco" binding:"required"`
	InscricaoEstadual  string   `json:"inscricao_estadual" binding:"required"`
	InscricaoMunicipal string   `json:"inscricao_municipal"`
	RegimeTributario   string   `json:"regime_tributario" binding:"required"`
	RazaoSocial        string   `json:"razao_social"`
	NomeFantasia       string   `json:"nome_fantasia"`
	Website            string   `json:"website"`
	Banco              string   `json:"banco"`
	Agencia            string   `json:"agencia"`
	Conta              string   `json:"conta"`
}

// UpdateRemetenteRequest representa a requisição para atualizar um remetente
type UpdateRemetenteRequest struct {
	Nome               string   `json:"nome"`
	CPFCNPJ            string   `json:"cpf_cnpj"`
	TipoPessoa         string   `json:"tipo_pessoa"`
	Email              string   `json:"email" binding:"omitempty,email"`
	Telefone           string   `json:"telefone"`
	Endereco           Endereco `json:"endereco"`
	InscricaoEstadual  string   `json:"inscricao_estadual"`
	InscricaoMunicipal string   `json:"inscricao_municipal"`
	RegimeTributario   string   `json:"regime_tributario"`
	RazaoSocial        string   `json:"razao_social"`
	NomeFantasia       string   `json:"nome_fantasia"`
	Website            string   `json:"website"`
	Banco              string   `json:"banco"`
	Agencia            string   `json:"agencia"`
	Conta              string   `json:"conta"`
}

// RemetenteResponse representa a resposta da API
type RemetenteResponse struct {
	ID                 string    `json:"id"`
	Nome               string    `json:"nome"`
	CPFCNPJ            string    `json:"cpf_cnpj"`
	TipoPessoa         string    `json:"tipo_pessoa"`
	Email              string    `json:"email"`
	Telefone           string    `json:"telefone"`
	Endereco           Endereco  `json:"endereco"`
	InscricaoEstadual  string    `json:"inscricao_estadual"`
	InscricaoMunicipal string    `json:"inscricao_municipal"`
	RegimeTributario   string    `json:"regime_tributario"`
	RazaoSocial        string    `json:"razao_social"`
	NomeFantasia       string    `json:"nome_fantasia"`
	Website            string    `json:"website"`
	Banco              string    `json:"banco"`
	Agencia            string    `json:"agencia"`
	Conta              string    `json:"conta"`
	CriadoEm           time.Time `json:"criado_em"`
	AtualizadoEm       time.Time `json:"atualizado_em"`
}
