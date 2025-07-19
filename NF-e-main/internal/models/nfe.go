package models

import "time"

// NFe representa uma Nota Fiscal Eletrônica
type NFe struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Numero      string    `json:"numero" binding:"required"`
	Serie       string    `json:"serie" binding:"required"`
	DataEmissao time.Time `json:"data_emissao" binding:"required"`
	DataSaida   time.Time `json:"data_saida" binding:"required"`

	// Dados do remetente
	Remetente Remetente `json:"remetente" binding:"required"`

	// Dados do destinatário
	Destinatario Destinatario `json:"destinatario" binding:"required"`

	// Dados do produto/serviço
	Itens []ItemNFe `json:"itens" binding:"required,min=1"`

	// Valores
	ValorTotal  float64 `json:"valor_total"`
	ValorICMS   float64 `json:"valor_icms"`
	ValorIPI    float64 `json:"valor_ipi"`
	ValorPIS    float64 `json:"valor_pis"`
	ValorCOFINS float64 `json:"valor_cofins"`

	// Status
	Status string `json:"status"` // "rascunho", "autorizada", "cancelada"

	// Metadados
	CriadoEm     time.Time `json:"criado_em"`
	AtualizadoEm time.Time `json:"atualizado_em"`
}

// ItemNFe representa um item da NFe
type ItemNFe struct {
	ID            string  `json:"id" bson:"_id,omitempty"`
	Codigo        string  `json:"codigo" binding:"required"`
	Descricao     string  `json:"descricao" binding:"required"`
	Quantidade    float64 `json:"quantidade" binding:"required,gt=0"`
	UnidadeMedida string  `json:"unidade_medida" binding:"required"`
	ValorUnitario float64 `json:"valor_unitario" binding:"required,gt=0"`
	ValorTotal    float64 `json:"valor_total"`

	// Impostos
	AliquotaICMS   float64 `json:"aliquota_icms"`
	AliquotaIPI    float64 `json:"aliquota_ipi"`
	AliquotaPIS    float64 `json:"aliquota_pis"`
	AliquotaCOFINS float64 `json:"aliquota_cofins"`
}

// CreateNFERequest representa a requisição para criar uma NFe
type CreateNFERequest struct {
	Numero       string       `json:"numero" binding:"required"`
	Serie        string       `json:"serie" binding:"required"`
	DataEmissao  time.Time    `json:"data_emissao" binding:"required"`
	DataSaida    time.Time    `json:"data_saida" binding:"required"`
	Remetente    Remetente    `json:"remetente" binding:"required"`
	Destinatario Destinatario `json:"destinatario" binding:"required"`
	Itens        []ItemNFe    `json:"itens" binding:"required,min=1"`
}

// NFeResponse representa a resposta da API
type NFeResponse struct {
	ID           string       `json:"id"`
	Numero       string       `json:"numero"`
	Serie        string       `json:"serie"`
	DataEmissao  time.Time    `json:"data_emissao"`
	DataSaida    time.Time    `json:"data_saida"`
	Remetente    Remetente    `json:"remetente"`
	Destinatario Destinatario `json:"destinatario"`
	Itens        []ItemNFe    `json:"itens"`
	ValorTotal   float64      `json:"valor_total"`
	ValorICMS    float64      `json:"valor_icms"`
	ValorIPI     float64      `json:"valor_ipi"`
	ValorPIS     float64      `json:"valor_pis"`
	ValorCOFINS  float64      `json:"valor_cofins"`
	Status       string       `json:"status"`
	CriadoEm     time.Time    `json:"criado_em"`
	AtualizadoEm time.Time    `json:"atualizado_em"`
}
